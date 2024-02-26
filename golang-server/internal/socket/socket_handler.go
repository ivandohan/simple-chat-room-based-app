package socket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// return origin == "http://localhost:3000"
		return true
	},
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request CreateRoomRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[request.ID] = &Room{
		ID: request.ID,
		Name: request.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, request)
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// path looks like /websocket/join-room/:roomId?userId=1&username=user
	roomId := c.Param("roomId")
	userId := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		Conn: conn,
		Message: make(chan *Message, 10),
		ID: userId,
		RoomID: roomId,
		Username: username,
	}

	message := &Message{
		Content: "A new user has joined the room.",
		RoomID: roomId,
		Username: username,
	}
}