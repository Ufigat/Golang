package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/fault"
	"gateway/pkg/response/user"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetUserCars(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("userService"), "/user/cars?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return echo.ErrInternalServerError
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("GetUserCars ", err.Error())

			return echo.ErrInternalServerError
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var userCarsResp user.UserCarsResponse

	err = json.NewDecoder(resp.Body).Decode(&userCarsResp)
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, userCarsResp)
}

func GetUserEngines(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("userService"), "/user/cars-engine?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return echo.ErrInternalServerError
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse

		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("GetUserEngines ", err.Error())

			return echo.ErrInternalServerError
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var userEnginesResp user.UserEnginesResponse

	err = json.NewDecoder(resp.Body).Decode(&userEnginesResp)
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, userEnginesResp)
}
