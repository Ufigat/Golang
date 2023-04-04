package delivery

import (
	"encoding/json"
	"engine/internal/app/domain"
	"engine/internal/app/usecase"
	"engine/pkg/request/engine"
	"engine/pkg/response/fault"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PostEngineUserCars(c echo.Context) error {
	var ucfer engine.UserCarsForEnginesRequest
	err := json.NewDecoder(c.Request().Body).Decode(&ucfer)
	if err != nil {
		log.Println("PostEngineUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	for _, Id := range ucfer.EngineID {
		if Id <= 0 {
			return echo.ErrBadRequest
		}
	}

	response, err := usecase.GetEngines(&ucfer)
	if err != nil {
		log.Println("PostEngineUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}

func GetEngine(c echo.Context) error {
	engineID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Println("GetEngine", err.Error())
		return echo.ErrBadRequest
	}

	var em domain.Engine
	em.ID = engineID
	err = em.ValidationID()
	if err != nil {
		log.Println("GetEngine ", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, fault.NewFaultResponse(err.Error()))
	}

	response, err := usecase.GetEngine(&em)
	if err != nil {
		log.Println("GetEngine ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, response)
}
