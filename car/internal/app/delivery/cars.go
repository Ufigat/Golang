package delivery

import (
	"car/internal/app/domain"
	"car/internal/app/usecase"
	"car/pkg/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func GetCars(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		log.Errorln("GetCars ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: err.Error()})
	}

	var car domain.Car
	car.UserID = userID
	err = car.ValidationUserID()
	if err != nil {
		log.Infoln("GetCars ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: err.Error()})
	}

	response, err := usecase.GetUserWithCar(&car)
	if err != nil {
		log.Errorln("GetCars ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: response})
}

func GetUserCarsWithEngines(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetUserCarsWithEngines ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: err.Error()})
	}

	var car domain.Car
	car.UserID = userID
	err = car.ValidationUserID()
	if err != nil {
		log.Infoln("GetUserCarsWithEngines ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: err.Error()})
	}

	response, err := usecase.GetUserWithCarEngines(&car)
	if err != nil {
		log.Errorln("GetUserCarsWithEngines ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}

func GetCarsWithEnginesByBrand(c echo.Context) error {
	var car domain.Car
	car.Brand = c.QueryParam("brand")

	err := car.ValidationBrand()
	if err != nil {
		log.Infoln("GetCarsWithEnginesByBrand ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: err.Error()})
	}

	response, err := usecase.GetCarWithEnginesByBrand(&car)
	if err != nil {
		log.Errorln("GetCarsWithEnginesByBrand ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: response})
}

// func GetCarEngine(c echo.Context) error {
// 	carID, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		log.Errorln("GetCarEngine ", err.Error())

// 		return echo.ErrInternalServerError
// 	}

// 	var car domain.Car
// 	car.ID = carID
// 	err = car.ValidationID()
// 	if err != nil {
// 		log.Infoln("GetCarEngine ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, fault.NewResponse(err.Error()))
// 	}

// 	response, err := usecase.GetCarEngine(&car)
// 	if err != nil {
// 		log.Errorln("GetCarEngine ", err.Error())

// 		return echo.ErrBadRequest
// 	}

// 	return c.JSON(http.StatusOK, &util.Response{Data: response})
// }
