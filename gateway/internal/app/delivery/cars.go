package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/car"
	"gateway/pkg/response/fault"
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

	if resp.StatusCode >= http.StatusBadRequest {
		var fault fault.FaultResponse

		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("GetCarEngineByBrand ", err.Error())

			return echo.ErrInternalServerError
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var carBrandResp car.CarBrandWithEngineResponse

	err = json.NewDecoder(resp.Body).Decode(&carBrandResp)
	if err != nil {
		log.Errorln("GetUserEngines ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, carBrandResp)
}

func GetCarEngine(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/engines?id=", c.Param("id")))
	if err != nil {
		log.Errorln("GetCarEngine microservice error", err.Error())

		return echo.ErrInternalServerError
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse

		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("GetCarEngine ", err.Error())

			return echo.ErrInternalServerError
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var carEngineResp car.CarWithEngineResponse

	err = json.NewDecoder(resp.Body).Decode(&carEngineResp)
	if err != nil {
		log.Errorln("GetUserEngine ", err.Error())

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, carEngineResp)
}
