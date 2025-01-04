package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connection() {
	dsn := "root:root@tcp(localhost:3306)/book_store?charset=utf8&parseTime=True&loc=Local"
	data, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = data
}

func GetDB() *gorm.DB {
	return db
}
