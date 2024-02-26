package socket

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, isRoomExist := h.Rooms[client.RoomID]; isRoomExist {
				room := h.Rooms[client.RoomID]

				if _, isClientExist := room.Clients[client.ID]; !isClientExist {
					room.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, isRoomExist := h.Rooms[client.RoomID]; isRoomExist {
				if _, ok := h.Rooms[client.RoomID].Clients[client.ID]; ok {

					if len(h.Rooms[client.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "User left the chat room.",
							RoomID:   client.RoomID,
							Username: client.Username,
						}
					}

					delete(h.Rooms[client.RoomID].Clients, client.ID)
					close(client.Message)
				}
			}

		case message := <-h.Broadcast:
			if _, isRoomExist := h.Rooms[message.RoomID]; isRoomExist {

				for _, client := range h.Rooms[message.RoomID].Clients {
					client.Message <- message
				}
			}
		}
	}
}