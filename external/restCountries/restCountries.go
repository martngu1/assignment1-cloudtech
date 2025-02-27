package restCountries

import (
	"assignment1/external"
	"assignment1/models"
	"encoding/json"
	"log"
)

// apiResponse is used to decode the REST Countries API response.
type apiResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		Png string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}

// RequestInfo sends a GET request to the REST Countries API, decodes the response, and maps it to models.CountryInfo.
func RequestInfo(url string) models.CountryInfo {
	var apiResp []apiResponse
	body := external.SendGetRequest(url)
	if err := json.Unmarshal(body, &apiResp); err != nil {
		log.Printf("Error decoding API response: %v", err)
		return models.CountryInfo{}
	}
	if len(apiResp) == 0 {
		log.Printf("No country data found")
		return models.CountryInfo{}
	}

	// Map the first element of the API response to models.CountryInfo
	resp := apiResp[0]
	var capital string
	if len(resp.Capital) > 0 {
		capital = resp.Capital[0]
	}

	return models.CountryInfo{
		Name:       resp.Name.Common,
		Continents: resp.Continents,
		Population: resp.Population,
		Languages:  resp.Languages,
		Borders:    resp.Borders,
		Flag:       resp.Flags.Png,
		Capital:    capital,
		Cities:     []string{}, // will be filled later
	}
}
