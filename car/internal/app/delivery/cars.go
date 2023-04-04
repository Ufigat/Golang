package delivery

import (
	"car/internal/app/domain"
	"car/internal/app/usecase"
	"car/pkg/response/fault"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCars(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetCars ", err.Error())
		return echo.ErrBadRequest
	}

	var cm domain.Car
	cm.UserID = userID
	err = cm.ValidationUserID()
	if err != nil {
		log.Println("GetCars ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetUserWithCar(&cm)
	if err != nil {
		log.Println("GetCars ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetUserCarsWithEngines(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetUserCarsWithEngines ", err.Error())
		return echo.ErrBadRequest
	}

	var cm domain.Car
	cm.UserID = userID
	err = cm.ValidationUserID()
	if err != nil {
		log.Println("GetUserCarsWithEngines ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetUserWithCarWithEngines(&cm)
	if err != nil {
		log.Println("GetUserCarsWithEngines ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetCarsWithEnginesByBrand(c echo.Context) error {
	var cm domain.Car
	cm.Brand = c.QueryParam("brand")

	err := cm.ValidationBrand()
	if err != nil {
		log.Println("GetCarsWithEnginesByBrand ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetCarWithEnginesByBrand(&cm)
	if err != nil {
		log.Println("GetCarsWithEnginesByBrand ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetCarEngine(c echo.Context) error {
	carID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetCarEngine ", err.Error())
		return echo.ErrBadRequest
	}

	var cm domain.Car
	cm.ID = carID
	err = cm.ValidationID()
	if err != nil {
		log.Println("GetCarEngine ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetCarEngine(&cm)
	if err != nil {
		log.Println("GetCarEngine ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}
