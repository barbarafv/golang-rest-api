package service

import (
	"app/source/controllers/requests"
	"app/source/controllers/responses"
	"app/source/domain/entities"
	"app/source/domain/exception"
	"app/source/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func UpdatePlanet(request *requests.PlanetRequest, id int) {

	FindPlanetById(id)

	planet := mapToEntityPlanet(request)

	err := repository.UpdatePlanet(&planet, id)

	if err != nil {
		log.Panic("<UpdatePlanet> An error ocurred during update", err)
	}
}

func FindPlanets() *[]responses.PlanetResponse {
	planetsResponse := []responses.PlanetResponse{}
	planets, err := repository.FindPlanets()

	for _, planet := range planets {
		planetsResponse = append(planetsResponse,
			mapToResponsePlanet(strconv.Itoa(planet.Id), &planet))
	}

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select", err)
	}
	return &planetsResponse
}

func FindPlanetById(id int) *responses.PlanetResponse {

	result, err := repository.FindPlanetById(id)

	if err != nil {
		panic(exception.NewNotFoundException(fmt.Sprintf("Planet %d was not found", id)))
	}

	planetResponse := mapToResponsePlanet(strconv.Itoa(id), result)

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select by id", err)
	}
	return &planetResponse
}

func InsertPlanet(request *requests.PlanetRequest) {

	planet := mapToEntityPlanet(request)

	planetByName := repository.ExistsPlanetByName(planet.Name)

	if planetByName {
		panic(&exception.HttpException{StatusCode: http.StatusBadRequest,
			Message: fmt.Sprintf("Planet %s aready exist", planet.Name)})
	}
	err := repository.InsertPlanet(&planet)

	if err != nil {
		log.Panic("<InsertPlanet> An error ocurred during insert", err)
	}
}

func DeletePlanet(id int) {

	err := repository.DeletePlanet(id)

	if err != nil {
		if strings.Contains(fmt.Sprint(err), "record not found") {
			panic(&exception.HttpException{StatusCode: http.StatusNotFound,
				Message: fmt.Sprintf("Planet cannot deleted because id %d was not found", id)})
		}
		log.Panic("An error ocurred during delete", err)
	}
}

func mapToResponsePlanet(id string, planet *entities.Planet) (response responses.PlanetResponse) {

	return responses.PlanetResponse{
		Id:         id,
		Name:       planet.Name,
		Climate:    planet.Climate,
		Land:       planet.Land,
		Atmosphere: planet.Atmosphere,
	}
}

func mapToEntityPlanet(request *requests.PlanetRequest) entities.Planet {

	return entities.Planet{
		Name:       request.Name,
		Climate:    request.Climate,
		Land:       request.Land,
		Atmosphere: request.Atmosphere,
	}
}
