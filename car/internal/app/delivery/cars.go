package delivery

import (
	"car/internal/app/infrastructure/repository"
	"car/pkg/rabbitmq"
	"car/pkg/response/fault"
	"car/pkg/util"
	"net/http"
	"strconv"

	carReq "car/pkg/request/car"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Delivery struct {
	Conn *rabbitmq.Connect
}

func GetCars(c echo.Context) error {
	var carsIDs []int
	err := c.Bind(&carsIDs)
	if err != nil {
		log.Errorln("GetCars #1 ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarByUser(carsIDs)
	if err != nil {
		log.Errorln("GetCars #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

func GetCarEngines(c echo.Context) error {
	var carsIDs []int
	err := c.Bind(&carsIDs)
	if err != nil {
		log.Errorln("GetCarEngines #1 ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarEngineByUser(carsIDs)
	if err != nil {
		log.Errorln("GetCarEngines #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

func (d *Delivery) GetCarsByBrand() {
	for dc := range d.Conn.QueueChannel["GetCar"].DeliveryChan {

		req := &carReq.Request{
			Brand: string(dc.Body),
		}

		err := req.ValidationBrand()
		if err != nil {
			log.Errorln("GetCarsByBrand #1 ", err.Error())

			d.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		resp, err := repository.GetCarEngineByBrand(req)
		if err != nil {
			log.Errorln("GetCarsByBrand #2 ", err.Error())

			d.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())})
			continue
		}

		d.ProduceMessage(&util.Response{Data: resp})
	}
}

func GetCarEngine(c echo.Context) error {
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorln("GetCarEngine #1 ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	req := &carReq.Request{
		ID: carID,
	}

	err = req.ValidationID()
	if err != nil {
		log.Infoln("GetCarEngine #2 ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarEngine(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}
