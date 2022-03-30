package entities

import (
	"errors"
)

var (
	ErrNameRequired    = errors.New("Planet name is required")
	ErrPlanetNotExists = errors.New("Planet not exists")
	ErrPlanetExists    = errors.New("Planet name aready exist")
)

type Planet struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	Climate    string `gorm:"column:climate"`
	Land       string `gorm:"column:land"`
	Atmosphere string `gorm:"column:atmosphere"`
}

func (b *Planet) TableName() string {
	return "planet"
}
