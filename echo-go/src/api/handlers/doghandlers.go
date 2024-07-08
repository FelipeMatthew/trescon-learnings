package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
}

func AddDog(c echo.Context) error {
	dog := Dog{}

	err := c.Bind(&dog)
	if err != nil {
		log.Printf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: $#v", &dog)
	return c.String(http.StatusOK, "Dog added successfuly")
}