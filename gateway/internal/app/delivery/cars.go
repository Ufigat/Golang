package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/car"
	"gateway/pkg/response/fault"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCarEnginesByBrand(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/car-engines-brand?brand=", c.QueryParam("brand")))
	if err != nil {
		log.Println("GetUserCars microservice error", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetCarEngineByBrand ", err.Error())
			return echo.ErrBadRequest
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var crs car.CarBrandWithEngineResponse
	err = json.NewDecoder(resp.Body).Decode(&crs)
	if err != nil {
		log.Println("GetUserEngines ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, crs)
}

func GetCarEngine(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/car-engines?id=", c.QueryParam("id")))
	if err != nil {
		log.Println("GetCarEngine microservice error", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Println("GetCarEngine ", err.Error())
			return echo.ErrBadRequest
		}

		return c.JSON(resp.StatusCode, fault)
	}

	var crs car.CarWithEngineResponse
	err = json.NewDecoder(resp.Body).Decode(&crs)
	if err != nil {
		log.Println("GetUserEngine ", err.Error())
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, crs)
}
