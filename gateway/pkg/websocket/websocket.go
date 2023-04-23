package websocket

import (
	"gateway/pkg/util"
	"log"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Room *Room

	Conn *websocket.Conn

	Send chan *util.Response

	WritePumpClose chan bool

	ID int
}

type Room struct {
	Unregister chan *Client
	Register   chan *Client
	Clients    map[int]*Client
}

func (r *Room) Work() {

	for {
		select {
		case client := <-r.Register:
			r.Clients[client.ID] = client
		case client := <-r.Unregister:
			delete(r.Clients, client.ID)
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.WritePumpClose <- true
	}()
	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("readPump close: %v", err)

				c.Room.Unregister <- c
			}
			break
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
		close(c.WritePumpClose)
		close(c.Send)
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				break
			}

			if err := c.Conn.WriteJSON(message); err != nil {
				c.Conn.WriteJSON("error")
			}

		case <-c.WritePumpClose:
			return
		}
	}
}
