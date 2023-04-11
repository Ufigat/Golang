package delivery

import (
	"net/http"
	"strconv"
	"user/internal/app/infrastructure/repository"
	"user/pkg/request/user"
	"user/pkg/response/fault"
	"user/pkg/util"

	"github.com/labstack/echo/v4"

	log "github.com/sirupsen/logrus"
)

func GetUserCars(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Errorln("GetUserCars #1", err.Error())

		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	err = user.ValidationID(userID)
	if err != nil {
		log.Infoln("GetUserCars #2", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetUser(userID)
	if err != nil {
		log.Infoln("GetUserCars #3", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}
