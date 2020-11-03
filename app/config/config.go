package config

import (
	"implementasi-mvc/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@/simple_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB" + err.Error())
	}

	db.AutoMigrate(new(model.Account), new(model.TransactionModel))
	return db
}
