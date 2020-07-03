package observer

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/require"

	"github.com/binance-chain/oracle-relayer/common"
	"github.com/binance-chain/oracle-relayer/executor/mock"
	"github.com/binance-chain/oracle-relayer/model"
	"github.com/binance-chain/oracle-relayer/util"
)

func TestObserver_fetchBlock_error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bscExecutor := mock.NewMockBscExecutor(ctrl)
	bscExecutor.EXPECT().GetBlockAndPackages(gomock.Any()).AnyTimes().Return(nil, errors.New("error"))

	ob := NewObserver(db, config, bscExecutor)
	err = ob.fetchBlock(1, 2, "1")
	require.NotNil(t, err, "error should not be nil")

	require.Contains(t, err.Error(), "get block info error")
}

func TestObserver_fetchBlock_error_wrongParentHash(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bscExecutor := mock.NewMockBscExecutor(ctrl)
	bscExecutor.EXPECT().GetBlockAndPackages(gomock.Any()).AnyTimes().Return(
		&common.BlockAndPackageLogs{
			Height:          3,
			BlockHash:       "3",
			ParentBlockHash: "2_1",
		}, nil)

	ob := NewObserver(db, config, bscExecutor)

	blockLog1 := &model.BlockLog{
		Height:     1,
		BlockHash:  "1",
		ParentHash: "0",
		BlockTime:  0,
	}
	db.Create(blockLog1)

	blockLog2 := &model.BlockLog{
		Height:     2,
		BlockHash:  "2",
		ParentHash: "1",
		BlockTime:  0,
	}
	db.Create(blockLog2)

	packageLog2 := &model.CrossChainPackageLog{
		ChainId:         96,
		OracleSequence:  1,
		PackageSequence: 1,
		ChannelId:       2,
		Height:          2,
		TxHash:          "tx_hash",
	}
	db.Create(packageLog2)

	err = ob.fetchBlock(2, 3, "2")
	require.Nil(t, err, "error should be nil")

	deletedBlockLog := &model.BlockLog{}
	err = db.Where("height = ?", blockLog2.Height).First(&deletedBlockLog).Error

	require.Equal(t, err, gorm.ErrRecordNotFound, "error should be ErrRecordNotFound")

	deletedPackage := &model.CrossChainPackageLog{}
	err = db.Where("height = ?", packageLog2.Height).First(&deletedPackage).Error
	require.Equal(t, err, gorm.ErrRecordNotFound, "error should be ErrRecordNotFound")
}

func TestObserver_fetchBlock_error_rightParentHash(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := util.GetTestConfig()
	db, err := util.PrepareDB(config)
	require.Nil(t, err, "create db error")

	bscExecutor := mock.NewMockBscExecutor(ctrl)
	bscExecutor.EXPECT().GetBlockAndPackages(gomock.Any()).AnyTimes().Return(
		&common.BlockAndPackageLogs{
			Height:          3,
			BlockHash:       "3",
			ParentBlockHash: "2",
			Packages: []interface{}{
				&model.CrossChainPackageLog{
					ChainId:         96,
					OracleSequence:  2,
					PackageSequence: 2,
					ChannelId:       2,
					Height:          3,
					TxHash:          "tx_hash_1",
				},
				&model.CrossChainPackageLog{
					ChainId:         96,
					OracleSequence:  2,
					PackageSequence: 3,
					ChannelId:       2,
					Height:          3,
					TxHash:          "tx_hash_2",
				},
			},
		}, nil)

	ob := NewObserver(db, config, bscExecutor)

	blockLog1 := &model.BlockLog{
		Height:     1,
		BlockHash:  "1",
		ParentHash: "0",
		BlockTime:  0,
	}
	db.Create(blockLog1)

	blockLog2 := &model.BlockLog{
		Height:     2,
		BlockHash:  "2",
		ParentHash: "1",
		BlockTime:  0,
	}
	db.Create(blockLog2)

	err = ob.fetchBlock(2, 3, "2")
	require.Nil(t, err, "error should be nil")

	newBlockLog := &model.BlockLog{}
	err = db.Where("height = ?", 3).First(&newBlockLog).Error

	require.Nil(t, err, "error should be nil")
	require.Equal(t, newBlockLog.Height, int64(3), "height should be 3")

	newPackages := make([]*model.CrossChainPackageLog, 0)
	err = db.Where("height = ?", 3).Find(&newPackages).Error

	require.Nil(t, err, "error should be nil")
	require.Equal(t, len(newPackages), 2, "length of packages should be 2")
}
