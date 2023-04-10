package delivery

import (
	"engine/internal/app/infrastructure/repository"
	"engine/pkg/request/engine"
	engineResp "engine/pkg/response/engine"
	"engine/pkg/response/fault"
	"engine/pkg/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func PostEngineUserCars(c echo.Context) error {
	var engineIDsReq engine.IDsRequest
	err := c.Bind(&engineIDsReq)
	if err != nil {
		log.Errorln("PostEngineUserCars #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	for i := range engineIDsReq.EngineID {
		if engineIDsReq.EngineID[i] <= 0 {

			return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse("ID is not valid")})
		}
	}

	response, err := repository.GetEngines(&engineIDsReq)
	if err != nil {
		log.Errorln("PostEngineUserCars #2 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: response})
}

func GetEngine(c echo.Context) error {
	engineID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		log.Errorln("GetEngine #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &engineResp.Response{Error: fault.NewResponse(err.Error())})
	}

	engineReq := &engine.Request{
		ID: engineID,
	}

	err = engineReq.ValidationID()
	if err != nil {
		log.Infoln("GetEngine #2 ", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &engineResp.Response{Error: fault.NewResponse(err.Error())})
	}

	response, err := repository.GetEngine(engineReq)
	if err != nil {
		log.Errorln("GetEngine #3 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &engineResp.Response{Engine: response, Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: response})
}
