package requests

type UpdatePlanetRequest struct {
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Land       string `json:"land"`
	Atmosphere string `json:"atmosphere"`
}
