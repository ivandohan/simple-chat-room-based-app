package router

import (
	"golang-server/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
}

func Start(address string) error {
	return r.Run(address)
}