package repositories

import (
	"github.com/jinzhu/gorm"
)

//TransactionRepository handles the operations related to transactions on the database
type TransactionRepository interface {
}

type transactionRepository struct {
	db *gorm.DB
}

//NewTransactionRepository returns a new TransactionRepository instance
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
