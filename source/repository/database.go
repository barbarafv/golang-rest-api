package repository

import (
	"app/source/configuration"
	"app/source/domain/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(configuration.Config.DBDriver, BuildDBConfig())

	if err != nil {
		log.Panic("An error ocurred during try to connect a database ", err)
	}

	db.AutoMigrate(&entities.Planet{})

}

func BuildDBConfig() string {
	connectionString := configuration.Config.DBUser + ":" + configuration.Config.DBPass + "@" + configuration.Config.DBSource
	return connectionString
}
