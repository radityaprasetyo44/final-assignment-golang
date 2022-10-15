package configs

import (
	"final-assignment/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Env.MySQL.User,
		Env.MySQL.Password,
		Env.MySQL.Host,
		Env.MySQL.Port,
		Env.MySQL.DBName))
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Photo{})
	db.AutoMigrate(models.Comment{})
	db.AutoMigrate(models.SocialMedia{})

	return db
}
