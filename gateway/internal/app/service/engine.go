package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	engineRes "gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetEngines(engineIDs []int) (*engineRes.DataResponse, error) {
	value, err := json.Marshal(&engineIDs)
	if err != nil {
		log.Errorln("CarEngines #1 ", err.Error())

		return nil, &fault.Response{Message: err.Error()}
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(fmt.Sprint(viper.GetString("services.engine"), "/engines"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("CarEngines #2 ", err.Error())

		return nil, &fault.Response{Message: err.Error()}
	}

	defer resp.Body.Close()

	var carEnigneResp engineRes.DataResponse

	err = json.NewDecoder(resp.Body).Decode(&carEnigneResp)
	if err != nil {
		log.Errorln("CarEngines #3 ", err.Error())

		return nil, &fault.Response{Message: err.Error()}
	}

	if carEnigneResp.Error != nil {
		log.Errorln("CarEngines #4 ", carEnigneResp.Error.Message)

		return nil, &fault.Response{Message: carEnigneResp.Error.Message}
	}

	return &carEnigneResp, nil
}

func GetEngine(engineId int) (*engineRes.EnigneResponse, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(fmt.Sprint(viper.GetString("services.engine"), "/engine?id=", engineId))
	if err != nil {
		log.Errorln("CarEngine #1 ", err.Error())

		return nil, &fault.Response{Message: err.Error()}
	}

	defer resp.Body.Close()

	var dataResp engineRes.EnigneResponse

	err = json.NewDecoder(resp.Body).Decode(&dataResp)
	if err != nil {
		log.Errorln("CarEngine #2 ", err.Error())

		return nil, &fault.Response{Message: err.Error()}
	}

	if dataResp.Error != nil {
		log.Errorln("CarEngine #3 ", dataResp.Error.Message)

		return nil, &fault.Response{Message: dataResp.Error.Message}
	}

	return &dataResp, nil
}
