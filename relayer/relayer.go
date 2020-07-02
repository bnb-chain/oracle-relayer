package relayer

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jinzhu/gorm"

	"github.com/binance-chain/oracle-relayer/common"
	"github.com/binance-chain/oracle-relayer/executor/bbc"
	"github.com/binance-chain/oracle-relayer/model"
	"github.com/binance-chain/oracle-relayer/util"
)

type Relayer struct {
	DB          *gorm.DB
	BBCExecutor *bbc.Executor
	Config      *util.Config
}

func NewRelayer(db *gorm.DB, bbcExecutor *bbc.Executor, cfg *util.Config) *Relayer {
	return &Relayer{
		DB:          db,
		BBCExecutor: bbcExecutor,
		Config:      cfg,
	}
}

func (r *Relayer) Main() {
	go r.RelayPackages()

	go r.Alert()
}

func (r *Relayer) RelayPackages() {
	for {
		err := r.process(r.Config.ChainConfig.BSCChainId)
		if err != nil {
			time.Sleep(time.Duration(r.Config.ChainConfig.RelayInterval) * time.Millisecond)
		}
	}
}

func (r *Relayer) process(chainId uint16) error {
	sequence, err := r.BBCExecutor.GetCurrentSequence(chainId)
	if err != nil {
		util.Logger.Errorf("get current sequence error: chainId=%d, err=%s",
			chainId, err.Error())
		return err
	}

	util.Logger.Infof("current sequence, chain_id=%d, seq=%d", chainId, sequence)

	claimLogs := make([]*model.CrossChainPackageLog, 0)
	err = r.DB.Where("oracle_sequence = ? and chain_id = ? and status = ?",
		sequence, chainId, model.PackageStatusConfirmed).Order("tx_index asc").Find(&claimLogs).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			util.Logger.Errorf("query claim log error: err=%s", err.Error())
		}
		return err
	}

	if len(claimLogs) == 0 {
		return fmt.Errorf("no packages found")
	}

	prophecy, err := r.BBCExecutor.GetProphecy(chainId, sequence)
	if err != nil {
		util.Logger.Errorf("get prophecy error: err=%s", err.Error())
		return err
	}

	validatorAddress := r.BBCExecutor.GetAddress()
	if prophecy != nil && prophecy.ValidatorClaims != nil && prophecy.ValidatorClaims[validatorAddress.String()] != "" {
		return fmt.Errorf("already claimed")
	}

	packages := make(msg.Packages, 0, len(claimLogs))
	for _, claimLog := range claimLogs {
		payload, err := hex.DecodeString(claimLog.PayLoad)
		if err != nil {
			return fmt.Errorf("decode payload error, payload=%s", claimLog.PayLoad)
		}

		pack := msg.Package{
			ChannelId: types.IbcChannelID(claimLog.ChannelId),
			Sequence:  claimLog.PackageSequence,
			Payload:   payload,
		}
		packages = append(packages, pack)
	}

	encodedPackages, err := rlp.EncodeToBytes(packages)
	if err != nil {
		return fmt.Errorf("encode packages error, err=%s", err.Error())
	}

	util.Logger.Infof("claim, chain_id=%d, seq=%d, payload=%s",
		chainId, sequence, hex.EncodeToString(encodedPackages))
	txHash, err := r.BBCExecutor.Claim(chainId, uint64(sequence), encodedPackages)
	if err != nil {
		util.Logger.Errorf("claim error: err=%s", err.Error())
		return err
	}

	r.DB.Model(model.CrossChainPackageLog{}).Where("oracle_sequence = ? and chain_id = ?", sequence, chainId).Update(map[string]interface{}{
		"status":        model.PackageStatusClaimed,
		"claim_tx_hash": txHash,
		"update_time":   time.Now().Unix(),
	})
	return nil
}

func (r *Relayer) Alert() {
	for {
		time.Sleep(common.PackageDelayAlertInterval)

		claimLog := &model.CrossChainPackageLog{}

		err := r.DB.Where("chain_id = ? and status = ?",
			r.Config.ChainConfig.BSCChainId, model.PackageStatusConfirmed).Order("oracle_sequence asc").First(&claimLog).Error
		if err != nil {
			util.Logger.Errorf("query claim log error: err=%s", err.Error())
			continue
		}

		if claimLog.Id == 0 {
			continue
		}

		if time.Now().Unix()-claimLog.UpdateTime > r.Config.AlertConfig.PackageDelayAlertThreshold {
			alertMsg := fmt.Sprintf("[%s] cross chain package was confirmed but not relayed, confiremd_time=%s, sequence=%d",
				r.Config.AlertConfig.Moniker, time.Unix(claimLog.UpdateTime, 0).String(), claimLog.OracleSequence)

			util.SendTelegramMessage(r.Config.AlertConfig.TelegramBotId, r.Config.AlertConfig.TelegramChatId, alertMsg)
		}
	}
}
