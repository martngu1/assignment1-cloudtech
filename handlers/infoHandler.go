package handlers

import (
	"assignment1/constants"
	"assignment1/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		/*  //Splits the URL.Path by separating "/"
		splitPath := strings.Split(r.URL.Path, "/")
		if len(splitPath) < 4 {
			log.Printf(constants.ErrorPath)
			http.Error(w, constants.ErrorPath, http.StatusBadRequest)
			return
		}

		//Declares countryCode (ISO2) by taking the last segment of the split path
		ISO2 := splitPath[len(splitPath)-1]
		if len(ISO2) != 2 {
			log.Printf(constants.ErrorPathParameter)
			http.Error(w, constants.ErrorPathParameter, http.StatusBadRequest)
			return
		}
		*/
		// Extracts the two-letter country code from the URL path
		ISO2 := r.PathValue("two_letter_country_code")
		if len(ISO2) != 2 {
			log.Printf(constants.ErrorPathParameter)
			http.Error(w, constants.ErrorPathParameter, http.StatusBadRequest)
			return
		}

		limitStr := r.URL.Query().Get("limit")
		limit := 10 // Default limit (0 for all cities)

		if limitStr != "" {
			var err error
			limit, err = strconv.Atoi(limitStr)
			if err != nil || limit < 0 {
				log.Printf("Invalid limit parameter: %s", limitStr)
				http.Error(w, "Invalid limit parameter, must be a positive integer", http.StatusBadRequest)
				return
			}
		}

		countryInfo, err := services.GetCountryInfo(ISO2, limit)
		if err != nil {
			log.Printf("errornotfound")
			http.Error(w, "errornotfound", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err2 := json.NewEncoder(w).Encode(countryInfo)
		if err2 != nil {
			log.Printf("bad")
			http.Error(w, "error", http.StatusBadRequest)
		}

	}
}
