package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn/ginEssential/models"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ginEssential.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}