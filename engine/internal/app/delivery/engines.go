package delivery

import (
	"encoding/json"
	"engine/internal/app/usecase"
	"engine/pkg/request/engine"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostEngineUserCars(c echo.Context) error {
	var ucfer engine.UserCarsForEngineRequest
	json.NewDecoder(c.Request().Body).Decode(&ucfer)

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
