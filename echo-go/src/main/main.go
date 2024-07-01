package main

import (
	"fmt"
	"github.com/FelipeMatthew/go-learnings/src/functions"

	"github.com/labstack/echo"
)



func main() {
	fmt.Println("Working w echo")

	e := echo.New()

	e.GET("/", functions.GettingStart)

	e.Start(":8000")
}
