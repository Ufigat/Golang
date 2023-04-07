package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetUserCars(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("userService"), "/user/cars?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	defer resp.Body.Close()

	var userCarsResp util.Response

	err = json.NewDecoder(resp.Body).Decode(&userCarsResp)
	if err != nil {
		log.Errorln("GetUserCars ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	if userCarsResp.Error != nil {
		log.Errorln("GetUserCars ", userCarsResp.Error)

		return c.JSON(http.StatusUnprocessableEntity, userCarsResp)
	}

	return c.JSON(resp.StatusCode, userCarsResp)
}

func GetUserEngines(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("userService"), "/user/cars-engine?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	defer resp.Body.Close()

	var userEnginesResp util.Response

	err = json.NewDecoder(resp.Body).Decode(&userEnginesResp)
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	if userEnginesResp.Error != nil {
		log.Errorln("GetUserCars ", userEnginesResp.Error)

		return c.JSON(http.StatusUnprocessableEntity, userEnginesResp)
	}

	return c.JSON(resp.StatusCode, userEnginesResp)
}
