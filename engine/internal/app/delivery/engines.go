package delivery

import (
	"engine/internal/app/domain"
	"engine/internal/app/usecase"
	"engine/pkg/request/engine"
	engineResp "engine/pkg/response/engine"
	"engine/pkg/response/fault"
	"engine/pkg/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func PostEngineUserCars(c echo.Context) error {
	var engineIdsReq engine.IDsRequest

	err := c.Bind(&engineIdsReq)
	if err != nil {
		log.Errorln("PostEngineUserCars ", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	for i := range engineIdsReq.EngineID {
		if engineIdsReq.EngineID[i] <= 0 {
			return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse("Id is not valid")})
		}
	}

	response, err := usecase.GetEngines(&engineIdsReq)
	if err != nil {
		log.Errorln("PostEngineUserCars ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	fmt.Println(&util.Response{Data: response})

	return c.JSON(http.StatusOK, &util.Response{Data: response})
}

func GetEngine(c echo.Context) error {
	engineID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetEngine", err.Error())

		return c.JSON(http.StatusBadRequest, &engineResp.Response{Error: fault.NewResponse(err.Error())})
	}

	var engineModel domain.Engine
	engineModel.ID = engineID
	err = engineModel.ValidationID()

	if err != nil {
		log.Infoln("GetEngine ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &engineResp.Response{Error: fault.NewResponse(err.Error())})
	}

	response, err := usecase.GetEngine(&engineModel)
	if err != nil {
		log.Errorln("GetEngine ", err.Error())

		return c.JSON(http.StatusInternalServerError, &engineResp.Response{Engine: response, Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, response)
}
