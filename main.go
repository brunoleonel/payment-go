package main

import (
	"fmt"

	"github.com/brunoleonel/payment/app/http/controllers"
	"github.com/brunoleonel/payment/app/repositories"
	"github.com/brunoleonel/payment/app/services"
	"github.com/brunoleonel/payment/infra/database"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const version = "v1"

var db = database.Connect()

var accountRepository = repositories.NewAccountRepository(db)
var accountService = services.NewAccountService(accountRepository)

var transactionRepository = repositories.NewTransactionRepository(db)
var transactionService = services.NewTransactionService(transactionRepository, accountService)

var operationRepository = repositories.NewOperationRepository(db)
var operationService = services.NewOperationService(operationRepository, transactionService, accountService)

func route(path string) string {
	return fmt.Sprintf("%s/%s", version, path)
}

func main() {
	app := iris.New()

	mvc.Configure(app.Party(route("accounts")), accounts)
	mvc.Configure(app.Party(route("transactions")), transactions)
	mvc.Configure(app.Party(route("payments")), payments)

	app.Run(
		iris.Addr("0.0.0.0:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}

func accounts(app *mvc.Application) {
	app.Register(accountService)
	app.Handle(new(controllers.AccountController))
}

func transactions(app *mvc.Application) {
	app.Register(transactionService)
	app.Handle(new(controllers.TransactionController))
}

func payments(app *mvc.Application) {
	app.Register(operationService)
	app.Handle(new(controllers.PaymentController))
}
