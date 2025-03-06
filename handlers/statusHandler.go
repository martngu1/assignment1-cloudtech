package handlers

import (
	"assignment1/services"
	"encoding/json"
	"log"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		statusData, err := services.GetStatus()
		if err != nil {
			log.Printf("errornotfound")
			http.Error(w, "errornotfound", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err1 := json.NewEncoder(w).Encode(statusData)
		if err1 != nil {
			log.Printf("bad")
			http.Error(w, "error", http.StatusBadRequest)
		}
	}

}
