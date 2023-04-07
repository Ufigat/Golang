package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetCarEnginesByBrand(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/engines-brand?brand=", c.Param("brand")))
	if err != nil {
		log.Errorln("GetCarEnginesByBrand microservice error ", err.Error())

		return echo.ErrInternalServerError
	}

	defer resp.Body.Close()

	var brandEngineResp util.Response

	err = json.NewDecoder(resp.Body).Decode(&brandEngineResp)
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(resp.StatusCode, brandEngineResp)
}

func GetCarEngine(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/engines?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetCarEngine microservice error", err.Error())

		return echo.ErrInternalServerError
	}

	defer resp.Body.Close()

	var carEngineResp util.Response

	err = json.NewDecoder(resp.Body).Decode(&carEngineResp)
	if err != nil {
		log.Errorln("GetUserEngine ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(resp.StatusCode, carEngineResp)
}
