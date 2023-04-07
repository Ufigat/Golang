package delivery

import (
	"car/internal/app/infrastructure/repository"
	"car/pkg/response/fault"
	"car/pkg/util"
	"net/http"

	carReq "car/pkg/request/car"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func PostCars(c echo.Context) error {
	var req carReq.IDsRequest
	err := c.Bind(&req)
	if err != nil {
		log.Errorln("PostCars #1", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarByUser(&req)
	if err != nil {
		log.Errorln("PostCars #2", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

func PostCarEngines(c echo.Context) error {
	var req carReq.IDsRequest
	err := c.Bind(&req)
	if err != nil {
		log.Errorln("PostCars #1", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarEngineByUser(&req)
	if err != nil {
		log.Errorln("PostCars #2", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

func GetCarsByBrand(c echo.Context) error {
	var req carReq.IDsRequest
	err := c.Bind(&req)
	if err != nil {
		log.Errorln("GetCarsByBrand #1", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	resp, err := repository.GetCarByUser(&req)
	if err != nil {
		log.Errorln("GetCarsByBrand #2", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	return c.JSON(http.StatusOK, &util.Response{Data: resp})
}

// func GetUserCarsWithEngines(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		log.Errorln("GetUserCarsWithEngines ", err.Error())

// 		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	var car domain.Car
// 	car.UserID = userID
// 	err = car.ValidationUserID()
// 	if err != nil {
// 		log.Infoln("GetUserCarsWithEngines ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	response, err := usecase.GetUserWithCarEngines(&car)
// 	if err != nil {
// 		log.Errorln("GetUserCarsWithEngines ", err.Error())

// 		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	return c.JSON(http.StatusOK, response)
// }

// func GetCarsWithEnginesByBrand(c echo.Context) error {
// 	var car domain.Car
// 	car.Brand = c.QueryParam("brand")

// 	err := car.ValidationBrand()
// 	if err != nil {
// 		log.Infoln("GetCarsWithEnginesByBrand ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	response, err := usecase.GetCarWithEnginesByBrand(&car)
// 	if err != nil {
// 		log.Errorln("GetCarsWithEnginesByBrand ", err.Error())

// 		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	return c.JSON(http.StatusOK, &util.Response{Data: response})
// }

// func GetCarEngine(c echo.Context) error {
// 	carID, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		log.Errorln("GetCarEngine ", err.Error())

// 		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	var car domain.Car
// 	car.ID = carID
// 	err = car.ValidationID()
// 	if err != nil {
// 		log.Infoln("GetCarEngine ", err.Error())

// 		return c.JSON(http.StatusUnprocessableEntity, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	response, err := usecase.GetCarEngine(&car)
// 	if err != nil {
// 		log.Errorln("GetCarEngine ", err.Error())

// 		return c.JSON(http.StatusBadRequest, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	return c.JSON(http.StatusOK, &util.Response{Data: response})
// }
