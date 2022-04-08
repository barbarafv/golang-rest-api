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

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func UpdatePlanet(request *requests.PlanetRequest, id int) {

	err := validadePlanet(id)

	if err != nil {
		log.Panic(err)
	}

	planet := CreatePlanet(request)

	err = repository.UpdatePlanet(&planet, id)

	if err != nil {
		log.Panic("<UpdatePlanet> An error ocurred during update", err)
	}

}

func FindPlanets() *[]responses.PlanetResponse {
	planetsResponse := []responses.PlanetResponse{}
	planets, err := repository.FindPlanets()

	for _, planet := range planets {
		planetsResponse = append(planetsResponse,
			CreatePlanetResponse(strconv.Itoa(planet.Id), &planet))
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

	planetResponse := CreatePlanetResponse(strconv.Itoa(id), result)

	if err != nil {
		log.Panic("<FindPlanetById> An error ocurred during select by id", err)
	}
	return &planetResponse
}

func InsertPlanet(request *requests.PlanetRequest) {

	planet := CreatePlanet(request)

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
		if err == gorm.ErrRecordNotFound {
			panic(&exception.HttpException{StatusCode: http.StatusNotFound,
				Message: fmt.Sprintf("Planet cannot deleted because id %d was not found", id)})
		}
		log.Panic("An error ocurred during delete", err)
	}
}

func validadePlanet(id int) error {

	planetToUpdate, err := repository.FindPlanetById(id)

	if err != nil {
		panic(&exception.HttpException{StatusCode: http.StatusNotFound,
			Message: fmt.Sprintf("Planet with id %d not found", id)})
	}

	if planetToUpdate == nil {
		return entities.ErrPlanetNotExists
	}

	return nil
}

func CreatePlanetResponse(id string, planet *entities.Planet) (response responses.PlanetResponse) {

	return responses.PlanetResponse{
		Id:         id,
		Name:       planet.Name,
		Climate:    planet.Climate,
		Land:       planet.Land,
		Atmosphere: planet.Atmosphere,
	}

}

func CreatePlanet(request *requests.PlanetRequest) entities.Planet {

	return entities.Planet{
		Name:       request.Name,
		Climate:    request.Climate,
		Land:       request.Land,
		Atmosphere: request.Atmosphere,
	}
}
