package services

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
	"github.com/brunoleonel/payment/app/repositories"
)

//OperationService handles the business logic related to payments
type OperationService interface {
	CreatePayment(resource *resources.Payment) (transaction models.Transaction, err *resources.Error)
}

type operationService struct {
	repository         repositories.OperationRepository
	transactionService TransactionService
}

//NewOperationService returns a new OperationService instance
func NewOperationService(repository repositories.OperationRepository, transactionService TransactionService) OperationService {
	return &operationService{
		repository:         repository,
		transactionService: transactionService,
	}
}

//CreatePayment handles the logic to create a payment
func (service *operationService) CreatePayment(resource *resources.Payment) (transaction models.Transaction, err *resources.Error) {
	transaction, err = service.transactionService.ListPendent(resource.AccountID)
	if err != nil {
		return
	}

	originalBalance := transaction.Balance
	balance := originalBalance + resource.Amount
	//limit := originalBalance - resource.Amount

	if balance > 0 {
		transaction.Balance = 0
		service.transactionService.Update(&transaction)

		var creditTransaction models.Transaction
		creditTransaction.OperationTypeID = 4
		creditTransaction.AccountID = resource.AccountID
		creditTransaction.Amount = balance

		service.transactionService.CreateCreditTransaction(&creditTransaction)
	}

	transaction.Balance = balance
	service.transactionService.Update(&transaction)

	return
}
