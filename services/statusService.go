package services

import (
	"assignment1/constants"
	"assignment1/models"
	"log"
	"net/http"
	"time"
)

var StartTime time.Time

func checkStatus(url string) int {

	resp, err := http.Head(url)
	if err != nil {
		log.Printf("Error in response to %s: %s", url, err)
		return http.StatusInternalServerError
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func GetStatus() (*models.Status, error) {
	countriesNowStatus := checkStatus(constants.CountriesNow + "countries")
	countriesRestStatus := checkStatus(constants.RESTCountries + "alpha/no")

	upTime := int(time.Since(StartTime).Seconds())

	return &models.Status{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: countriesRestStatus,
		Version:          constants.Version,
		Uptime:           upTime,
	}, nil
}
