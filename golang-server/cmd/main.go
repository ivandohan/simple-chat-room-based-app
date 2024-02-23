package main

import (
	"golang-server/database"
	"log"
)

func main() {
	_, err := database.NewDatabase()

	if err != nil {
		log.Fatalf("Could not initialize database connection: %s\n", err)
	}
}