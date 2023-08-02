package config

import (
	"LK_blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/LK_blog?parseTime=true"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})

	return db, nil
}
