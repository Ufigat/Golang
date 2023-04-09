package service

import (
	"encoding/json"
	"fmt"
	"gateway/pkg/response/user"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetUser(userID string) (*user.DataResponse, error) {

	httpResp, err := http.Get(fmt.Sprint(viper.GetString("userService"), "/user/", userID, "/cars"))
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

	return &resp, nil
}
