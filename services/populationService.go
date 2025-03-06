package services

import (
	"assignment1/constants"
	"assignment1/external/countriesNow"
	"assignment1/external/restCountries"
	"assignment1/models"
	"log"
	"math"
	"strconv"
	"strings"
)

func GetPopulation(ISO2 string, limit string) (*models.Population, error) {
	urlRest := constants.RESTCountries + "alpha/" + ISO2
	urlNow := constants.CountriesNow + "countries/population"

	countriesRestData := restCountries.RequestInfo(urlRest)

	countriesNowPopulation := countriesNow.RequestPopulation(urlNow, countriesRestData.Name)

	if limit != "" {
		countriesNowPopulation = getLimitPopulation(limit, countriesNowPopulation)
	}

	return &models.Population{
		Mean:   getMean(countriesNowPopulation),
		Values: countriesNowPopulation}, nil

}

func getLimitPopulation(limitStr string, populationData []models.PopulationData) []models.PopulationData {
	timeRange := strings.Split(limitStr, "-")
	if len(timeRange) != 2 {
		return nil
	}

	startYear, err1 := strconv.Atoi(timeRange[0])
	endYear, err2 := strconv.Atoi(timeRange[1])
	if err1 != nil || err2 != nil {
		log.Printf("Invalid limit parameter: Start Year: %d and End Year: %d", startYear, endYear)
		return nil
	}

	var newPopulation []models.PopulationData
	for _, data := range populationData {
		if data.Year >= startYear && data.Year <= endYear {
			newPopulation = append(newPopulation, data)
		}
	}
	return newPopulation
}

func getMean(populationData []models.PopulationData) int {
	sum := 0
	for _, data := range populationData {
		sum += data.Value
	}
	if len(populationData) == 0 {
		return 0
	}

	mean := int(math.Round(float64(sum) / float64(len(populationData))))

	return mean
}
