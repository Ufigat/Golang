package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectDB() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("host"), viper.GetInt("port"), viper.GetString("user"),
		viper.GetString("password"), viper.GetString("dbname"))

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("Successfully connected DB!")

	return nil
}
