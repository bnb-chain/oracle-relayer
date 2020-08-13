package bsc

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/binance-chain/oracle-relayer/model"
)

var (
	CrossChainPackageEventName = "crossChainPackage"

	CrossChainPackageEventHash = common.HexToHash("0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe")
)

type ContractEvent interface {
	ToTxLog(log *types.Log) interface{}
}

type CrossChainPackageEvent struct {
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
}

func (ev *CrossChainPackageEvent) ToTxLog(log *types.Log) interface{} {
	pack := &model.CrossChainPackageLog{
		ChainId:         ev.ChainId,
		OracleSequence:  ev.OracleSequence,
		PackageSequence: ev.PackageSequence,
		ChannelId:       ev.ChannelId,
		PayLoad:         hex.EncodeToString(ev.Payload),
		BlockHash:       log.BlockHash.Hex(),
		TxHash:          log.TxHash.String(),
		TxIndex:         log.TxIndex,
		Height:          int64(log.BlockNumber),
	}
	return pack
}

func ParseCrossChainPackageEvent(abi *abi.ABI, log *types.Log) (*CrossChainPackageEvent, error) {
	var ev CrossChainPackageEvent

	err := abi.Unpack(&ev, CrossChainPackageEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.OracleSequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes()).Uint64()
	ev.PackageSequence = big.NewInt(0).SetBytes(log.Topics[2].Bytes()).Uint64()
	ev.ChannelId = uint8(big.NewInt(0).SetBytes(log.Topics[3].Bytes()).Uint64())

	return &ev, nil
}
