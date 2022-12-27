package db

import (
	"time"

	"gorm.io/datatypes"
)

type BlockChainLog struct {
	Address     string `gorm:"column:address"`
	BlockHash   string `gorm:"column:block_hash"`
	BlockNumber string `gorm:"column:block_number"`
	Data        string `gorm:"column:data"`
	GasPrice    string `gorm:"column:gas_price"`
	GasUsed     string `gorm:"column:gas_used"`
	LogIndex    string `gorm:"column:log_index"`
	TimeStamp   string `gorm:"column:time_stamp"`
	// Topics      []string `gorm:"column:topics"`
	Topics datatypes.JSON `gorm:"column:topics"`
	// ID     uint   `gorm:"primary_key"`
	TransactionHash  string `gorm:"column:transaction_hash;primary_key"`
	TransactionIndex string `gorm:"column:transaction_index"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
