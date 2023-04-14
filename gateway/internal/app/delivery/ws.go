package delivery

import (
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Ws struct {
	Room *websocket.Room
}

func (w *Ws) WsConnect(c echo.Context) error {
	clientID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorln("WsConnect #1 ", err.Error())
		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Errorln("WsConnect #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	client := &websocket.Client{Conn: conn, Send: make(chan *util.Response), WritePumpClose: make(chan bool), ID: clientID}

	w.Room.Clients[clientID] = client

	go client.ReadPump()
	go client.WritePump()

	client.Send <- &util.Response{Data: &util.Client{ID: clientID}}
	return nil
}
