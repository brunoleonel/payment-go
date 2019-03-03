package services

import (
	"github.com/brunoleonel/payment/app/models"

	"github.com/brunoleonel/payment/app/adapters"
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/repositories"
)

//AccountService handles the business logic related to accounts
type AccountService interface {
	Create(resource *resources.Account) *resources.Account
	Update(id int64, resource *resources.Account) (response *resources.Account, err *resources.Error)
	GetLimits() *[]*resources.Account
	Find(id int64) (model *models.Account, err *resources.Error)
	AdjustAccountLimit(id int64, value float32, operation int64) (err *resources.Error)
}

type accountService struct {
	Repository repositories.AccountRepository
}

//NewAccountService returns a new AccountService instance
func NewAccountService(repository repositories.AccountRepository) AccountService {
	return &accountService{
		Repository: repository,
	}
}

//Create handles business logic of account creation
func (service *accountService) Create(resource *resources.Account) *resources.Account {
	var adapter adapters.AccountAdapter
	model := adapter.ToEntity(resource)
	service.Repository.Create(model)
	response := adapter.FromEntity(model)
	return response
}

//Update handles business logic of account update
func (service *accountService) Update(id int64, resource *resources.Account) (response *resources.Account, err *resources.Error) {
	var adapter adapters.AccountAdapter
	model, err := service.Repository.Find(id)

	if err != nil {
		return
	}

	model.AvailableCreditLimit += resource.AvailableCreditLimit.Amount
	model.AvailableWithdrawalLimit += resource.AvailableWithdrawalLimit.Amount
	model = service.Repository.Update(model)
	response = adapter.FromEntity(model)
	return
}

//GetLimits returns the limits of the accounts
func (service *accountService) GetLimits() *[]*resources.Account {
	var adapter adapters.AccountAdapter
	accounts := service.Repository.List()
	list := adapter.FromEntityCollection(accounts)
	return list
}

//Find finds an account
func (service *accountService) Find(id int64) (model *models.Account, err *resources.Error) {
	model, err = service.Repository.Find(id)
	return
}

//AdjustAccountLimit adjusts acount limits after payment
func (service *accountService) AdjustAccountLimit(id int64, value float32, operation int64) (err *resources.Error) {
	account, err := service.Find(id)

	if err != nil {
		return
	}

	if operation != 3 {
		account.AvailableCreditLimit += value
		service.Repository.Update(account)
		return
	}

	account.AvailableWithdrawalLimit += value
	service.Repository.Update(account)
	return
}
