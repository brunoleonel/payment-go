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
	var resources []resources.Payment
	ctx.ReadJSON(&resources)
	response := c.Service.CreatePayment(resources)
	ctx.JSON(response)
}
