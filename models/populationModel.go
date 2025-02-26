package models

type Population struct {
	Mean   int              `json:"mean"`
	Values []PopulationData `json:"values"`
}

type PopulationData struct {
	Year  int `json:"year"`
	Count int `json:"value"`
}
