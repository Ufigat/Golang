package delivery

import (
	"net/http"
	"strconv"
	"user/internal/app/domain"
	"user/internal/app/usecase"
	"user/pkg/response/fault"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func GetUserCars(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return echo.ErrBadRequest
	}

	var user domain.User
	user.ID = userID
	err = user.ValidationID()
	if err != nil {
		log.Infoln("GetUserCars ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	resp, err := usecase.GetUserWithCar(&user)
	if err != nil {
		log.Errorln("GetUserCars after GetUserWithCar ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func GetUserCarEngines(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return echo.ErrBadRequest
	}

	var user domain.User
	user.ID = userID
	err = user.ValidationID()
	if err != nil {
		log.Infoln("GetUserCars after valid ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	resp, err := usecase.GetUserWithCarEngines(&user)
	if err != nil {
		log.Errorln("GetUserCars after GetUserWithCarEngines ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, resp)
}
