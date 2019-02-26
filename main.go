package main

import (
	"fmt"

	"github.com/brunoleonel/payment/app/http/controllers"
	"github.com/brunoleonel/payment/app/repositories"
	"github.com/brunoleonel/payment/app/services"
	"github.com/brunoleonel/payment/infra/database"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const version = "v1"

var db = database.Connect()

func route(path string) string {
	return fmt.Sprintf("%s/%s", version, path)
}

func main() {
	app := iris.New()

	mvc.Configure(app.Party(route("accounts")), accounts)
	mvc.Configure(app.Party(route("transactions")), transactions)
	mvc.Configure(app.Party(route("payments")), payments)

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func accounts(app *mvc.Application) {
	accountRepository := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepository)
	app.Register(accountService)
	app.Handle(new(controllers.AccountController))
}

func transactions(app *mvc.Application) {
	transactionRepository := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepository)
	app.Register(transactionService)
	app.Handle(new(controllers.TransactionController))
}

func payments(app *mvc.Application) {
	operationRepository := repositories.NewOperationRepository(db)
	operationService := services.NewOperationService(operationRepository)
	app.Register(operationService)
	app.Handle(new(controllers.PaymentController))
}
