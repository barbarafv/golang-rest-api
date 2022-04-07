package responses

type PlanetResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Land       string `json:"land"`
	Atmosphere string `json:"atmosphere"`
}

