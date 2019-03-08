package services

import (
	"math"

	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
	"github.com/brunoleonel/payment/app/repositories"
)

//OperationService handles the business logic related to payments
type OperationService interface {
	CreatePayment(resourceList []resources.Payment) (payments *resources.PaymentResponse)
}

type operationService struct {
	repository         repositories.OperationRepository
	transactionService TransactionService
	accountService     AccountService
}

//NewOperationService returns a new OperationService instance
func NewOperationService(repository repositories.OperationRepository, transactionService TransactionService, accountService AccountService) OperationService {
	return &operationService{
		repository:         repository,
		transactionService: transactionService,
		accountService:     accountService,
	}
}

//CreatePayment handles the logic to create a payment
func (service *operationService) CreatePayment(resourceList []resources.Payment) (response *resources.PaymentResponse) {

	response = &resources.PaymentResponse{}
	failedList := make([]resources.FailedPayment, 0)
	paidList := make([]resources.Payment, 0)

	for _, resource := range resourceList {

		transaction, err := service.transactionService.ListPendent(resource.AccountID)
		if err != nil {
			failed := resources.FailedPayment{
				Payment: resource,
				Reason:  err.Message,
			}
			failedList = append(failedList, failed)
			continue
		}

		originalBalance := transaction.Balance
		balance := originalBalance + resource.Amount

		if balance >= 0 {
			limit := math.Abs(float64(originalBalance))
			transaction.Balance = 0
			service.transactionService.Update(&transaction)
			service.accountService.AdjustAccountLimit(resource.AccountID, float32(limit), transaction.OperationTypeID)

			if balance > 0 {
				var creditTransaction models.Transaction
				creditTransaction.OperationTypeID = 4
				creditTransaction.AccountID = resource.AccountID
				creditTransaction.Amount = balance
				service.transactionService.CreateCreditTransaction(&creditTransaction)
			}

			paidList = append(paidList, resource)
			continue
		}

		limit := resource.Amount
		transaction.Balance = balance
		service.transactionService.Update(&transaction)
		service.accountService.AdjustAccountLimit(resource.AccountID, float32(limit), transaction.OperationTypeID)
		paidList = append(paidList, resource)
	}

	response.Failed = failedList
	response.Paid = paidList

	return
}
