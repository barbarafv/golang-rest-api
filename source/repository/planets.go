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

	dbResult := DB.Where("id = ?", id).First(&planet)

	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &planet, nil
}

func UpdatePlanet(planet *entities.Planet, id int) (err error) {

	if err := DB.Model(&entities.Planet{Id: id}).Updates(planet).Error; err != nil {
		return err
	}

	return nil
}

func DeletePlanet(id int) (err error) {

	planet := entities.Planet{}

	if err := DB.Where("id = ?", id).Delete(planet).Error; err != nil {
		return err
	}
	return nil
}

func InsertPlanet(planet *entities.Planet) (err error) {

	if err := DB.Create(planet).Error; err != nil {
		return err
	}
	return nil
}

func ExistsPlanetByName(name string) bool {

	result := entities.Planet{}
	dbResult := DB.Where("name = ?", name).Find(&result)

	exists := dbResult.RowsAffected > 0

	return exists
}
