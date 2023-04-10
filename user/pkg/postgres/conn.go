package postgres

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("postgresql.host"), viper.GetInt("postgresql.port"), viper.GetString("postgresql.user"),
		viper.GetString("postgresql.password"), viper.GetString("postgresql.dbname"))

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Infoln("Successfully connected DB!")

	return nil
}
