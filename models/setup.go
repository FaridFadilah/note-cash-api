package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=my_cash port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err.Error())
	}

	db.AutoMigrate(&ReportTransaction{}, &AuthToken{}, &User{})

	return db
}