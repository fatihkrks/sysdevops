package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&category)
}

func (category Category) Get(where ...interface{}) Category {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}
	db.First(&category, where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var categories []Category
	db.Find(&categories, where...)
	return categories
}
