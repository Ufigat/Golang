package main

import (
	"car/routing"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routing.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
