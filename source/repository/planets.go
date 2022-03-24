package repository

import (
	entities "aplicacao/source/domain/entity"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func FindPlanets() (*[]entities.Planet, error) {

	var result []entities.Planet
	dbResult := DB.Find(&result)

	if err := dbResult.Error; err != nil {
		log.Panic("<FindPlanets> Error to find Planets ", err)
		return nil, err
	}
	return &result, nil
}

func FindPlanetById(id string) (*entities.Planet, error) {

	var result entities.Planet
	dbResult := DB.Where("id = ?", id).First(&result)

	if err := dbResult.Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func UpdatePlanet(planet *entities.Planet, id string) (err error) {

	idConv, _ := strconv.Atoi(id)
	DB.Model(&entities.Planet{Id: idConv}).Updates(planet)

	return nil
}

func DeletePlanet(planet *entities.Planet, id string) (err error) {

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
