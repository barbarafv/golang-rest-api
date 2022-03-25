package service

import (
	"aplicacao/source/domain/entities"
	"aplicacao/source/dto/requests"
	"aplicacao/source/dto/responses"
	"aplicacao/source/repository"
	"aplicacao/source/utils"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UpdatePlanet(request *requests.PlanetRequest, id int) {

	var planet entities.Planet

	validadePlanet(id)

	planet.UpdatePlanet(&request.Name, &request.Climate, &request.Land, &request.Atmosphere)

	err := repository.UpdatePlanet(&planet, id)

	if err != nil {
		log.Panic("<UpdatePlanet> An error ocurred during update", err)
	}

}

func FindPlanets() *[]responses.PlanetResponse {
	planetsResponse := []responses.PlanetResponse{}
	planets, err := repository.FindPlanets()

	for _, planet := range *planets {
		idConv := utils.ConvertToString(planet.Id)
		planetsResponse = append(planetsResponse, responses.CreatePlanetResponse(idConv, planet.Name, planet.Climate, planet.Land, planet.Atmosphere))
	}

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select", err)
	}
	return &planetsResponse
}

func FindPlanetById(id int) *responses.PlanetResponse {

	result, err := repository.FindPlanetById(id)

	utils.ConvertToString(id)

	planetResponse := responses.CreatePlanetResponse(utils.ConvertToString(id), result.Name, result.Climate, result.Land, result.Atmosphere)

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select by id", err)
	}
	return &planetResponse
}

func InsertPlanet(request *requests.PlanetRequest) {

	planet, err := entities.CreatePlanet(request.Name, request.Climate, request.Land, request.Atmosphere)

	planetByName := repository.ExistsPlanetByName(planet.Name)

	if planetByName {
		log.Panic("<InsertPlanet> Planet name aready exist!")
	}
	err = repository.InsertPlanet(&planet)

	if err != nil {
		log.Panic("<InsertPlanet> An error ocurred during insert", err)
	}
}

func DeletePlanet(id int) {

	err := repository.DeletePlanet(id)

	if err != nil {
		log.Panic("<DeletePlanet> An error ocurred during delete", err)
	}
}

func validadePlanet(id int) error {

	planetToUpdate, err := repository.FindPlanetById(id)

	if err != nil {
		log.Panic("<FindPlanetById> Error to find planet by id", err)
	}

	if planetToUpdate == nil {
		return entities.ErrPlanetNotExists
	}

	return nil
}
