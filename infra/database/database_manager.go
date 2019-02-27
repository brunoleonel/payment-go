package database

import (
	"github.com/brunoleonel/payment/app/models"
	"github.com/jinzhu/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Account{}).Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.OperationType{}).Set("gorm:table_options", "ENGINE=InnoDB")
	db.AutoMigrate(&models.Transaction{}).Set("gorm:table_options", "ENGINE=InnoDB")

	db.Model(&models.Transaction{}).AddForeignKey("Account_ID", "accounts(Account_ID)", "NO ACTION", "CASCADE")
	db.Model(&models.Transaction{}).AddForeignKey("OperationType_ID", "operation_types(OperationType_ID)", "NO ACTION", "CASCADE")

	db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('COMPRA A VISTA', 2);")
	db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('COMPRA PARCELADA', 1);")
	db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('SAQUE', 0);")
	db.Exec("INSERT INTO operation_types(Description, ChargeOrder) VALUES('PAGAMENTO', 0);")
}

//Connect connect on database
func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/payment?charset=utf8mb4&parseTime=True")
	if err != nil {
		defer db.Close()
		panic(err.Error())
	}

	go migrate(db)

	return db
}
