package models

//OperationType represents the table operations
type OperationType struct {
	OperationTypeID int64         `gorm:"column:OperationType_ID;auto_increment;primary_key"`
	Description     string        `gorm:"column:Description"`
	ChargeOrder     int8          `gorm:"column:ChargeOrder"`
	Transactions    []Transaction `gorm:"foreignkey:OperationType_ID"`
}
