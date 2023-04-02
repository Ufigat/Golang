package delivery

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserCars(c echo.Context) error {
	carID := c.QueryParam("id")

	resp, err := http.Get(fmt.Sprint("http://localhost:8083/user-cars?id=", carID))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, string(body[:]))
}

// func GetUserEngines(c echo.Context) error {
// 	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return c.JSON(http.StatusOK, response)
// }
