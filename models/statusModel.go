package models

type Status struct {
	CountriesNowAPI  int    `json:"countriesnowapi"`
	RestCountriesAPI int    `json:"countriesrestapi"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}
