package delivery

import (
	"encoding/json"
	"user/internal/app/infrastructure/repository"
	"user/pkg/rabbitmq"
	"user/pkg/request/user"
	"user/pkg/response/fault"
	"user/pkg/util"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Conn *rabbitmq.Connect
}

func (u *User) GetUserCars() {
	for d := range u.Conn.QueueChannel["GetUserCars"].DeliveryChan {
		var userID int

		err := json.Unmarshal(d.Body, &userID)
		if err != nil {
			log.Errorln("GetUserCars #1 ", err.Error())

			err = u.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendUserCars",
				"SendUserCars",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetUserCars #2 ", err.Error())

				break
			}
			continue
		}

		err = user.ValidationID(userID)
		if err != nil {
			log.Infoln("GetUserCars #3 ", err.Error())

			err = u.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendUserCars",
				"SendUserCars",
				"",
				false,
				false,
				"text/plain",
			)
			if err != nil {
				log.Fatalf("GetUserCars #4 ", err.Error())

				break
			}
			continue
		}

		resp, err := repository.GetUser(userID)
		if err != nil {
			log.Errorln("GetUserCars #5 ", err.Error())

			err = u.Conn.ProduceMessage(&util.Response{Error: fault.NewResponse(err.Error())},
				"SendUserCars",
				"SendUserCars",
				"",
				false,
				false,
				"text/plain")
			if err != nil {
				log.Fatalf("GetUserCars #6 ", err.Error())

				break
			}
			continue
		}

		err = u.Conn.ProduceMessage(&util.Response{Data: resp},
			"SendUserCars",
			"SendUserCars",
			"",
			false,
			false,
			"text/plain",
		)
		if err != nil {
			log.Fatalf("GetUserCars #7 ", err.Error())

			break
		}
	}
}
