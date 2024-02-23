package main

import (
	"golang-server/database"
	"golang-server/internal/user"
	"log"
)

func main() {
	dbConnection, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %s\n", err)
	}

	userRepository := user.NewRepository(dbConnection.GetDatabase())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
}