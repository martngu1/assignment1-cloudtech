package handlers

import (
	"assignment1/constants"
	"assignment1/services"
	"encoding/json"
	"log"
	"net/http"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		// Extracts the two-letter country code from the URL path
		ISO2 := r.PathValue("two_letter_country_code")
		if len(ISO2) != 2 {
			log.Printf(constants.ErrorPathParameter)
			http.Error(w, constants.ErrorPathParameter, http.StatusBadRequest)
			return
		}

		limitStr := r.URL.Query().Get("limit")

		populationData, err := services.GetPopulation(ISO2, limitStr)
		if err != nil {
			log.Printf("errornotfound")
			http.Error(w, "errornotfound", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err2 := json.NewEncoder(w).Encode(populationData)
		if err2 != nil {
			log.Printf("bad")
			http.Error(w, "error", http.StatusBadRequest)
		}
	}

}
