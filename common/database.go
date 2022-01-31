package common

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"learn/ginEssential/models"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	dsn := viper.GetString("datasource.dsn")

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&models.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}