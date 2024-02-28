package router

import (
	"golang-server/internal/socket"
	"golang-server/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, socketHandler *socket.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/websocket/create-room", socketHandler.CreateRoom)
	r.GET("/websocket/join-room/:roomId", socketHandler.JoinRoom)
	r.GET("/websocket/get-rooms", socketHandler.GetRoom)
	r.GET("/websocket/get-clients/:roomId", socketHandler.GetClients)
}

func Start(address string) error {
	return r.Run(address)
}