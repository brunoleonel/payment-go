package repositories

import (
	"fmt"
	"time"

	"github.com/brunoleonel/payment/app/http/resources"

	"github.com/brunoleonel/payment/app/models"
	"github.com/jinzhu/gorm"
)

//TransactionRepository handles the operations related to transactions on the database
type TransactionRepository interface {
	Create(model *models.Transaction) *models.Transaction
	FindPendent(accountID int64) (result models.Transaction, err *resources.Error)
	Update(model *models.Transaction) *models.Transaction
}

type transactionRepository struct {
	db *gorm.DB
}

//NewTransactionRepository returns a new TransactionRepository instance
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

//Create creates a new transaction
func (repository *transactionRepository) Create(model *models.Transaction) *models.Transaction {
	repository.db.Create(model)
	return model
}

//Update updates a transaction
func (repository *transactionRepository) Update(model *models.Transaction) *models.Transaction {
	repository.db.Save(model)
	return model
}

//FindPendent lists the pending transactions for an account
func (repository *transactionRepository) FindPendent(accountID int64) (result models.Transaction, err *resources.Error) {

	var joins = fmt.Sprintf(
		"%s %s %s",
		"INNER JOIN operation_types",
		"ON transactions.OperationType_ID = operation_types.OperationType_ID",
		"AND operation_types.OperationType_ID <> 4",
	)

	query := repository.db.Table("transactions")
	query = query.Joins(joins)
	query = query.Select(
		"transactions.*, operation_types.ChargeOrder",
	)
	query = query.Where("transactions.Account_ID = ?", accountID)
	query = query.Where("transactions.Balance < 0")
	query = query.Where("transactions.DueDate >= ?", time.Now())
	query = query.Order("operation_types.ChargeOrder ASC, transactions.EventDate DESC")
	notFound := query.First(&result).RecordNotFound()

	if notFound {
		err = &resources.Error{
			Code:    404,
			Message: "No pendent transaction found",
		}
	}

	return
}
