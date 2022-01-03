package controllers

import (
	"fmt"
	"freedom-senryu-online/handlers"
	"freedom-senryu-online/middlewares"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	clients = make(Clients)
)

type WebSocket struct{}

type Clients map[*websocket.Conn]bool

func (cws WebSocket) ControllWebSocket(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	clients[conn] = true
	ticker := time.NewTicker(time.Second)

	hu = handlers.User{}

	rid := c.Params.ByName("room_id")

	u := middlewares.GetCurrentUserInfo()
	uid := u.UID

	userInfo, err := hu.JoinRoom(uid, rid)

	fmt.Println("[user_info]")
	fmt.Println(userInfo)

	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	clients.WriteJSON(userInfo)

	defer ticker.Stop()
	defer hu.ExitRoom(uid)

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		clients.WriteMessage(t, msg)
	}
}

func (clients Clients) WriteJSON(v interface{}) error {
	for client := range clients {
		if err := client.WriteJSON(v); err != nil {
			return err
		}
	}
	return nil
}

func (clients Clients) WriteMessage(messageType int, data []byte) error {
	for client := range clients {
		if err := client.WriteMessage(messageType, data); err != nil {
			return err
		}
	}
	return nil
}
