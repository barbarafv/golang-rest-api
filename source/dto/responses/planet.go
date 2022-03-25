package responses

type PlanetResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Land       string `json:"land"`
	Atmosphere string `json:"atmosphere"`
}

func CreatePlanetResponse(id string, name string, climate string, land string, atmosphere string) PlanetResponse {
	return PlanetResponse{
		Id:         id,
		Name:       name,
		Climate:    climate,
		Land:       land,
		Atmosphere: atmosphere,
	}
}
