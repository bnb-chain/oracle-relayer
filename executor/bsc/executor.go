package bsc

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/binance-chain/oracle-relayer/common"
	abi2 "github.com/binance-chain/oracle-relayer/executor/bsc/abi"
	"github.com/binance-chain/oracle-relayer/util"
)

type Executor struct {
	Config *util.Config

	TokenHubAbi     abi.ABI
	ValidatorSetAbi abi.ABI

	tokenHubContractAddress     ethcmm.Address
	validatorSetContractAddress ethcmm.Address
	Client                      *ethclient.Client
}

func NewExecutor(provider string, config *util.Config) *Executor {
	tokenHubAbi, err := abi.JSON(strings.NewReader(abi2.TokenHubABI))
	if err != nil {
		panic("marshal abi error")
	}
	validatorSetAbi, err := abi.JSON(strings.NewReader(abi2.ValidatorSetABI))
	if err != nil {
		panic("marshal abi error")
	}

	client, err := ethclient.Dial(provider)
	if err != nil {
		panic("new eth client error")
	}

	return &Executor{
		Config:                      config,
		tokenHubContractAddress:     config.ChainConfig.BSCTokenHubContractAddress,
		validatorSetContractAddress: config.ChainConfig.BSCValidatorSetContractAddress,
		TokenHubAbi:                 tokenHubAbi,
		ValidatorSetAbi:             validatorSetAbi,
		Client:                      client,
	}
}

func (e *Executor) GetBlockAndTxs(height int64) (*common.BlockAndTxLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	header, err := e.Client.HeaderByNumber(ctxWithTimeout, big.NewInt(height))
	if err != nil {
		return nil, err
	}

	txLogs, err := e.GetLogs(header)
	if err != nil {
		return nil, err
	}

	return &common.BlockAndTxLogs{
		Height:          height,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		TxLogs:          txLogs,
	}, nil
}

func (e *Executor) GetLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{
		BindSuccessEventHash, BindRejectedEventHash, BindTimeoutEventHash, BindInvalidParameterEventHash,
		TransferInFailureTimeoutEventHash, TransferInFailureInsufficientBalanceEventHash, TransferInFailureUnboundTokenEventHash,
		TransferInFailureUnknownReasonEventHash, TransferOutEventHash, BatchTransferOutEventHash, BatchTransferOutAddrsEventHash,
		ValidatorFelonyEventHash,
	}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcmm.Address{e.tokenHubContractAddress, e.validatorSetContractAddress},
	})
	if err != nil {
		return nil, err
	}

	models := make([]interface{}, 0, len(logs))

	// get batch transfer out address events first
	transferOutAddrsEvents := make(map[int64]*BatchTransferOutAddrsEvent)
	for _, log := range logs {
		if bytes.Equal(log.Topics[0][:], BatchTransferOutAddrsEventHash[:]) {
			event, err := ParseBatchTransferOutAddrsEvent(&e.TokenHubAbi, &log)
			if err != nil {
				util.Logger.Errorf("parse event log error, er=%s", err.Error())
				continue
			}
			if event == nil {
				continue
			}
			transferOutAddrsEvents[event.Sequence.Int64()] = event
		}
	}

	for _, log := range logs {
		util.Logger.Info("get log: %d, %s, %s", log.BlockNumber, hex.EncodeToString(log.Topics[0][:]), log.TxHash.String())
		event, err := e.parseEvent(&log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		// skip batch transfer out addrs events
		if bytes.Equal(log.Topics[0][:], BatchTransferOutAddrsEventHash[:]) {
			continue
		}

		var txLog interface{}
		if bytes.Equal(log.Topics[0][:], BatchTransferOutEventHash[:]) {
			txLog, err = ParseBatchTransferOutEventToTxLog(&e.TokenHubAbi, &log, transferOutAddrsEvents)
			if err != nil {
				return nil, err
			}
		} else if bytes.Equal(log.Topics[0][:], ValidatorFelonyEventHash[:]) {
			txLog, err = ParseValidatorFelonyEventToTxLog(&e.ValidatorSetAbi, header, &log, e.Config.ChainConfig.BSCChainId)
		} else {
			txLog = event.ToTxLog(&log)
		}

		models = append(models, txLog)
	}
	return models, nil
}

func (e *Executor) parseEvent(log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Address[:], e.validatorSetContractAddress[:]) {
		return ParseValidatorSetEvent(&e.ValidatorSetAbi, log)
	} else {
		return ParseTokenHubEvent(&e.TokenHubAbi, log)
	}
}
