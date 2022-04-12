package repository

import (
	"app/source/domain/entities"

	"gorm.io/gorm"
)

func FindPlanets() ([]entities.Planet, error) {
	planets := []entities.Planet{}
	return planets, db.Find(&planets).Error
}

func FindPlanetById(id int) (*entities.Planet, error) {
	planets := entities.Planet{}
	return &planets, db.Where("id = ?", id).First(&planets).Error
}

func UpdatePlanet(planet *entities.Planet, id int) error {
	return db.Model(&entities.Planet{Id: id}).Updates(planet).Error
}

func DeletePlanet(id int) error {

	dbResult := db.Where("id = ?", id).Delete(entities.Planet{})

	if dbResult.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return dbResult.Error
}

func InsertPlanet(planet *entities.Planet) error {
	return db.Create(planet).Error
}

func ExistsPlanetByName(name string) bool {
	dbResult := db.Where("name = ?", name).Find(&entities.Planet{})
	return dbResult.RowsAffected > 0
}
