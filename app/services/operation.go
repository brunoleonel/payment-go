package services

import (
	"github.com/brunoleonel/payment/app/repositories"
)

//OperationService handles the business logic related to payments
type OperationService interface {
}

type operationService struct {
	repository         repositories.OperationRepository
	transactionService TransactionService
}

//NewOperationService returns a new OperationService instance
func NewOperationService(repository repositories.OperationRepository) OperationService {
	return &operationService{
		repository: repository,
	}
}
