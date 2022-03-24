package entities

import (
	"errors"
)

var (
	ErrNameRequired = errors.New("Planet name is required")
)

type Planet struct {
	Id         int
	Name       string
	Climate    string
	Land       string
	Atmosphere string
}

func CreatePlanet(name string, climate string, land string, atmosphere string) (Planet, error) {
	if name == "" {
		return Planet{}, ErrNameRequired
	}

	return Planet{
		Name:       name,
		Climate:    climate,
		Land:       land,
		Atmosphere: atmosphere,
	}, nil
}

func (p *Planet) UpdatePlanet(newName *string, newClimate *string, newLand *string, newAtmosphere *string) error {

	if newName != nil {
		p.Name = *newName
	}
	if newClimate != nil {
		p.Climate = *newClimate
	}
	if newLand != nil {
		p.Climate = *newLand
	}
	if newAtmosphere != nil {
		p.Atmosphere = *newAtmosphere
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	return nil
}

func (b *Planet) TableName() string {
	return "planet"
}