package repositories

import (
	"github.com/brunoleonel/payment/app/models"
	"github.com/jinzhu/gorm"
)

//TransactionRepository handles the operations related to transactions on the database
type TransactionRepository interface {
	Create(model *models.Transaction) *models.Transaction
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

//Create creates a new transaction
func (repository *transactionRepository) Create(model *models.Transaction) *models.Transaction {
	repository.db.Create(model)
	return model
}
