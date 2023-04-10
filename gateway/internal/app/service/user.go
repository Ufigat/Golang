package service

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/fault"
	"gateway/pkg/response/user"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetUser(userID string) (*user.DataResponse, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	httpResp, err := client.Get(fmt.Sprint(viper.GetString("services.user"), "/user/", userID, "/cars"))
	if err != nil {
		log.Errorln("GetUserCars #1 ", err.Error())

		return nil, err
	}

	defer httpResp.Body.Close()

	var resp user.DataResponse

	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		log.Errorln("GetUserCars #2 ", err.Error())

		return nil, err
	}

	if resp.Error != nil {
		log.Errorln("GetUserCars #3 ", resp.Error.Message)

		return nil, &fault.Response{Message: resp.Error.Message}
	}

	return &resp, nil
}
