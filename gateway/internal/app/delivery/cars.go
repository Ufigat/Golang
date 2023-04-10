package delivery

import (
	"gateway/internal/app/service"
	"gateway/pkg/response/car"
	"gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func GetCarEnginesByBrand(c echo.Context) error {
	carResp, err := service.GetCarsEngineByBrand(c.Param("brand"))
	if err != nil {
		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	var engineIDs []int

	for i := range carResp.Data {
		engineIDs = append(engineIDs, carResp.Data[i].EngineID)
	}

	engineDataResp, err := service.PostEngines(engineIDs)
	if err != nil {
		log.Errorln("GetCarEnginesByBrand #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	carEngineByBrandResp := &car.EngineByBrandResponse{
		Name:    c.Param("brand"),
		Engines: engineDataResp.Data,
	}

	return c.JSON(http.StatusOK, &util.Response{Data: carEngineByBrandResp})
}

func GetCarEngine(c echo.Context) error {
	carResp, err := service.GetCar(c.Param("id"))
	if err != nil {
		log.Errorln("GetCarEngine #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	engineResp, err := service.GetEngine(carResp.Data.EngineID)
	if err != nil {
		log.Errorln("GetCarEngine #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp := &engine.ForCar{
		ID:     carResp.Data.ID,
		Engine: engineResp.Data,
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}
