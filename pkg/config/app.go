package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	database *gorm.DB
)

func Connect() {
	db, err := gorm.Open("mysql", "root:12345678@/golangbase?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	database = db
}

func GetDB() *gorm.DB {
	return database
}
