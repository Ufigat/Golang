package engine

import (
	"bytes"
	"car/pkg/request/engine"
	engineRes "car/pkg/response/engine"
	"car/pkg/response/fault"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CarEngines(engineRequest *engine.EnginesRequest) (*engineRes.DataResponse, error) {
	value, err := json.Marshal(engineRequest)
	if err != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, err
	}
	fmt.Println(fmt.Sprint(viper.GetString("engineService"), "/engines"))
	fmt.Println(*engineRequest)
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

	fmt.Println(carEnigneRespLinks)
	return &carEnigneRespLinks, nil
}

func CarEngine(engineRequest *engine.EngineRequest) (*engineRes.DataResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("engineService"), "/engine?id=", engineRequest.EngineID))
	if err != nil {
		log.Errorln("CarEngine ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	var dataResp engineRes.DataResponse

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
