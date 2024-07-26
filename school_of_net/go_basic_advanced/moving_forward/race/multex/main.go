package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// * go run -race main.go
// Vai mostrar se o sistema tem alguma race condition

var m sync.Mutex // Protege dados que rodam em paralelo
var result int

func main() {
	go runProcess("P1", 10)
	go runProcess("P2", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()   // *
		result++   // Vai travar e deixar protegido ate atribuir o valor para evitar race condition
		m.Unlock() // *
		fmt.Println(name, "->", i, "Partial result: ", result)
	}
}
