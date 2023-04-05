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

func CarEngines(engineRequest *engine.EnginesRequest) ([]engineRes.EngineResponse, error) {
	value, err := json.Marshal(engineRequest)
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

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("CarEngines ", err.Error())

			return nil, err
		}

		return nil, &fault
	}

	var cers []engineRes.EngineResponse
	err = json.NewDecoder(resp.Body).Decode(&cers)
	if err != nil {
		log.Errorln("CarEngines ", err.Error())

		return nil, err
	}

	return cers, nil
}

func CarEngine(engineRequest *engine.EngineRequest) (*engineRes.EngineResponse, error) {
	resp, err := http.Get(fmt.Sprint(viper.GetString("engineService"), "/engine?id=", engineRequest.EngineID))
	if err != nil {
		log.Errorln("CarEngine ", err.Error())

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		var fault fault.FaultResponse
		err = json.NewDecoder(resp.Body).Decode(&fault)
		if err != nil {
			log.Errorln("CarEngine ", err.Error())

			return nil, err
		}

		return nil, &fault
	}

	var cers engineRes.EngineResponse
	err = json.NewDecoder(resp.Body).Decode(&cers)
	if err != nil {
		log.Errorln("CarEngine ", err.Error())

		return nil, err
	}

	return &cers, nil
}
