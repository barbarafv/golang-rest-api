package entity

type Planet struct {
	Id         int    `json:"id;"`
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Land       string `json:"land"`
	Atmosphere string `json:"atmosphere"`
}

func (b *Planet) TableName() string {
	return "planet"
}
