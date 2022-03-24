package repository

import (
	"aplicacao/source/configuration"
	entities "aplicacao/source/domain/entity"
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
	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/dbtest?charset=utf8&parseTime=True&loc=Local", os.Getenv(DB_USER))
	//connectionString := "root:MyApplication92@tcp(127.0.0.1:3306)/dbtest?charset=utf8&parseTime=True&loc=Local"
	return connectionString
}
