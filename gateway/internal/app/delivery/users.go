package delivery

import (
	"gateway/internal/app/service"
	"gateway/pkg/response/fault"
	"gateway/pkg/response/user"
	"gateway/pkg/util"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func GetUserCars(c echo.Context) error {
	userDataResp, err := service.GetUser(c.Param("id"))
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	var carIDs []int

	for i := range userDataResp.Data {
		carIDs = append(carIDs, userDataResp.Data[i].CarID)
	}

	carDataResp, err := service.GetCars(carIDs)
	if err != nil {
		log.Errorln("GetUserCars #2 ", err.Error())

	}

	userCarsResp := &user.UserCarsResponse{
		ID:   userDataResp.Data[0].ID,
		Name: userDataResp.Data[0].Name,
		Cars: carDataResp.Data}

	return c.JSON(http.StatusOK, &util.Response{Data: userCarsResp})
}

func GetUserEngines(c echo.Context) error {
	userDataResp, err := service.GetUser(c.Param("id"))
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	var carIDs []int
	for i := range userDataResp.Data {
		carIDs = append(carIDs, userDataResp.Data[i].CarID)
	}

	carEngineDataResp, err := service.GetCarsEngine(carIDs)
	if err != nil {
		log.Errorln("GetUserCars #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	var enginesIDs []int

	for i := range carEngineDataResp.Data {
		enginesIDs = append(enginesIDs, carEngineDataResp.Data[i].EngineID)
	}

	enginesResp, err := service.PostEngines(enginesIDs)
	if err != nil {
		log.Errorln("GetUserCars #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	userEnginesResp := &user.UserEnginesResponse{
		ID:      userDataResp.Data[0].ID,
		Name:    userDataResp.Data[0].Name,
		Engines: enginesResp.Data}

	return c.JSON(http.StatusOK, &util.Response{Data: userEnginesResp})
}
