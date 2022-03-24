package service

import (
	entities "aplicacao/source/domain/entity"
	"aplicacao/source/repository"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UpdatePlanet(planet *entities.Planet, id string) (err error) {

	planetToUpdate, err := repository.FindPlanetById(id)

	if planetToUpdate == nil {
		log.Panic("Planet not exists")
	}

	err = repository.UpdatePlanet(planet, id)

	if err != nil {
		log.Panic("<UpdatePlanet> An error ocurred during update", err)
	}
	return nil
}

func FindPlanets() *[]entities.Planet {
	result, err := repository.FindPlanets()

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select", err)
	}
	return result
}

func FindPlanetById(id string) *entities.Planet {
	result, err := repository.FindPlanetById(id)

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select by id", err)
	}
	return result
}

func InsertPlanet(planet *entities.Planet) (err error) {

	planetByName := repository.ExistsPlanetByName(planet.Name)

	if planetByName {
		log.Panic("<InsertPlanet> Planet name aready exist!")
	}
	err = repository.InsertPlanet(planet)

	if err != nil {
		log.Panic("<InsertPlanet> An error ocurred during insert", err)
	}
	return nil
}

func DeletePlanet(id string) (err error) {

	var planet entities.Planet

	err = repository.DeletePlanet(&planet, id)

	if err != nil {
		log.Panic("<DeletePlanet> An error ocurred during delete", err)
	}
	return nil
}
