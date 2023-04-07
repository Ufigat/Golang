package delivery

import (
	"fmt"
	"net/http"
	"strconv"
	"user/internal/app/domain"
	"user/internal/app/infrastructure/repository"
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

	user := &domain.User{
		ID: userID,
	}

	err = user.ValidationID()
	if err != nil {
		log.Infoln("GetUserCars #2", err.Error())

		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetUser(user)
	if err != nil {
		log.Infoln("GetUserCars #3", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	fmt.Println("resp resp resp", resp)

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

// func GetUserCarEngines(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		log.Errorln("GetUserCars ", err.Error())

// 		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	var user domain.User
// 	user.ID = userID
// 	err = user.ValidationID()
// 	if err != nil {
// 		log.Infoln("GetUserCars after valid ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	resp, err := usecase.GetUserWithCarEngines(&user)
// 	if err != nil {
// 		log.Errorln("GetUserCars after GetUserWithCarEngines ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	return c.JSON(http.StatusOK, &util.Response{Data: resp})
// }
