package constants

const (
	CountriesNow  = "http://129.241.150.113:3500/api/v0.1/"
	RESTCountries = "http://129.241.150.113:8080/v3.1/"
)

const (
	Iso2 = "{two_letter_country_code}"
)

const (
	Version        = "v1"
	DefaultPath    = "/countryinfo/" + Version
	InfoPath       = DefaultPath + "/info/" + Iso2
	PopulationPath = DefaultPath + "/population/" + Iso2
	StatusPath     = DefaultPath + "/Status/"
)

const (
	ErrorDefault       = "The following error has occurred: "
	ErrorPath          = ErrorDefault + "Invalid path"
	ErrorPathParameter = ErrorDefault + "Invalid path parameter"
	ErrorRequest       = ErrorDefault + "Cannot create request"
	ErrorLimit         = ErrorDefault + "Invalid limit"
)
