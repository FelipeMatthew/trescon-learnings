package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Moto struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func AddMoto(c echo.Context) error {
	moto := Moto{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&moto)
	if err != nil {
		log.Printf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your moto: $#v", &moto)
	return c.String(http.StatusOK, "Moto added successfuly")

}