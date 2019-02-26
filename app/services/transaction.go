package services

import (
	"github.com/brunoleonel/payment/app/repositories"
)

//TransactionService handles the business logic related to transactions
type TransactionService interface {
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
