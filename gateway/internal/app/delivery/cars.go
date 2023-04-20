package delivery

import (
	"gateway/internal/app/usecase"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Car struct {
	Conn *rabbitmq.Connect
	Room *websocket.Room
}

func (car *Car) GetCarEnginesByBrand(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("client"))
	if err != nil {
		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	usecase := usecase.Usacase{Room: car.Room, Conn: car.Conn}
	go usecase.GetCarEnginesByBrand(userID, c.Param("brand"))

	return c.JSON(http.StatusOK, &util.Response{Data: "request in processing"})
}

func (car *Car) GetCarEngine(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("client"))
	if err != nil {
		log.Errorln("GetCarEngine #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	usecase := usecase.Usacase{Room: car.Room, Conn: car.Conn}
	go usecase.GetCarEngine(userID, c.Param("car"))

	return c.JSON(http.StatusOK, &util.Response{Data: "request in processing"})
}
