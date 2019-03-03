package repositories

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
	"github.com/jinzhu/gorm"
)

//AccountRepository handles the operations related to accounts on the database
type AccountRepository interface {
	Create(account *models.Account) *models.Account
	Update(account *models.Account) *models.Account
	Find(id int64) (account *models.Account, err *resources.Error)
	List() []models.Account
}

type accountRepository struct {
	DB *gorm.DB
}

//NewAccountRepository returns a new AccountRepository instance
func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		DB: db,
	}
}

//Create creates an new account
func (repository *accountRepository) Create(account *models.Account) *models.Account {
	repository.DB.Create(account)
	return account
}

//Update updates an account
func (repository *accountRepository) Update(account *models.Account) *models.Account {
	repository.DB.Save(account)
	return account
}

//Find find an account
func (repository *accountRepository) Find(id int64) (account *models.Account, err *resources.Error) {
	account = &models.Account{}
	notFound := repository.DB.First(account, id).RecordNotFound()
	if notFound {
		err = &resources.Error{
			Code:    404,
			Message: "Account not found",
		}
	}
	return
}

//List the accounts
func (repository *accountRepository) List() []models.Account {
	accounts := make([]models.Account, 0)
	repository.DB.Find(&accounts)
	return accounts
}
