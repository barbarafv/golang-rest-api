package repository

import (
	"aplicacao/source/domain/entities"
	"aplicacao/source/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func FindPlanets() (*[]entities.Planet, error) {

	var result []models.PlanetDBModel
	dbResult := DB.Find(&result)

	planets := dbPlanetToService(result)

	if err := dbResult.Error; err != nil {
		log.Panic("<FindPlanets> Error to find Planets ", err)
		return nil, err
	}
	return &planets, nil
}

func FindPlanetById(id int) (*entities.Planet, error) {

	var planetDBModel models.PlanetDBModel
	dbResult := DB.Where("id = ?", id).First(&planetDBModel)

	if err := dbResult.Error; err != nil {
		return nil, err
	}

	planet, err := entities.CreatePlanet(planetDBModel.Name, planetDBModel.Climate, planetDBModel.Land, planetDBModel.Atmosphere)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func UpdatePlanet(planet *entities.Planet, id int) (err error) {

	planetDBModel := dbPlanetFromService(planet)

	if err := DB.Model(&models.PlanetDBModel{Id: id}).Updates(planetDBModel).Error; err != nil {
		return err
	}

	return nil
}

func DeletePlanet(id int) (err error) {

	planetDBModel := models.PlanetDBModel{}

	if err := DB.Where("id = ?", id).Delete(planetDBModel).Error; err != nil {
		return err
	}
	return nil
}

func InsertPlanet(planet *entities.Planet) (err error) {

	planetDBModel := dbPlanetFromService(planet)

	if err := DB.Create(planetDBModel).Error; err != nil {
		return err
	}
	return nil
}

func ExistsPlanetByName(name string) bool {
	result := models.PlanetDBModel{}
	dbResult := DB.Where("name = ?", name).Find(&result)

	exists := dbResult.RowsAffected > 0

	return exists
}

func dbPlanetToService(dbPlanets []models.PlanetDBModel) []entities.Planet {
	var planets []entities.Planet

	for _, dbPlanet := range dbPlanets {
		planets = append(planets, entities.UnmarshalPlanet(dbPlanet.Id, dbPlanet.Name, dbPlanet.Climate, dbPlanet.Land, dbPlanet.Atmosphere))
	}

	return planets
}

func dbPlanetFromService(planet *entities.Planet) *models.PlanetDBModel {
	return &models.PlanetDBModel{
		Id:         planet.Id,
		Name:       planet.Name,
		Climate:    planet.Climate,
		Land:       planet.Land,
		Atmosphere: planet.Atmosphere,
	}
}
