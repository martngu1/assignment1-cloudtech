package countriesNow

import (
	"assignment1/external"
	"assignment1/models"
	"encoding/json"
	"log"
)

func RequestCities(url string, country string) []string {

	resp := models.CityResponse{}

	dataBody := models.CountryRequest{
		Name: country,
	}

	body := external.SendPostRequest(url, dataBody)

	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("error", err)
	}
	return resp.Data

}

func RequestPopulation(url string, country string) []models.PopulationData {
	var resp models.PopulationResponse

	dataBody := models.CountryRequest{
		Name: country,
	}

	body := external.SendPostRequest(url, dataBody)

	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("Error unmarshaling response: %v", err)
		return nil
	}

	return resp.Data.PopulationCounts
}
