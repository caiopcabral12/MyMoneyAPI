package database

import (
	md "MyMoneyAPI/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	stringConnection := "host=localhost user=caiopcabral password=0102 dbname=MyMoneyAPI port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Error while connecting to Database!")
	}
	DB.AutoMigrate(
		&md.Client{},
		&md.CreditCard{},
		&md.Payable{},
		&md.Receivable{},
		&md.Savings{},
	)
}
