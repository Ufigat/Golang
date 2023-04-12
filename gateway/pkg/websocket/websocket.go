package websocket

import (
	"gateway/pkg/response/car"
	"log"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Conn *websocket.Conn

	Send chan car.DataResponse

	WritePumpClose chan bool
}

func (c *Client) ReadPump() {
	defer func() {
		c.Conn.Close()
		c.WritePumpClose <- true
		close(c.WritePumpClose)
		close(c.Send)
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
