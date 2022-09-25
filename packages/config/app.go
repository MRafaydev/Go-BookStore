package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	db = d
}

func GETDB() *gorm.DB {
	return db
}