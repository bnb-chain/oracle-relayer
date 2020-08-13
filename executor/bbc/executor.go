package bbc

import (
	"fmt"

	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"

	"github.com/binance-chain/oracle-relayer/util"
)

type Executor struct {
	config    *util.Config
	RpcClient rpc.DexClient
}

func NewExecutor(provider string, network types.ChainNetwork, cfg *util.Config) (*Executor, error) {
	rpcClient := rpc.NewRPCClient(provider, network)

	return &Executor{
		config:    cfg,
		RpcClient: rpcClient,
	}, nil
}

func getKeyManager(config *util.ChainConfig) (keys.KeyManager, error) {
	var bnbMnemonic string
	if config.BBCKeyType == util.KeyTypeAWSMnemonic {
		awsMnemonic, err := util.GetSecret(config.BBCAWSSecretName, config.BBCAWSRegion)
		if err != nil {
			return nil, err
		}
		bnbMnemonic = awsMnemonic
	} else {
		bnbMnemonic = config.BBCMnemonic
	}

	return keys.NewMnemonicKeyManager(bnbMnemonic)
}

func (e *Executor) GetAddress() types.ValAddress {
	keyManager, err := getKeyManager(e.config.ChainConfig)
	if err != nil {
		return types.ValAddress{}
	}
	return types.ValAddress(keyManager.GetAddr())
}

func (e *Executor) GetProphecy(chainId uint16, sequence int64) (*msg.Prophecy, error) {
	prop, err := e.RpcClient.GetProphecy(types.IbcChainID(chainId), sequence)
	if err != nil {
		return nil, err
	}
	return prop, err
}

func (e *Executor) Claim(chainId uint16, sequence uint64, payload []byte) (string, error) {
	keyManager, err := getKeyManager(e.config.ChainConfig)
	if err != nil {
		return "", fmt.Errorf("get key manager error, err=%s", err.Error())
	}
	e.RpcClient.SetKeyManager(keyManager)
	defer e.RpcClient.SetKeyManager(nil)

	res, err := e.RpcClient.Claim(types.IbcChainID(chainId), sequence, payload, rpc.Commit)
	if err != nil {
		return "", err
	}
	if res.Code != 0 {
		return "", fmt.Errorf("claim error, code=%d, log=%s", res.Code, res.Log)
	}
	util.Logger.Infof("claim success, tx_hash=%s", res.Hash.String())
	return res.Hash.String(), nil
}

func (e *Executor) GetCurrentSequence(chainId uint16) (int64, error) {
	sequence, err := e.RpcClient.GetCurrentOracleSequence(types.IbcChainID(chainId))
	if err != nil {
		return 0, err
	}
	return sequence, nil
}
