package constants

const (
	CountriesNowAPI  = "http://129.241.150.113:3500/api/v0.1/"
	RESTCountriesAPI = "http://129.241.150.113:8080/v3.1/"
)

const (
	Iso2 = "{two_letter_country_code}"
)

const (
	Version        = "v1"
	DefaultPath    = "/countryinfo/" + Version
	InfoPath       = DefaultPath + "/info/"
	PopulationPath = DefaultPath + "/population/"
	StatusPath     = DefaultPath + "/Status/"
)
