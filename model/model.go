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

type TxStatus int

const (
	TxStatusInit      TxStatus = 0
	TxStatusConfirmed TxStatus = 1
	TxStatusClaimed   TxStatus = 2
)

type ClaimLog struct {
	Id        int64
	Sequence  int64
	ClaimType int8
	Claim     string `gorm:"type:text"`

	Status       TxStatus
	BlockHash    string
	TxHash       string
	Height       int64
	ConfirmedNum int64
	CreateTime   int64
	UpdateTime   int64
}

func (ClaimLog) TableName() string {
	return "claim_log"
}

func InitTables(db *gorm.DB) {
	if !db.HasTable(&BlockLog{}) {
		db.CreateTable(&BlockLog{})
		db.Model(&BlockLog{}).AddUniqueIndex("idx_block_log_height", "height")
		db.Model(&BlockLog{}).AddIndex("idx_block_log_create_time", "create_time")
	}

	if !db.HasTable(&ClaimLog{}) {
		db.CreateTable(&ClaimLog{})
	}
}
