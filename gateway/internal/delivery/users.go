package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/user"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserCars(c echo.Context) error {
	userID := c.QueryParam("id")

	resp, err := http.Get(fmt.Sprint("http://localhost:8083/user-cars?id=", userID))
	if err != nil {
		log.Println("GetUserCars user microservice error", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	var crs user.UserCarsResponse
	json.NewDecoder(resp.Body).Decode(&crs)

	return c.JSON(http.StatusOK, crs)
}

func GetUserEngines(c echo.Context) error {
	userId := c.QueryParam("id")
	resp, err := http.Get(fmt.Sprint("http://localhost:8083/user-cars-engine?id=", userId))
	if err != nil {
		log.Println("Error from query", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	var uer user.UserEnginesResponse
	json.NewDecoder(resp.Body).Decode(&uer)

	return c.JSON(http.StatusOK, uer)
}
