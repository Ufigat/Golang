package websocket

import (
	"fmt"
	"gateway/pkg/util"
	"log"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Conn *websocket.Conn

	Send chan *util.Response

	WritePumpClose chan bool

	ID int
}

type Room struct {
	Clients map[int]*Client
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
		fmt.Println("connection close")
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
			fmt.Println("defer after ReadPump")
			return
		}
	}
}
