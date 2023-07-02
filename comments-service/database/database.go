package database

import (
	commentClient "comments/clients/comments"
	commentModel "comments/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "project" //variable de entorno para nombre de la base de datos
	DBUser := "user"     //variable de entorno para el usuario de la base de datos
	//DBPass := ""
	DBPass := "password" //variable de entorno para la pass de la base de datos
	DBHost := "mysql_database"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	commentClient.Db = db

}

func StartDbEngine() {

	// We need to migrate all classes model.
	db.AutoMigrate(&commentModel.Comment{})

	log.Info("Finishing Migration Database Tables")
}
