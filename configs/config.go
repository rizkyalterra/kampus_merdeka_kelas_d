package configs

import (
	"kelasd/models/news"
	"kelasd/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CnnectDB() {
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/kel_d?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database tidak connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(users.User{}, news.News{})
}
