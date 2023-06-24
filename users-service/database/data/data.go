package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	userModel "user/models"
)

func InsertData(db *gorm.DB) {
	// Insert data
	log.Info("Inserting data...")

	//Inserting users
	err := db.First(&userModel.User{}).Error

	if err != nil {
		db.Create(&userModel.User{Name: "Lautaro", LastName: "Saenz", UserName: "lautarose", Email: "lauti@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Julian", LastName: "Gergolet", UserName: "juligergolet", Email: "juli@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Hernan", LastName: "Lachampionliga", UserName: "hernanchampion", Email: "hernan@gmail.com", Pwd: "hola123"})
		db.Create(&userModel.User{Name: "Saul", LastName: "Hudson", UserName: "slash", Email: "slashGNR@gmail.com", Pwd: "hola123"})
	}

	log.Info("Data inserted")
}
