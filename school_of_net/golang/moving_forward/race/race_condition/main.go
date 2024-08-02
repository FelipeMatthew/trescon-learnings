package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go run -race main.go
// Vai mostrar se o sistema tem alguma race condition

var result int

func main() {
	go runProcess("P1", 10)
	go runProcess("P2", 10)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result) // 10 - O certo seria 20
}

func runProcess(name string, total int) {
	// O result ele esta sendo chamado e aplicado em z
	// Mas quando faz o outro loop ele pergunta qual o valor do result e coloca esse valor novamente repetindo
	// Ao inves de ser 20x o result ele so vai sair 10 por conta que fica atribuindo mais de uma vez antes de atribuir novamente
	// Por isso nome de corrida, um chega mais rapido e pega o valor

	// Interfere no bom funcionamento da aplicação - Race condition
	// Seus resultados podem estar sendo comprometos

	// Processos rodando um encima do outro encavalando seu sistema.
	for i := 0; i < total; i++ {
		z := result
		z++

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)

		result = z
		fmt.Println(name, "->", i, "Partial result: ", result)
	}
}
