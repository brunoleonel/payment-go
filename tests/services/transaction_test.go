package test

import (
	"testing"

	"github.com/brunoleonel/payment/app/http/resources"

	"github.com/brunoleonel/payment/app/services"

	mocks "github.com/brunoleonel/payment/tests/mocks/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateTransaction(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	accountRepository := mocks.NewAccountRepositoryMock(db)
	transactionRepository := mocks.NewTransactionRepositoryMock(db)
	accountService := services.NewAccountService(accountRepository)
	transactionService := services.NewTransactionService(transactionRepository, accountService)

	resource := &resources.Transaction{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          1,
	}

	result, _ := transactionService.Create(resource)

	assert.NotNil(result)
	assert.Equal(result.AccountID, int64(1))
	assert.Equal(result.OperationTypeID, int64(1))
	assert.Equal(result.Amount, float32(1))
}

func TestShouldntCreateTransactionByAccountNotFound(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	accountRepository := mocks.NewAccountRepositoryMock(db)
	transactionRepository := mocks.NewTransactionRepositoryMock(db)
	accountService := services.NewAccountService(accountRepository)
	transactionService := services.NewTransactionService(transactionRepository, accountService)

	resource := &resources.Transaction{
		AccountID:       2,
		OperationTypeID: 1,
		Amount:          1,
	}

	_, err := transactionService.Create(resource)

	assert.Error(err)
	assert.Equal(404, err.Code)
}

func TestShouldListPendentTransaction(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	accountRepository := mocks.NewAccountRepositoryMock(db)
	transactionRepository := mocks.NewTransactionRepositoryMock(db)
	accountService := services.NewAccountService(accountRepository)
	transactionService := services.NewTransactionService(transactionRepository, accountService)

	result, _ := transactionService.ListPendent(1)

	assert.Equal(int64(1), result.AccountID)
	assert.Equal(int64(1), result.TransactionID)
	assert.Equal(int64(1), result.OperationTypeID)
	assert.Equal(float32(-100), result.Amount)
	assert.Equal(float32(-100), result.Balance)
}

func TestShouldntListPendentTransactionByAccountNotFound(t *testing.T) {
	assert := assert.New(t)
	db := &gorm.DB{}
	accountRepository := mocks.NewAccountRepositoryMock(db)
	transactionRepository := mocks.NewTransactionRepositoryMock(db)
	accountService := services.NewAccountService(accountRepository)
	transactionService := services.NewTransactionService(transactionRepository, accountService)

	_, err := transactionService.ListPendent(2)

	assert.Error(err)
	assert.Equal(404, err.Code)
}
