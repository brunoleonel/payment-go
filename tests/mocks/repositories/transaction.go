package mocks

import (
	"time"

	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
	"github.com/brunoleonel/payment/app/repositories"
	"github.com/jinzhu/gorm"
)

type transactionRepositoryMock struct {
	db *gorm.DB
}

//NewTransactionRepositoryMock returns a new TransactionRepository instance
func NewTransactionRepositoryMock(db *gorm.DB) repositories.TransactionRepository {
	return &transactionRepositoryMock{
		db: db,
	}
}

//Create creates a new transaction
func (repository *transactionRepositoryMock) Create(model *models.Transaction) *models.Transaction {
	model.TransactionID = 1
	return model
}

//Update updates a transaction
func (repository *transactionRepositoryMock) Update(model *models.Transaction) *models.Transaction {
	return model
}

//FindPendent lists the pending transactions for an account
func (repository *transactionRepositoryMock) FindPendent(accountID int64) (result models.Transaction, err *resources.Error) {

	notFound := accountID != 1

	if notFound {
		err = &resources.Error{
			Code:    404,
			Message: "No pendent transaction found",
		}
		return
	}

	time := time.Now()

	result = models.Transaction{
		TransactionID:   1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -100,
		Balance:         -100,
		DueDate:         &time,
		EventDate:       &time,
	}

	return
}
