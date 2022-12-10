package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
}
