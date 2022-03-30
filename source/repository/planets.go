package repository

import (
	"aplicacao/source/domain/entities"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func FindPlanets() (*[]entities.Planet, error) {

	var planets []entities.Planet
	dbResult := DB.Find(&planets)

	if err := dbResult.Error; err != nil {
		log.Panic("<FindPlanets> Error to find Planets ", err)
		return nil, err
	}
	return &planets, nil
}

func FindPlanetById(id int) (*entities.Planet, error) {

	planet := entities.Planet{}

	dbResult := DB.Where("id = ?", id).First(&entities.Planet{})

	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &planet, nil
}

func UpdatePlanet(planet *entities.Planet, id int) error {
	return DB.Model(&entities.Planet{Id: id}).Updates(planet).Error
}

func DeletePlanet(id int) error {
	return DB.Where("id = ?", id).Delete(entities.Planet{}).Error
}

func InsertPlanet(planet *entities.Planet) error {
	return DB.Create(planet).Error
}

func ExistsPlanetByName(name string) bool {

	result := entities.Planet{}
	dbResult := DB.Where("name = ?", name).Find(&result)

	return dbResult.RowsAffected > 0
}
