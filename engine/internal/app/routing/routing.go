package routing

import (
	"engine/internal/app/delivery"
	database "engine/pkg/postgres"
	"engine/pkg/rabbitmq"
	"log"
)

func Init() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("fatal DB connect error: %s", err.Error())
	}

	conn := rabbitmq.NewConnect()

	err = rabbitmq.ConnRabbit(conn)
	if err != nil {
		log.Fatalf("fatal rabbitmq connect error: %s", err.Error())
	}

	mes := conn.ConsumeEnginesChan()

	go delivery.GetEngines(conn, mes)
	//go delivery.GetEngine(conn, mes)
}
