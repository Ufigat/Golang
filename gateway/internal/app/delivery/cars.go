package delivery

import (
	"encoding/json"
	"gateway/internal/app/service"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Delivery struct {
	Conn *rabbitmq.Connect
}

func (d *Delivery) GetCarEnginesByBrand(c echo.Context) error {
	conn, err := websocket.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	client := &websocket.Client{Conn: conn, Send: make(chan car.DataResponse)}

	go client.WritePump()
	go client.ReadPump()

	d.Conn.ProduceGetCar(c.Param("brand"))

	msgSendCar := <-d.Conn.ConsumeSendCar()

	var carsByBrandResp car.DataResponse

	err = json.Unmarshal(msgSendCar.Body, &carsByBrandResp)
	if err != nil {
		log.Errorln("GetCarsEngineByBrand #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	if carsByBrandResp.Error != nil {
		log.Errorln("GetCarsEngineByBrand #3 ", carsByBrandResp.Error.Message)
		client.Send <- carsByBrandResp

		return nil
	}

	client.Send <- carsByBrandResp
	return nil
}

// func GetCarEnginesByBrand(c echo.Context) error {
// 	carResp, err := service.GetCarsEngineByBrand(c.Param("brand"))
// 	if err != nil {
// 		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

// 		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	var engineIDs []int

// 	for i := range carResp.Data {
// 		engineIDs = append(engineIDs, carResp.Data[i].EngineID)
// 	}

// 	engineDataResp, err := service.GetEngines(engineIDs)
// 	if err != nil {
// 		log.Errorln("GetCarEnginesByBrand #2 ", err.Error())

// 		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	carEngineByBrandResp := &engine.ByBrandResponse{
// 		Brand:   c.Param("brand"),
// 		Engines: engineDataResp.Data,
// 	}

// 	return c.JSON(http.StatusOK, &util.Response{Data: carEngineByBrandResp})
// }

func GetCarEngine(c echo.Context) error {
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

// carResp, err := service.GetCarsEngineByBrand(c.Param("brand"))
// if err != nil {
// 	log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

// 	return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// }

// var engineIDs []int

// for i := range carResp.Data {
// 	engineIDs = append(engineIDs, carResp.Data[i].EngineID)
// }

// engineDataResp, err := service.GetEngines(engineIDs)
// if err != nil {
// 	log.Errorln("GetCarEnginesByBrand #2 ", err.Error())

// 	return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// }

// carEngineByBrandResp := &engine.ByBrandResponse{
// 	Brand:   c.Param("brand"),
// 	Engines: engineDataResp.Data,
// }

//return c.JSON(http.StatusInternalServerError, &util.Response{Data: "successful websocket connection"})
