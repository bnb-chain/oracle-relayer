package bsc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"

	types2 "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/binance-chain/oracle-relayer/model"
)

var (
	BindSuccessEventHash                          = common.HexToHash("0x8005b9354dd0ca4c5593805bcd00ea12b5fce8a2cc9bc15252f50fb2d17c09d2")
	BindRejectedEventHash                         = common.HexToHash("0x341e20b0b6b62cb3990e2d1f8bcb0a15e7d7fd446355a7be807face162285254")
	BindTimeoutEventHash                          = common.HexToHash("0x4781c2d0a33124fb32083581f5b48c93a59b71fd567ce2d4a56c89196baa2ccd")
	BindInvalidParameterEventHash                 = common.HexToHash("0x2117f993c9cc877c531b4e6bd55d822cb48b529fd003c80e5bd6c27b7c1c1702")
	TransferInFailureTimeoutEventHash             = common.HexToHash("0x8090e98e190cb0b05412d5c1a8cd5ee9af5d40da935335cef5d4179c7da63d79")
	TransferInFailureInsufficientBalanceEventHash = common.HexToHash("0x1de400dfa3e72ba83f12c6f1d8b9b85dc3d2aedc6eacc27b481267826aec7422")
	TransferInFailureUnboundTokenEventHash        = common.HexToHash("0x055f2adbd109a4e99b3821af55571cccb4981551d10e3846b21574d348572a59")
	TransferInFailureUnknownReasonEventHash       = common.HexToHash("0xcb6ddd4a252f58c1ff32f31fbb529dc35e8f6a81908f6211bbe7dfa94ef52f1f")
	TransferOutEventHash                          = common.HexToHash("0x5bd451c53ab05abd9855ceb52a469590655af1d732a4cfd67f1f9b53d74dc613")
	BatchTransferOutEventHash                     = common.HexToHash("0x00a18f0343865824d1375c23f5dd79fdf32a12f50400ef2591e52276f8378e31")
	BatchTransferOutAddrsEventHash                = common.HexToHash("0x8740bbd4e1a2505bf2908481adbf1056fb52f762152b702f6c65468f63c55cf8")
	ValidatorFelonyEventHash                      = common.HexToHash("0x7e770310e43f85c3dca97460dbe1484068514437298ff349e6052595a6ffbdb7")
)

const (
	BindSuccessEventName                          = "LogBindSuccess"
	BindRejectedEventName                         = "LogBindRejected"
	BindTimeoutEventName                          = "LogBindTimeout"
	BindInvalidParameterEventName                 = "LogBindInvalidParameter"
	TransferOutEventName                          = "LogTransferOut"
	BatchTransferOutEventName                     = "LogBatchTransferOut"
	BatchTransferOutAddrsEventName                = "LogBatchTransferOutAddrs"
	TransferInFailureTimeoutEventName             = "LogTransferInFailureTimeout"
	TransferInFailureInsufficientBalanceEventName = "LogTransferInFailureInsufficientBalance"
	TransferInFailureUnboundTokenEventName        = "LogTransferInFailureUnboundToken"
	TransferInFailureUnknownReasonEventName       = "LogTransferInFailureUnknownReason"
	ValidatorFelonyEventName                      = "validatorFelony"
)

type ContractEvent interface {
	ToTxLog(log *types.Log) interface{}
}

type BindSuccessEvent struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
	TotalSupply     *big.Int
	PeggyAmount     *big.Int
	Decimals        *big.Int
}

func (ev BindSuccessEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.UpdateBindClaim{
		Status:          msg.BindStatusSuccess,
		Symbol:          string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		ContractAddress: msg.SmartChainAddress(ev.ContractAddr),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeUpdateBind),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog
}

type BindRejectedEvent struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}

func (ev BindRejectedEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.UpdateBindClaim{
		Status:          msg.BindStatusRejected,
		Symbol:          string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		ContractAddress: msg.SmartChainAddress(ev.ContractAddr),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeUpdateBind),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}

	return &claimLog
}

type BindTimeoutEvent struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}

func (ev BindTimeoutEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.UpdateBindClaim{
		Status:          msg.BindStatusTimeout,
		Symbol:          string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		ContractAddress: msg.SmartChainAddress(ev.ContractAddr),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeUpdateBind),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}

	return &claimLog
}

type BindInvalidParameterEvent struct {
	Sequence        *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}

func (ev BindInvalidParameterEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.UpdateBindClaim{
		Status:          msg.BindStatusInvalidParameter,
		Symbol:          string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		ContractAddress: msg.SmartChainAddress(ev.ContractAddr),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeUpdateBind),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}

	return &claimLog
}

func ParseBindSuccessEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev BindSuccessEvent

	err := abi.Unpack(&ev, BindSuccessEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func ParseBindRejectedEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev BindRejectedEvent

	err := abi.Unpack(&ev, BindRejectedEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func ParseBindTimeoutEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev BindTimeoutEvent

	err := abi.Unpack(&ev, BindTimeoutEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func ParseBindInvalidParameterEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev BindInvalidParameterEvent

	err := abi.Unpack(&ev, BindInvalidParameterEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

type TransferInTimeoutEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
	ExpireTime      *big.Int
}

func ParseTransferInTimeoutEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferInTimeoutEvent

	err := abi.Unpack(&ev, TransferInFailureTimeoutEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (ev TransferInTimeoutEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.TransferOutRefundClaim{
		RefundAddress: types2.AccAddress(ev.RefundAddr[:]),
		Amount: types2.Coin{
			Denom:  string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
			Amount: ev.Amount.Int64(),
		},
		RefundReason: msg.Timeout,
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferOutRefund),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog
}

type TransferInInsufficientBalanceEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
	ActualBalance   *big.Int
}

func ParseTransferInInsufficientBalanceEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferInInsufficientBalanceEvent

	err := abi.Unpack(&ev, TransferInFailureInsufficientBalanceEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (ev TransferInInsufficientBalanceEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.TransferOutRefundClaim{
		RefundAddress: types2.AccAddress(ev.RefundAddr[:]),
		Amount: types2.Coin{
			Denom:  string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
			Amount: ev.Amount.Int64(),
		},
		RefundReason: msg.InsufficientBalance,
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferOutRefund),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog
}

type TransferInUnboundTokenEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}

func ParseTransferInUnboundedTokenEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferInUnboundTokenEvent

	err := abi.Unpack(&ev, TransferInFailureUnboundTokenEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (ev TransferInUnboundTokenEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.TransferOutRefundClaim{
		RefundAddress: types2.AccAddress(ev.RefundAddr[:]),
		Amount: types2.Coin{
			Denom:  string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
			Amount: ev.Amount.Int64(),
		},
		RefundReason: msg.UnboundToken,
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferOutRefund),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog
}

type TransferInUnknownReasonEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Bep2TokenAmount *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
}

func ParseTransferInUnknownReasonEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferInUnknownReasonEvent

	err := abi.Unpack(&ev, TransferInFailureUnknownReasonEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (ev TransferInUnknownReasonEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.TransferOutRefundClaim{
		RefundAddress: types2.AccAddress(ev.RefundAddr[:]),
		Amount: types2.Coin{
			Denom:  string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
			Amount: ev.Bep2TokenAmount.Int64(),
		},
		RefundReason: msg.Unknown,
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferOutRefund),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog
}

type TransferOutEvent struct {
	Sequence        *big.Int
	RefundAddr      common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
	ExpireTime      *big.Int
	RelayFee        *big.Int
}

func (ev TransferOutEvent) ToTxLog(log *types.Log) interface{} {
	claim := msg.TransferInClaim{
		ContractAddress:   msg.SmartChainAddress(ev.ContractAddr),
		RefundAddresses:   []msg.SmartChainAddress{msg.SmartChainAddress(ev.RefundAddr)},
		ReceiverAddresses: []types2.AccAddress{types2.AccAddress(ev.Recipient[:])},
		Amounts:           []int64{ev.Amount.Int64()},
		Symbol:            string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		RelayFee: types2.Coin{
			Denom:  msg.NativeToken,
			Amount: ev.RelayFee.Int64(),
		},
		ExpireTime: ev.ExpireTime.Int64(),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferIn),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}

	return &claimLog
}

func ParseTransferOutEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferOutEvent

	err := abi.Unpack(&ev, TransferOutEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

type BatchTransferOutEvent struct {
	Sequence        *big.Int
	Amounts         []*big.Int
	ContractAddr    common.Address
	Bep2TokenSymbol common.Hash
	ExpireTime      *big.Int
	RelayFee        *big.Int
}

func (ev BatchTransferOutEvent) ToTxLog(log *types.Log) interface{} {
	return nil
}

func ParseBatchTransferOutEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev BatchTransferOutEvent

	err := abi.Unpack(&ev, BatchTransferOutEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return ev, nil
}

type BatchTransferOutAddrsEvent struct {
	Sequence       *big.Int
	RecipientAddrs []common.Address
	RefundAddrs    []common.Address
}

func ParseBatchTransferOutAddrsEvent(abi *abi.ABI, log *types.Log) (*BatchTransferOutAddrsEvent, error) {
	var ev BatchTransferOutAddrsEvent

	err := abi.Unpack(&ev, BatchTransferOutAddrsEventName, log.Data)
	if err != nil {
		return nil, err
	}

	return &ev, nil
}

func ParseBatchTransferOutEventToTxLog(abi *abi.ABI, log *types.Log, batchTransferOutEvents map[int64]*BatchTransferOutAddrsEvent) (interface{}, error) {
	var ev BatchTransferOutEvent

	err := abi.Unpack(&ev, BatchTransferOutEventName, log.Data)
	if err != nil {
		return nil, err
	}

	amountsInt64 := make([]int64, 0, len(ev.Amounts))
	for _, amount := range ev.Amounts {
		amountsInt64 = append(amountsInt64, amount.Int64())
	}

	batchTransferOutAddrsEvent := batchTransferOutEvents[ev.Sequence.Int64()]
	if batchTransferOutEvents == nil {
		return nil, fmt.Errorf("transfer out addrs event does not exist")
	}

	refundAddressesConverted := make([]msg.SmartChainAddress, 0, len(batchTransferOutAddrsEvent.RefundAddrs))
	for _, addr := range batchTransferOutAddrsEvent.RefundAddrs {
		tmpAddr := msg.SmartChainAddress{}
		copy(tmpAddr[:], addr[:])
		refundAddressesConverted = append(refundAddressesConverted, tmpAddr)
	}

	receiverAddressesConverted := make([]types2.AccAddress, 0, len(batchTransferOutAddrsEvent.RecipientAddrs))
	for _, addr := range batchTransferOutAddrsEvent.RecipientAddrs {
		tmpAddr := make([]byte, types2.AddrLen, types2.AddrLen)
		copy(tmpAddr, addr[:])
		receiverAddressesConverted = append(receiverAddressesConverted, tmpAddr)
	}

	claim := msg.TransferInClaim{
		ContractAddress:   msg.SmartChainAddress(ev.ContractAddr),
		RefundAddresses:   refundAddressesConverted,
		ReceiverAddresses: receiverAddressesConverted,
		Amounts:           amountsInt64,
		Symbol:            string(bytes.Trim(ev.Bep2TokenSymbol[:], "\x00")),
		RelayFee: types2.Coin{
			Denom:  msg.NativeToken,
			Amount: ev.RelayFee.Int64(),
		},
		ExpireTime: ev.ExpireTime.Int64(),
	}

	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeTransferIn),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog, nil
}

type ValidatorFelonyEvent struct {
	Sequence  *big.Int
	Validator common.Address
	Amount    *big.Int
}

func ParseValidatorFelonyEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev ValidatorFelonyEvent

	err := abi.Unpack(&ev, ValidatorFelonyEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Sequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
	ev.Validator = common.BytesToAddress(log.Topics[2].Bytes())

	return ev, nil
}

func (ev ValidatorFelonyEvent) ToTxLog(log *types.Log) interface{} {
	return nil
}

func ParseValidatorFelonyEventToTxLog(abi *abi.ABI, header *types.Header, log *types.Log, chainId string) (interface{}, error) {
	var ev ValidatorFelonyEvent

	err := abi.Unpack(&ev, ValidatorFelonyEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.Sequence = big.NewInt(0).SetBytes(log.Topics[1].Bytes())
	ev.Validator = common.BytesToAddress(log.Topics[2].Bytes())

	claim := msg.SideDowntimeSlashClaim{
		SideConsAddr:  ev.Validator[:],
		SideHeight:    header.Number.Int64(),
		SideChainId:   chainId,
		SideTimestamp: int64(header.Time),
	}
	claimBz, err := json.Marshal(claim)
	if err != nil {
		panic(err)
	}

	claimLog := model.ClaimLog{
		Sequence:  ev.Sequence.Int64(),
		ClaimType: int8(msg.ClaimTypeDowntimeSlash),
		Claim:     string(claimBz),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return &claimLog, nil
}

func ParseTokenHubEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], BindSuccessEventHash[:]) {
		return ParseBindSuccessEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], BindRejectedEventHash[:]) {
		return ParseBindRejectedEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], BindTimeoutEventHash[:]) {
		return ParseBindTimeoutEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], BindInvalidParameterEventHash[:]) {
		return ParseBindInvalidParameterEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], TransferInFailureTimeoutEventHash[:]) {
		return ParseTransferInTimeoutEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], TransferInFailureInsufficientBalanceEventHash[:]) {
		return ParseTransferInInsufficientBalanceEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], TransferInFailureUnboundTokenEventHash[:]) {
		return ParseTransferInUnboundedTokenEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], TransferInFailureUnknownReasonEventHash[:]) {
		return ParseTransferInUnknownReasonEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], TransferOutEventHash[:]) {
		return ParseTransferOutEvent(abi, log)
	} else if bytes.Equal(log.Topics[0][:], BatchTransferOutEventHash[:]) {
		return ParseBatchTransferOutEvent(abi, log)
	}
	return nil, nil
}

func ParseValidatorSetEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], ValidatorFelonyEventHash[:]) {
		return ParseValidatorFelonyEvent(abi, log)
	}
	return nil, nil
}
