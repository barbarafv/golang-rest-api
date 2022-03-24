package repository

import (
	"aplicacao/source/configuration"
	entities "aplicacao/source/domain/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(configuration.Config.DBDriver, BuildDBConfig())

	if err != nil {
		log.Panic("An error ocurred during try to connect a database ", err)
	}

	DB.AutoMigrate(&entities.Planet{})

}

func BuildDBConfig() string {
	connectionString := configuration.Config.DBUser + ":" + configuration.Config.DBPass + "@" + configuration.Config.DBSource
	return connectionString
}
