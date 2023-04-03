package delivery

import (
	"log"
	"net/http"
	"strconv"
	"user/internal/app/domain"
	"user/internal/app/usecase"
	"user/pkg/response/fault"

	"github.com/labstack/echo/v4"
)

func GetUserCars(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	var um domain.User
	um.ID = userID
	err = um.ValidationID()
	if err != nil {
		log.Println("GetUserCars ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetUserWithCar(&um)
	if err != nil {
		log.Println("GetUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetUserCarEngines(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	var um domain.User
	um.ID = userID
	err = um.ValidationID()
	if err != nil {
		log.Println("GetUserCars after valid ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetUserWithCarEngines(&um)
	if err != nil {
		log.Println("GetUserCars after GetUserWithCarEngines", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}
