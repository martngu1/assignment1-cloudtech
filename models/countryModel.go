package models

type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"Cities"`
}

type CityResponse struct {
	Data []string `json:"data"`
}
type CityRequest struct {
	Name string `json:"country"`
}
