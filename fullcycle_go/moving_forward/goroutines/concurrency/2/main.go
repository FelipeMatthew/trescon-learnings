package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Waiting groups
var waitingGroup sync.WaitGroup

// Vai forçar a sempre trablhar com numero máximo de CPUS
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// fmt.Println(runtime.NumCPU()) // Saber quantas cpus seu dispositivo vai user
	// Quantos processos vão rodar:
	waitingGroup.Add(2) // Rodar dois processos

	go runProcess("P1", 10)
	go runProcess("P2", 10)

	waitingGroup.Wait() // Espera eles rodarem
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitingGroup.Done() // Assim que sair do loop ele vai finalizar o processo
}

// Paralelismo
// vai ser um core dedicado para cada threads, eles vao rodar exatamente ao mesmo tempo
// P1 ____________ (1 core)
// P2 ____________ (2 core)

// Concorrencia
// Vai rodar diferentes funções mas nunca ao mesmo tempo.
// P1 _ _ _ __ _ _ __ _ _ _
// P2  _ _ __  _ _ _ __ ___
