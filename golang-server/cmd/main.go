package main

import (
	"golang-server/database"
	"golang-server/internal/socket"
	"golang-server/internal/user"
	"golang-server/router"
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

	socketHub := socket.NewHub()
	socketHandler := socket.NewHandler(socketHub)

	router.InitRouter(userHandler, socketHandler)
	router.Start("0.0.0.0:8080")
}