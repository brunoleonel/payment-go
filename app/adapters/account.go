package adapters

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
)

//AccountAdapter convert a model to a resource and vice-versa
type AccountAdapter struct {
}

//ToEntity converts resources.Account on model.Account
func (adapter *AccountAdapter) ToEntity(resource *resources.Account) *models.Account {
	var model models.Account
	model.AvailableCreditLimit = resource.AvailableCreditLimit.Amount
	model.AvailableWithdrawalLimit = resource.AvailableWithdrawalLimit.Amount
	return &model
}

//FromEntity converts model.Account on resources.Account
func (adapter *AccountAdapter) FromEntity(model *models.Account) *resources.Account {
	return &resources.Account{
		ID: model.AccountID,
		AvailableCreditLimit: resources.AvailableCreditLimit{
			Amount: model.AvailableCreditLimit,
		},
		AvailableWithdrawalLimit: resources.AvailableWithdrawalLimit{
			Amount: model.AvailableWithdrawalLimit,
		},
	}
}

//FromEntityCollection converts model slice on resource slice
func (adapter *AccountAdapter) FromEntityCollection(models []models.Account) *[]*resources.Account {
	accounts := make([]*resources.Account, 0)
	for _, item := range models {
		account := adapter.FromEntity(&item)
		accounts = append(accounts, account)
	}
	return &accounts
}
