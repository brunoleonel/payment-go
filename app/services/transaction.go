package services

import (
	"time"

	"github.com/brunoleonel/payment/app/models"

	"github.com/brunoleonel/payment/app/adapters"
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/repositories"
)

//TransactionService handles the business logic related to transactions
type TransactionService interface {
	Create(resource *resources.Transaction) *resources.Transaction
}

type transactionService struct {
	repository repositories.TransactionRepository
}

//NewTransactionService returns a new TransactionService instance
func NewTransactionService(repository repositories.TransactionRepository) TransactionService {
	return &transactionService{
		repository: repository,
	}
}

//Create handles the business logic of transaction creation
func (service *transactionService) Create(resource *resources.Transaction) *resources.Transaction {
	var adapter adapters.TransactionAdapter

	model := adapter.ToEntity(resource)
	processed := processValue(model, model.Amount)
	transaction := service.repository.Create(processed)
	response := adapter.FromEntity(transaction)
	return response
}

func processValue(model *models.Transaction, balance float32) *models.Transaction {
	date := time.Now()
	due := date.AddDate(0, 0, 1)
	model.Balance = balance
	model.EventDate = &date
	model.DueDate = &due
	return model
}
