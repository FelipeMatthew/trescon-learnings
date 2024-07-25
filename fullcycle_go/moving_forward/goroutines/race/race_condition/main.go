package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func main() {
	go runProcess("P1", 10)
	go runProcess("P2", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result) // 10
}

func runProcess(name string, total int) {
	// O result ele esta sendo chamado e aplicado em z
	// Mas quando faz o outro loop ele pergunta qual o valor do result e coloca esse valor novamente repetindo
	// Ao inves de ser 20x o result ele so vai sair 10 por conta que fica atribuindo mais de uma vez antes de atribuir novamente
	// Por isso nome de corrida, um chega mais rapido e pega o valor
	for i := 0; i < total; i++ {
		z := result
		z++

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)

		result = z
		fmt.Println(name, "->", i, "Partial result: ", result)
	}
}
