package mocks

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
	"github.com/brunoleonel/payment/app/repositories"
	"github.com/jinzhu/gorm"
)

type accountRepositoryMock struct {
	DB *gorm.DB
}

//NewAccountRepositoryMock returns a new AccountRepository instance
func NewAccountRepositoryMock(db *gorm.DB) repositories.AccountRepository {
	return &accountRepositoryMock{
		DB: db,
	}
}

//Create creates an new account
func (repository *accountRepositoryMock) Create(account *models.Account) *models.Account {
	account.AccountID = 1
	return account
}

//Update updates an account
func (repository *accountRepositoryMock) Update(account *models.Account) *models.Account {
	return account
}

//Find find an account
func (repository *accountRepositoryMock) Find(id int64) (account *models.Account, err *resources.Error) {
	account = &models.Account{}
	notFound := id != 1
	if notFound {
		err = &resources.Error{
			Code:    404,
			Message: "Account not found",
		}
		return
	}

	account.AccountID = 1
	account.AvailableCreditLimit = 100
	account.AvailableWithdrawalLimit = 100
	account.Transactions = nil
	return
}

//List the accounts
func (repository *accountRepositoryMock) List() []models.Account {
	accounts := make([]models.Account, 0)
	for i := 0; i < 3; i++ {
		account := models.Account{
			AccountID:                int64(i + 1),
			AvailableCreditLimit:     100,
			AvailableWithdrawalLimit: 100,
			Transactions:             nil,
		}
		accounts = append(accounts, account)
	}
	return accounts
}
