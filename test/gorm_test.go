package test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	gorm.Model
	Title  string
	Code  string
	Price uint
}

func TestName(t *testing.T) {
	db := InitDb()
	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})
}

func InitDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	return db
}




