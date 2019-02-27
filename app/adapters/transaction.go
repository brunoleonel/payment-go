package adapters

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/models"
)

//TransactionAdapter converts a model to a resource and vice-versa
type TransactionAdapter struct {
}

//ToEntity converts resource.Transaction on model.Transaction
func (adapter *TransactionAdapter) ToEntity(resource *resources.Transaction) *models.Transaction {
	var model models.Transaction
	model.AccountID = resource.AccountID
	model.OperationTypeID = resource.OperationTypeID
	model.Amount = resource.Amount
	return &model
}

//FromEntity converts model.Transaction on resource.Transaction
func (adapter *TransactionAdapter) FromEntity(model *models.Transaction) *resources.Transaction {
	return &resources.Transaction{
		AccountID:       model.AccountID,
		OperationTypeID: model.OperationTypeID,
		Amount:          model.Amount,
	}
}
