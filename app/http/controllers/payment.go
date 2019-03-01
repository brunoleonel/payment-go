package controllers

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/services"
	"github.com/kataras/iris"
)

//PaymentController handles requests related to payments
type PaymentController struct {
	Service services.OperationService
}

//Post creates one or more payments
func (c *PaymentController) Post(ctx iris.Context) {
	var resource resources.Payment
	ctx.ReadJSON(&resource)
	response, err := c.Service.CreatePayment(&resource)
	if err != nil {
		ctx.StatusCode(err.Code)
		ctx.JSON(err)
		return
	}
	ctx.JSON(response)
}
