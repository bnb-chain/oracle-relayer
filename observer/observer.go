package observer

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/binance-chain/oracle-relayer/common"
	"github.com/binance-chain/oracle-relayer/executor/bsc"
	"github.com/binance-chain/oracle-relayer/model"
	"github.com/binance-chain/oracle-relayer/util"
)

type Observer struct {
	DB          *gorm.DB
	Config      *util.Config
	BscExecutor *bsc.Executor
}

func NewObserver(db *gorm.DB, cfg *util.Config, bscExecutor *bsc.Executor) *Observer {
	return &Observer{
		DB:          db,
		Config:      cfg,
		BscExecutor: bscExecutor,
	}
}

func (ob *Observer) Start() {
	go ob.Fetch(ob.Config.ChainConfig.BSCStartHeight)
	go ob.Prune()
	go ob.Alert()
}

func (ob *Observer) Fetch(startHeight int64) {
	for {
		curBlockLog := ob.GetCurrentBlockLog()

		nextHeight := curBlockLog.Height + 1
		if curBlockLog.Height == 0 && startHeight != 0 {
			nextHeight = startHeight
		}

		util.Logger.Infof("fetch block, height=%d", nextHeight)
		err := ob.fetchBlock(curBlockLog.Height, nextHeight, curBlockLog.BlockHash)
		if err != nil {
			util.Logger.Errorf("fetch block error, err=%s", err.Error())
			time.Sleep(2 * time.Second)
		}
	}
}

func (ob *Observer) fetchBlock(curHeight, nextHeight int64, curBlockHash string) error {
	blockAndTxLogs, err := ob.BscExecutor.GetBlockAndTxs(nextHeight)
	if err != nil {
		return fmt.Errorf("get block info error, height=%d, err=%s", nextHeight, err.Error())
	}

	parentHash := blockAndTxLogs.ParentBlockHash
	if curHeight != 0 && parentHash != curBlockHash {
		return ob.DeleteBlockAndTxs(curHeight)
	} else {
		nextBlockLog := model.BlockLog{
			BlockHash:  blockAndTxLogs.BlockHash,
			ParentHash: parentHash,
			Height:     blockAndTxLogs.Height,
			BlockTime:  blockAndTxLogs.BlockTime,
		}

		err := ob.SaveBlockAndTxs(&nextBlockLog, blockAndTxLogs.TxLogs)
		if err != nil {
			return err
		}

		err = ob.UpdateConfirmedNum(nextBlockLog.Height)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ob *Observer) DeleteBlockAndTxs(height int64) error {
	tx := ob.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("height = ?", height).Delete(model.BlockLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("height = ? and status = ?", height, model.TxStatusInit).Delete(model.ClaimLog{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ob *Observer) UpdateConfirmedNum(height int64) error {
	err := ob.DB.Model(model.ClaimLog{}).Where("status = ?", model.TxStatusInit).Updates(
		map[string]interface{}{
			"confirmed_num": gorm.Expr("? - height", height),
			"update_time":   time.Now().Unix(),
		}).Error
	if err != nil {
		return err
	}

	err = ob.DB.Model(model.ClaimLog{}).Where("status = ? and confirmed_num >= ? ",
		model.TxStatusInit, ob.Config.ChainConfig.BSCConfirmNum).Updates(
		map[string]interface{}{
			"status":      model.TxStatusConfirmed,
			"update_time": time.Now().Unix(),
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (ob *Observer) Prune() {
	for {
		curOtherChainBlockLog := ob.GetCurrentBlockLog()
		ob.DB.Where("height < ?", curOtherChainBlockLog.Height-common.ObserverMaxBlockNumber).Delete(model.BlockLog{})
		time.Sleep(common.ObserverPruneInterval)
	}
}

func (ob *Observer) SaveBlockAndTxs(blockLog *model.BlockLog, txLogs []interface{}) error {
	tx := ob.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(blockLog).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, txLog := range txLogs {
		if err := tx.Create(txLog).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (ob *Observer) GetCurrentBlockLog() *model.BlockLog {
	blockLog := model.BlockLog{}
	ob.DB.Order("height desc").First(&blockLog)
	return &blockLog
}

func (ob *Observer) Alert() {
	for {
		curOtherChainBlockLog := ob.GetCurrentBlockLog()
		if curOtherChainBlockLog.Height > 0 {
			if time.Now().Unix()-curOtherChainBlockLog.CreateTime > ob.Config.AlertConfig.BlockUpdateTimeOut {
				msg := fmt.Sprintf("[%s] last smart chain block fetched at %s, height=%d",
					ob.Config.AlertConfig.Moniker, time.Unix(curOtherChainBlockLog.CreateTime, 0).String(), curOtherChainBlockLog.Height)
				util.SendTelegramMessage(ob.Config.AlertConfig.TelegramBotId, ob.Config.AlertConfig.TelegramChatId, msg)
			}
		}

		time.Sleep(common.ObserverAlertInterval)
	}
}
