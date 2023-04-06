package car

import (
	"encoding/json"
	"fmt"

	"net/http"
	"user/internal/app/domain"
	"user/pkg/response/car"
	"user/pkg/response/engine"
	"user/pkg/response/fault"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetCars(userModel *domain.User) (*car.DataResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/cars?id=", userModel.ID))
	if err != nil {
		log.Errorln("GetCars ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var carResp car.DataResponse

	err = json.NewDecoder(resp.Body).Decode(&carResp)
	if err != nil {
		log.Errorln("GetCars ", err.Error())

		return nil, err
	}

	if carResp.Error != "" {
		return nil, &fault.Response{Message: carResp.Error}
	}

	return &carResp, nil
}

func GetCarsWithEngine(userModel *domain.User) (*engine.LinksResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("carService"), "/car/user-engines?id=", userModel.ID))
	if err != nil {
		log.Errorln("GetCarsWithEngine ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var engineRespLinks engine.LinksResponse

	err = json.NewDecoder(resp.Body).Decode(&engineRespLinks)
	if err != nil {
		log.Errorln("GetCarsWithEngine ", err.Error())

		return nil, err
	}

	if engineRespLinks.Error != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, engineRespLinks.Error
	}

	return &engineRespLinks, nil
}
