package delivery

import (
	"gateway/internal/app/usecase"
	"gateway/pkg/rabbitmq"
	"gateway/pkg/response/fault"
	"gateway/pkg/util"
	"gateway/pkg/websocket"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Conn *rabbitmq.Connect
	Room *websocket.Room
}

// func (car *Car) GetCarEnginesByBrand(c echo.Context) error {
// 	userID, err := strconv.Atoi(c.Param("client"))
// 	if err != nil {
// 		log.Errorln("GetCarEnginesByBrand #1 ", err.Error())

// 		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
// 	}

// 	usecase := usecase.Usacase{Room: car.Room, Conn: car.Conn}
// 	go usecase.GetCarEnginesByBrand(userID, c.Param("brand"))

//		return c.JSON(http.StatusOK, &util.Response{Data: "request in processing"})
//	}
func (u *User) GetUserCars(c echo.Context) error {
	clientID, err := strconv.Atoi(c.Param("client"))
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	usecase := usecase.Usacase{Room: u.Room, Conn: u.Conn}
	go usecase.GetUserCars(clientID, c.Param("id"))

	return c.JSON(http.StatusOK, &util.Response{Data: "request in processing"})

	// userDataResp, err := service.GetUser(c.Param("id"))
	// if err != nil {
	// 	log.Errorln("GetUserCars #1 ", err.Error())

	// 	return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	// }

	// carDataResp, err := service.GetCars(userDataResp.Data.CarIDs)
	// if err != nil {
	// 	log.Errorln("GetUserCars #2 ", err.Error())

	// }

	// userCarsResp := &user.UserCarsResponse{
	// 	ID:   userDataResp.Data.ID,
	// 	Name: userDataResp.Data.Name,
	// 	Cars: carDataResp.Data,
	// }

	// return c.JSON(http.StatusOK, &util.Response{Data: userCarsResp})
}

func (u *User) GetUserEngines(c echo.Context) error {
	clientID, err := strconv.Atoi(c.Param("client"))
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	}

	usecase := usecase.Usacase{Room: u.Room, Conn: u.Conn}
	go usecase.GetUserEngines(clientID, c.Param("id"))

	return c.JSON(http.StatusOK, &util.Response{Data: "request in processing"})

	// carEngineDataResp, err := service.GetCarsEngine(userDataResp.Data.CarIDs)
	// if err != nil {
	// 	log.Errorln("GetUserCars #2 ", err.Error())

	// 	return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	// }

	// var enginesIDs []int

	// for i := range carEngineDataResp.Data {
	// 	enginesIDs = append(enginesIDs, carEngineDataResp.Data[i].EngineID)
	// }

	// enginesResp, err := service.GetEngines(enginesIDs)
	// if err != nil {
	// 	log.Errorln("GetUserCars #2 ", err.Error())

	// 	return c.JSON(http.StatusInternalServerError, &util.Response{Error: fault.NewResponse(err.Error())})
	// }

	// userEnginesResp := &user.UserEnginesResponse{
	// 	ID:      userDataResp.Data.ID,
	// 	Name:    userDataResp.Data.Name,
	// 	Engines: enginesResp.Data,
	// }

	// return c.JSON(http.StatusOK, &util.Response{Data: userEnginesResp})
}
