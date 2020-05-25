package relayer

import (
	"fmt"
	"time"

	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/jinzhu/gorm"

	"github.com/binance-chain/oracle-relayer/common"
	"github.com/binance-chain/oracle-relayer/executor/bbc"
	"github.com/binance-chain/oracle-relayer/model"
	"github.com/binance-chain/oracle-relayer/util"
)

type Relayer struct {
	DB          *gorm.DB
	BBCExecutor *bbc.Executor
}

func NewRelayer(db *gorm.DB, bbcExecutor *bbc.Executor) *Relayer {
	return &Relayer{
		DB:          db,
		BBCExecutor: bbcExecutor,
	}
}

func (r *Relayer) Main() {
	claimTypes := []msg.ClaimType{
		msg.ClaimTypeUpdateBind,
		msg.ClaimTypeTransferOutRefund,
		msg.ClaimTypeTransferIn,
		msg.ClaimTypeDowntimeSlash,
	}

	for _, claimType := range claimTypes {
		go func(claimType msg.ClaimType) {
			for {
				err := r.process(claimType)
				if err != nil {
					time.Sleep(common.RelayerInterval)
				}
			}
		}(claimType)
	}
}

func (r *Relayer) process(claimType msg.ClaimType) error {
	sequence, err := r.BBCExecutor.GetCurrentSequence(claimType)
	if err != nil {
		util.Logger.Errorf("get current sequence error: claim_type=%s, err=%s",
			msg.ClaimTypeToString(claimType), err.Error())
		return err
	}

	util.Logger.Infof("current sequence, claim_type=%s, seq=%d", msg.ClaimTypeToString(claimType), sequence)

	prophecy, err := r.BBCExecutor.GetProphecy(claimType, sequence)
	if err != nil {
		util.Logger.Errorf("get prophecy error: err=%s", err.Error())
		return err
	}

	validatorAddress := r.BBCExecutor.GetAddress()
	if prophecy != nil && prophecy.ValidatorClaims != nil && prophecy.ValidatorClaims[validatorAddress.String()] != "" {
		return fmt.Errorf("already claimed")
	}

	claimLog := &model.ClaimLog{}
	err = r.DB.Where("sequence = ? and claim_type = ? and status = ?", sequence, int8(claimType),
		model.TxStatusConfirmed).First(&claimLog).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			util.Logger.Errorf("query claim log error: err=%s", err.Error())
		}
		return err
	}

	util.Logger.Infof("claim, type=%s, seq=%d, claim=%s",
		msg.ClaimTypeToString(msg.ClaimType(claimLog.ClaimType)), claimLog.Sequence, claimLog.Claim)
	err = r.BBCExecutor.Claim(msg.ClaimType(claimLog.ClaimType), claimLog.Sequence, claimLog.Claim)
	if err != nil {
		util.Logger.Errorf("claim error: err=%s", err.Error())
		return err
	}

	r.DB.Model(claimLog).Update(map[string]interface{}{
		"status":      model.TxStatusClaimed,
		"update_time": time.Now().Unix(),
	})
	return nil
}
