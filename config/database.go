package config

import (
	"fmt"
	"golang-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Impossible de se connecter à la base de données")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.RoomModel{})
	db.AutoMigrate(&models.Room{})

	fmt.Println("Base de données connectée !")
	DB = db
}
