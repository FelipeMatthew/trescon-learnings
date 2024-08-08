package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Car struct {
	Name string `json:"name"`
	Year int    `json:"year"`
}

func AddCar(c echo.Context) error {
	car := Car{}

	defer c.Request().Body.Close()

	// Body
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &car)
	if err != nil {
		log.Printf("Failed unmarhsaling the request body %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your car: $#v", &car)
	return c.String(http.StatusOK, "Car added successfuly")
}