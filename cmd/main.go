package main

import (
	"Travel-agency/internal/tour"

	"github.com/rs/zerolog/log"

	"net/http"
)

func main() {
	mux := http.NewServeMux()
	tourStorage := tour.NewInMemStorage()
	tourService := tour.NewService(tourStorage)
	tourHandler := tour.NewHandler(tourService)

	mux.HandleFunc("GET /tours", tourHandler.GetTours)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
