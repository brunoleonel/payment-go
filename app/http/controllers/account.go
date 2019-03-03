package controllers

import (
	"github.com/brunoleonel/payment/app/http/resources"
	"github.com/brunoleonel/payment/app/services"
	"github.com/kataras/iris"
)

//AccountController handles requests related to accounts
type AccountController struct {
	Service services.AccountService
}

//GetLimits Returns the limits
func (c *AccountController) GetLimits(ctx iris.Context) {
	accounts := c.Service.GetLimits()
	ctx.JSON(&accounts)
}

//PatchBy Updates the limits of the identified account
func (c *AccountController) PatchBy(id int64, ctx iris.Context) {
	var account resources.Account
	ctx.ReadJSON(&account)
	response, err := c.Service.Update(id, &account)

	if err != nil {
		ctx.StatusCode(err.Code)
		ctx.JSON(err.Message)
		return
	}
	ctx.JSON(&response)
}

//Post Creates a new Account with limits
func (c *AccountController) Post(ctx iris.Context) {
	var account resources.Account
	ctx.ReadJSON(&account)
	response := c.Service.Create(&account)
	ctx.JSON(response)
}
