package models

//Account represents the table accounts
type Account struct {
	AccountID                int64         `gorm:"column:Account_ID;auto_increment;primary_key"`
	AvailableCreditLimit     float32       `gorm:"column:AvailableCreditLimit"`
	AvailableWithdrawalLimit float32       `gorm:"column:AvailableWithdrawalLimit"`
	Transactions             []Transaction `gorm:"foreignkey:Account_ID"`
}
