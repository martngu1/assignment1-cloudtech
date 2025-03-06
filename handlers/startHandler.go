package handlers

import (
	"assignment1/constants"
	"fmt"
	"net/http"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := fmt.Sprintf("This API does not provide a service at the root endpoint %s.\n"+
		"Please add one of the following paths: \n%s\n%s\n%s",
		r.URL.Path, constants.InfoPath, constants.PopulationPath, constants.StatusPath)

	_, err := fmt.Fprint(w, response)
	if err != nil {
		http.Error(w, "Error when returning output, please try again!", http.StatusInternalServerError)
	}
}
