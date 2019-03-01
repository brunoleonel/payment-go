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
	Create(resource *resources.Transaction) (response *resources.Transaction, err *resources.Error)
	ListPendent(accountID int64) (transaction models.Transaction, err *resources.Error)
	AdjustLimits(account *models.Account, transaction *models.Transaction)
	CreateCreditTransaction(model *models.Transaction)
	Update(model *models.Transaction)
}

type transactionService struct {
	repository     repositories.TransactionRepository
	accountService AccountService
}

//NewTransactionService returns a new TransactionService instance
func NewTransactionService(repository repositories.TransactionRepository, accountService AccountService) TransactionService {
	return &transactionService{
		repository:     repository,
		accountService: accountService,
	}
}

//Create handles the business logic of transaction creation
func (service *transactionService) Create(resource *resources.Transaction) (response *resources.Transaction, err *resources.Error) {
	var adapter adapters.TransactionAdapter

	account := service.accountService.Find(resource.AccountID)

	err = service.checkLimits(account, resource)
	if err != nil {
		return
	}

	model := adapter.ToEntity(resource)
	model.Amount = -model.Amount
	processed := service.processValue(model, model.Amount)
	transaction := service.repository.Create(processed)

	service.AdjustLimits(account, processed)

	processed.Amount = -processed.Amount
	response = adapter.FromEntity(transaction)
	return
}

func (service *transactionService) processValue(model *models.Transaction, balance float32) *models.Transaction {
	date := time.Now()
	due := date.AddDate(0, 0, 1)
	model.Balance = balance
	model.EventDate = &date
	model.DueDate = &due
	return model
}

func (service *transactionService) checkLimits(account *models.Account, resource *resources.Transaction) (err *resources.Error) {
	err = service.checkCreditLimit(account, resource)
	if err != nil {
		return
	}

	err = service.checkWithdrawalLimit(account, resource)
	if err != nil {
		return
	}

	return
}

func (service *transactionService) checkCreditLimit(account *models.Account, resource *resources.Transaction) (err *resources.Error) {
	creditOp := resource.OperationTypeID != 3
	hasLimit := account.AvailableCreditLimit >= resource.Amount
	if creditOp && !hasLimit {
		err = &resources.Error{
			Code:    400,
			Message: "Credit limit exceeded",
		}
	}
	return
}

func (service *transactionService) checkWithdrawalLimit(account *models.Account, resource *resources.Transaction) (err *resources.Error) {
	withdOp := resource.OperationTypeID == 3
	hasLimit := account.AvailableWithdrawalLimit >= resource.Amount
	if withdOp && !hasLimit {
		err = &resources.Error{
			Code:    400,
			Message: "Withdrawal limit exceeded",
		}
	}
	return
}

//AdjustLimits handles the logic of limit adjustment
func (service *transactionService) AdjustLimits(account *models.Account, transaction *models.Transaction) {
	var adapter adapters.AccountAdapter

	if transaction.OperationTypeID != 3 {
		account.AvailableCreditLimit += transaction.Amount
		resource := adapter.FromEntity(account)
		service.accountService.Update(account.AccountID, resource)
		return
	}

	account.AvailableWithdrawalLimit += transaction.Amount
	resource := adapter.FromEntity(account)
	service.accountService.Update(account.AccountID, resource)
}

//ListPendent handle the logic to list the pendent transactions
func (service *transactionService) ListPendent(accountID int64) (result models.Transaction, err *resources.Error) {
	result, err = service.repository.FindPendent(accountID)
	return
}

//CreateCreditTransaction handles the business logic of transaction creation for funds
func (service *transactionService) CreateCreditTransaction(model *models.Transaction) {
	processed := service.processValue(model, model.Amount)
	service.repository.Create(processed)
	return
}

//Update updates a transaction when there's a model in hand
func (service *transactionService) Update(model *models.Transaction) {
	service.repository.Update(model)
}
