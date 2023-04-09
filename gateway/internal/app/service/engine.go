package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	engineReq "gateway/pkg/request/engine"
	engineRes "gateway/pkg/response/engine"
	"gateway/pkg/response/fault"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CarEngines(engineIDs []int) (*engineRes.DataResponse, error) {

	engineIDsReq := &engineReq.EnginesRequest{EngineIDs: engineIDs}
	value, err := json.Marshal(&engineIDsReq)
	if err != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, err
	}

	resp, err := http.Post(fmt.Sprint(viper.GetString("engineService"), "/engines"), "application/json", bytes.NewBuffer(value))
	if err != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var carEnigneRespLinks engineRes.DataResponse

	err = json.NewDecoder(resp.Body).Decode(&carEnigneRespLinks)
	if err != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, err
	}

	if carEnigneRespLinks.Error != "" {
		log.Errorln("CarEngines ", err.Error())

		return nil, &fault.Response{Message: carEnigneRespLinks.Error}
	}

	return &carEnigneRespLinks, nil
}

func CarEngine(engineId int) (*engineRes.EnigneResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("engineService"), "/engine?id=", engineId))
	if err != nil {
		log.Errorln("CarEngine ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var dataResp engineRes.EnigneResponse

	err = json.NewDecoder(resp.Body).Decode(&dataResp)
	if err != nil {
		log.Errorln("CarEngine ", err.Error())

		return nil, err
	}

	if dataResp.Error != "" {
		log.Errorln("CarEngine ", err.Error())

		return nil, &fault.Response{Message: dataResp.Error}
	}

	return &dataResp, nil
}
