package database

import (
	"fmt"
	"os"
	"time"

	"github.com/brunoleonel/payment/app/models"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Account{}).Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.OperationType{}).Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Transaction{}).Set("gorm:table_options", "ENGINE=InnoDB")

	db.Model(&models.Transaction{}).AddForeignKey("Account_ID", "accounts(Account_ID)", "NO ACTION", "CASCADE")
	db.Model(&models.Transaction{}).AddForeignKey("OperationType_ID", "operation_types(OperationType_ID)", "NO ACTION", "CASCADE")

	var records []models.OperationType
	db.Find(&records)

	if len(records) <= 0 {
		db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('COMPRA A VISTA', 2);")
		db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('COMPRA PARCELADA', 1);")
		db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('SAQUE', 0);")
		db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('PAGAMENTO', 0);")
	}
}

//Connect connect on database
func Connect() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DATABASE"),
	)

	time.Sleep(20 * time.Second)

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		defer db.Close()
		panic(err.Error())
	}

	go migrate(db)
	db.LogMode(true)
	return db
}
