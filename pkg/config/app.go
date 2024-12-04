package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
