package main

import (
	"assignment1/constants"
	"assignment1/handlers"
	"assignment1/services"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	router := http.NewServeMux()
	services.StartTime = time.Now()

	router.HandleFunc(constants.InfoPath, handlers.InfoHandler)
	router.HandleFunc(constants.PopulationPath, handlers.PopulationHandler)
	router.HandleFunc(constants.StatusPath, handlers.StatusHandler)

	log.Println("Running on port", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
