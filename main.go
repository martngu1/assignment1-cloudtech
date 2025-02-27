package main

import (
	"assignment1/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	router := http.NewServeMux()

	router.HandleFunc("/countryinfo/v1/info/", handlers.InfoHandler)
	router.HandleFunc("/countryinfo/v1/population/", handlers.PopulationHandler)
	router.HandleFunc("/countryinfo/v1/status/", handlers.StatusHandler)

	log.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
