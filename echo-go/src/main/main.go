package main

import (
	"fmt"
	
	"github.com/FelipeMatthew/go-learnings/src/router"
)

func main() {
	fmt.Println("Working w echo")

	e := router.New()

	e.Start(":8000")
}
