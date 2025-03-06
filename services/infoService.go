package services

import (
	"assignment1/constants"
	"assignment1/external/countriesNow"
	"assignment1/external/restCountries"
	"assignment1/models"
)

func GetCountryInfo(ISO2 string, limit int) (*models.CountryInfo, error) {
	urlRest := constants.RESTCountries + "alpha/" + ISO2
	urlNow := constants.CountriesNow + "countries/cities"

	countryData := restCountries.RequestInfo(urlRest)
	cityData := countriesNow.RequestCities(urlNow, countryData.Name)

	countryData.Cities = cityData

	if limit > 0 && len(countryData.Cities) > limit {
		countryData.Cities = countryData.Cities[:limit]
	}

	return &countryData, nil
}
