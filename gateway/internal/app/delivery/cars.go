package delivery

import (
	"gateway/internal/app/service"
	"gateway/internal/app/usecase"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"
	"net/http"

	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

type Delivery struct {
	Conn        *rabbitmq.Connect
	SendCar     <-chan amqp.Delivery
	SendEngines <-chan amqp.Delivery
}

func (d *Delivery) GetCarEnginesByBrand(c echo.Context) error {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	client := &websocket.Client{Conn: conn, Send: make(chan *util.Response), WritePumpClose: make(chan bool)}

	go client.WritePump()

	usecase := &usecase.Usacase{Conn: d.Conn, SendCar: d.SendCar, SendEngines: d.SendEngines}

	go usecase.GetCarEnginesByBrand(client, c.Param("brand"))

	return nil
}

func (d *Delivery) GetCarEngine(c echo.Context) error {
	carResp, err := service.GetCar(c.Param("id"))
	if err != nil {
		log.Errorln("GetCarEngine #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	engineResp, err := service.GetEngine(carResp.Data.EngineID)
	if err != nil {
		log.Errorln("GetCarEngine #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp := &engine.ForCar{
		ID:     carResp.Data.ID,
		Engine: engineResp.Data,
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}
