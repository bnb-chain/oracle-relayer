package relayer

import (
	"errors"
	"testing"

	"github.com/bnb-chain/go-sdk/types/msg"

	"github.com/bnb-chain/go-sdk/common/types"

	"github.com/binance-chain/oracle-relayer/model"

	"github.com/golang/mock/gomock"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/require"

	"github.com/binance-chain/oracle-relayer/executor/mock"
	"github.com/binance-chain/oracle-relayer/util"
)

func TestRelayer_process_getSequenceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(0), errors.New("get sequence error"))

	relayer := NewRelayer(db, bbcExecutor, config)
	err = relayer.process(96)
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "get sequence error")
}

func TestRelayer_process_emptyLogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(1), nil)

	relayer := NewRelayer(db, bbcExecutor, config)
	err = relayer.process(96)
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "no packages found")
}

func TestRelayer_process_getProphecyError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(1), nil)
	bbcExecutor.EXPECT().GetProphecy(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, errors.New("get prophecy error"))

	relayer := NewRelayer(db, bbcExecutor, config)

	packageLog := &model.CrossChainPackageLog{
		ChainId:         96,
		OracleSequence:  1,
		PackageSequence: 1,
		ChannelId:       2,
		Height:          2,
		Status:          1,
		TxHash:          "tx_hash",
	}
	db.Create(packageLog)

	err = relayer.process(96)
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "get prophecy error")
}

func TestRelayer_process_alreadyClaimed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	validatorAddr, err := types.AccAddressFromBech32("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	require.Nil(t, err, "error should be nil")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(1), nil)
	bbcExecutor.EXPECT().GetProphecy(gomock.Any(), gomock.Any()).AnyTimes().Return(&msg.Prophecy{
		ID:              "1",
		Status:          msg.Status{},
		ClaimValidators: nil,
		ValidatorClaims: map[string]string{
			types.ValAddress(validatorAddr).String(): "claim",
		},
	}, nil)
	bbcExecutor.EXPECT().GetAddress().AnyTimes().Return(types.ValAddress(validatorAddr))

	relayer := NewRelayer(db, bbcExecutor, config)

	packageLog := &model.CrossChainPackageLog{
		ChainId:         96,
		OracleSequence:  1,
		PackageSequence: 1,
		ChannelId:       2,
		Height:          2,
		Status:          1,
		TxHash:          "tx_hash",
	}
	db.Create(packageLog)

	err = relayer.process(96)
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "already claimed")
}

func TestRelayer_process_claimError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	validatorAddr, err := types.AccAddressFromBech32("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	require.Nil(t, err, "error should be nil")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(1), nil)
	bbcExecutor.EXPECT().GetProphecy(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	bbcExecutor.EXPECT().GetAddress().AnyTimes().Return(types.ValAddress(validatorAddr))
	bbcExecutor.EXPECT().Claim(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("", errors.New("claim error"))

	relayer := NewRelayer(db, bbcExecutor, config)

	packageLog := &model.CrossChainPackageLog{
		ChainId:         96,
		OracleSequence:  1,
		PackageSequence: 1,
		ChannelId:       2,
		Height:          2,
		Status:          1,
		TxHash:          "tx_hash",
	}
	db.Create(packageLog)

	err = relayer.process(96)
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "claim error")
}

func TestRelayer_process_claimSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	validatorAddr, err := types.AccAddressFromBech32("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	require.Nil(t, err, "error should be nil")

	bbcExecutor := mock.NewMockBbcExecutor(ctrl)
	bbcExecutor.EXPECT().GetCurrentSequence(gomock.Any()).AnyTimes().Return(int64(1), nil)
	bbcExecutor.EXPECT().GetProphecy(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)
	bbcExecutor.EXPECT().GetAddress().AnyTimes().Return(types.ValAddress(validatorAddr))
	bbcExecutor.EXPECT().Claim(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return("tx_hash", nil)

	relayer := NewRelayer(db, bbcExecutor, config)

	packageLog := &model.CrossChainPackageLog{
		ChainId:         96,
		OracleSequence:  1,
		PackageSequence: 1,
		ChannelId:       2,
		Height:          2,
		Status:          1,
		TxHash:          "tx_hash",
	}
	db.Create(packageLog)

	err = relayer.process(96)
	require.Nil(t, err, "error should be nil")

	newPackage := &model.CrossChainPackageLog{}
	err = db.Where("height = ?", 2).First(newPackage).Error
	require.Nil(t, err, "error should be nil")

	require.Equal(t, newPackage.TxHash, "tx_hash")
	require.Equal(t, newPackage.Status, model.PackageStatusClaimed)
}
