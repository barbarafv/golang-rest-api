package requests

type PlanetRequest struct {
	Name       string `json:"name" binding:"required"`
	Climate    string `json:"climate"`
	Land       string `json:"land"`
	Atmosphere string `json:"atmosphere"`
}
