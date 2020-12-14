package bsc

import (
	"context"
	"math/big"
	"math/rand"
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

	CrossChainAbi abi.ABI
	Clients       []*ethclient.Client

	crossChainContractAddress ethcmm.Address
}

// NewExecutor returns the bsc executor instance
func NewExecutor(providers []string, config *util.Config) *Executor {
	crossChainAbi, err := abi.JSON(strings.NewReader(abi2.CrossChainABI))
	if err != nil {
		panic("marshal abi error")
	}

	clients := initClients(providers)

	return &Executor{
		Config:        config,
		CrossChainAbi: crossChainAbi,
		Clients:       clients,

		crossChainContractAddress: config.ChainConfig.BSCCrossChainContractAddress,
	}
}

func initClients(providers []string) []*ethclient.Client {
	clients := make([]*ethclient.Client, 0)

	for _, provider := range providers {
		client, err := ethclient.Dial(provider)
		if err != nil {
			panic("new eth client error")
		}
		clients = append(clients, client)
	}

	return clients
}

func (e *Executor) getClient() *ethclient.Client {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	idx := r.Intn(len(e.Clients))
	return e.Clients[idx]
}

// GetBlockAndPackages returns the block and cross-chain packages of the given height
func (e *Executor) GetBlockAndPackages(height int64) (*common.BlockAndPackageLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := e.getClient()
	header, err := client.HeaderByNumber(ctxWithTimeout, big.NewInt(height))
	if err != nil {
		return nil, err
	}

	packageLogs, err := e.GetLogs(client, header)
	if err != nil {
		return nil, err
	}

	return &common.BlockAndPackageLogs{
		Height:          height,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		Packages:        packageLogs,
	}, nil
}

// GetLogs return the cross-chain packages of the given height
func (e *Executor) GetLogs(client *ethclient.Client, header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{CrossChainPackageEventHash}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcmm.Address{e.crossChainContractAddress},
	})
	if err != nil {
		return nil, err
	}

	packageModels := make([]interface{}, 0, len(logs))

	for _, log := range logs {
		util.Logger.Infof("get log: %d, %s, %s", log.BlockNumber, log.Topics[0].String(), log.TxHash.String())

		event, err := ParseCrossChainPackageEvent(&e.CrossChainAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}

		if event == nil {
			continue
		}

		packageModel := event.ToTxLog(&log)
		packageModels = append(packageModels, packageModel)
	}
	return packageModels, nil
}
