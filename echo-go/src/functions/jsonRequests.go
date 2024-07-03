package functions

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
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


type Moto struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
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

type Dog struct {
	Name string `json:"name"`
	Breed string  `json:"breed"`
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