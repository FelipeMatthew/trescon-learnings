package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Rodar um junto com o outro
	// Roda mais coisas ao mesmo tempo
	go runProcess("P1", 10)
	go runProcess("P2", 10)

	var s string
	// Vai escanear e se tiver mudança vai agir
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
