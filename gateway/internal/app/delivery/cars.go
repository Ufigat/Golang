package delivery

import (
	"github.com/labstack/echo/v4"
)

func GetCarEnginesByBrand(c echo.Context) error {
	// resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/engines-brand?brand=", c.Param("brand")))
	// if err != nil {
	// 	log.Errorln("GetCarEnginesByBrand microservice error ", err.Error())

	// 	return echo.ErrInternalServerError
	// }

	// defer resp.Body.Close()

	// err = json.NewDecoder(resp.Body).Decode(&util.Response)
	// if err != nil {
	// 	log.Errorln("GetUserEngines ", err.Error())

	// 	return echo.ErrInternalServerError
	// }

	// return c.JSON(http.StatusOK, answer)
	return nil
}

func GetCarEngine(c echo.Context) error {
	// resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/engines?id=", c.Param("id")))
	// if err != nil {
	// 	log.Errorln("GetCarEngine microservice error", err.Error())

	// 	return echo.ErrInternalServerError
	// }

	// defer resp.Body.Close()

	// var carEngineResp car.EngineResponse

	// err = json.NewDecoder(resp.Body).Decode(&carEngineResp)
	// if err != nil {
	// 	log.Errorln("GetUserEngine ", err.Error())

	// 	return echo.ErrInternalServerError
	// }

	// return c.JSON(http.StatusOK, carEngineResp)
	return nil
}
