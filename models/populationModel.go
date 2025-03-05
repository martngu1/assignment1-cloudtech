package models

type Population struct {
	Mean   int              `json:"mean"`
	Values []PopulationData `json:"values"`
}

type PopulationData struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}
type PopulationResponse struct {
	Data struct {
		PopulationCounts []PopulationData `json:"populationCounts"`
	} `json:"data"`
}
