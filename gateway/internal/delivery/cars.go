package delivery

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/car"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCarEngineByBrand(c echo.Context) error {
	resp, err := http.Get(fmt.Sprint("http://localhost:8081/car-engines-brand?brand=", c.QueryParam("brand")))
	if err != nil {
		log.Println("GetUserCars user microservice error", err.Error())
		return echo.ErrBadRequest
	}

	defer resp.Body.Close()

	var crs car.CarResponseWithEngineByBrand
	json.NewDecoder(resp.Body).Decode(&crs)

	return c.JSON(http.StatusOK, crs)
}
