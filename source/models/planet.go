package models

type PlanetDBModel struct {
	Id         int    `gorm:"column:id;primaryKey"`
	Name       string `gorm:"column:name"`
	Climate    string `gorm:"column:climate"`
	Land       string `gorm:"column:land"`
	Atmosphere string `gorm:"column:atmosphere"`
}

func (b *PlanetDBModel) TableName() string {
	return "planet"
}
