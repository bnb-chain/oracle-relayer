package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BlockLog struct {
	Id         int64
	Chain      string
	BlockHash  string
	ParentHash string
	Height     int64
	BlockTime  int64
	CreateTime int64
}

func (BlockLog) TableName() string {
	return "block_log"
}

func (l *BlockLog) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	return nil
}

type PackageStatus int

const (
	PackageStatusInit      PackageStatus = 0
	PackageStatusConfirmed PackageStatus = 1
	PackageStatusClaimed   PackageStatus = 2
)

type CrossChainPackageLog struct {
	Id              int64
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	PayLoad         string `gorm:"type:text"`
	TxIndex         uint

	Status       PackageStatus
	BlockHash    string
	TxHash       string
	Height       int64
	ConfirmedNum int64
	CreateTime   int64
	UpdateTime   int64
}

func (CrossChainPackageLog) TableName() string {
	return "cross_chain_package_log"
}

func InitTables(db *gorm.DB) {
	if !db.HasTable(&BlockLog{}) {
		db.CreateTable(&BlockLog{})
		db.Model(&BlockLog{}).AddUniqueIndex("idx_block_log_height", "height")
		db.Model(&BlockLog{}).AddIndex("idx_block_log_create_time", "create_time")
	}

	if !db.HasTable(&CrossChainPackageLog{}) {
		db.CreateTable(&CrossChainPackageLog{})
		db.Model(&CrossChainPackageLog{}).AddIndex("idx_package_log_channel_seq", "channel_id", "oracle_sequence")
		db.Model(&CrossChainPackageLog{}).AddIndex("idx_package_log_height", "height")
		db.Model(&CrossChainPackageLog{}).AddIndex("idx_package_log_status", "status")
	}
}
