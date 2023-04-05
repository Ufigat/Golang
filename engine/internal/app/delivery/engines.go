package delivery

import (
	"engine/internal/app/domain"
	"engine/internal/app/usecase"
	"engine/pkg/request/engine"
	"engine/pkg/response/fault"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func PostEngineUserCars(c echo.Context) error {
	var ucfer engine.UserCarsForEnginesRequest

	err := c.Bind(&ucfer)
	if err != nil {
		log.Errorln("PostEngineUserCars ", err.Error())

		return echo.ErrBadRequest
	}

	for i := range ucfer.EngineID {
		if ucfer.EngineID[i] <= 0 {
			return echo.ErrBadRequest
		}
	}

	response, err := usecase.GetEngines(&ucfer)
	if err != nil {
		log.Errorln("PostEngineUserCars ", err.Error())

		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetEngine(c echo.Context) error {
	engineID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetEngine", err.Error())

		return echo.ErrBadRequest
	}

	var em domain.Engine
	em.ID = engineID
	err = em.ValidationID()
	if err != nil {
		log.Infoln("GetEngine ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetEngine(&em)
	if err != nil {
		log.Errorln("GetEngine ", err.Error())

		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}
