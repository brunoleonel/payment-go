package controllers

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/services"
	"github.com/kataras/iris"
)

//TransactionController Handles requests related to transactions
type TransactionController struct {
	Service services.TransactionService
}

//Post Creates a transaction
func (c *TransactionController) Post(ctx iris.Context) {
	var resource resources.Transaction
	ctx.ReadJSON(&resource)
	response, err := c.Service.Create(&resource)
	if err != nil {
		ctx.StatusCode(err.Code)
		ctx.JSON(err)
		return
	}
	ctx.JSON(response)
}
