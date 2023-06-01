package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"graphql-sample/model"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	DB = db
}
