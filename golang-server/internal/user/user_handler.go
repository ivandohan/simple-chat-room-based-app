package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var userRequest CreateUserRequest
	
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.Service.CreateUser(c.Request.Context(), &userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	var userRequest LoginUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser, err := h.Service.LoginService(c.Request.Context(), &userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("login-jwt", currentUser.accessToken, 3600, "/", "localhost", false, true)

	response := &LoginUserResponse{
		Username: currentUser.Username,
		ID: currentUser.ID,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("login-jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}