package main

import (
	"log"
	"net/http"

	"github.com/Sahil2k07/go-SOLID/src/database"
	"github.com/Sahil2k07/go-SOLID/src/routes"
)

func main() {
	database.DbConnect()
	defer database.DbDisconnect()

	router := routes.AppRouter()

	port := ":2120"

	log.Printf("Server is running on port %s\n", port)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error while starting server %s\n", err)
	}
}
