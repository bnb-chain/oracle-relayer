package executor

import (
	"github.com/bnb-chain/go-sdk/common/types"
	"github.com/bnb-chain/go-sdk/types/msg"

	"github.com/binance-chain/oracle-relayer/common"
)

type BbcExecutor interface {
	GetAddress() types.ValAddress
	GetCurrentSequence(chainId uint16) (int64, error)
	GetProphecy(chainId uint16, sequence int64) (*msg.Prophecy, error)

	Claim(chainId uint16, sequence uint64, payload []byte) (string, error)
}

type BscExecutor interface {
	GetBlockAndPackages(height int64) (*common.BlockAndPackageLogs, error)
}
