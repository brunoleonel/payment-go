package models

import (
	"time"
)

//Transaction represents the table transactions
type Transaction struct {
	TransactionID   int64      `gorm:"column:Transaction_ID;auto_increment;primary_key"`
	AccountID       int64      `gorm:"column:Account_ID"`
	OperationTypeID int64      `gorm:"column:OperationType_ID"`
	Amount          float32    `gorm:"column:Amount"`
	Balance         float32    `gorm:"column:Balance"`
	EventDate       *time.Time `gorm:"column:EventDate"`
	DueDate         *time.Time `gorm:"column:DueDate"`
}
