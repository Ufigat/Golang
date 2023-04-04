package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/fault"
	"gateway/pkg/response/user"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserCars(c echo.Context) error {
	userID := c.QueryParam("id")
	resp, err := http.Get(fmt.Sprint("http://localhost:8083/user-cars?id=", userID))
	if err != nil {
		log.Println("GetUserCars ", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetUserCars ", err.Error())
			return echo.ErrBadRequest
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var crs user.UserCarsResponse
	err = json.NewDecoder(resp.Body).Decode(&crs)
	if err != nil {
		log.Println("GetUserEngines ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, crs)
}

func GetUserEngines(c echo.Context) error {
	userId := c.QueryParam("id")
	resp, err := http.Get(fmt.Sprint("http://localhost:8083/user-cars-engine?id=", userId))
	if err != nil {
		log.Println("GetUserEngines ", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetUserEngines ", err.Error())
			return echo.ErrBadRequest
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var uer user.UserEnginesResponse
	err = json.NewDecoder(resp.Body).Decode(&uer)
	if err != nil {
		log.Println("GetUserEngines ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, uer)
}
