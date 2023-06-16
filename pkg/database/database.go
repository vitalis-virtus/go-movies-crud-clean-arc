package database

// package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var (
	db *gorm.DB
)

// func Connect return connected db or error
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:rootpassword@tcp(localhost:3306)/movies?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")

	return db, nil

}
