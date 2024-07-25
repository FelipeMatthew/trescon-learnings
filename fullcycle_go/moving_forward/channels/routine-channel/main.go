package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()

	time.Sleep(time.Second)

	// Revesamento - Exemplo de natação - So vai sair proximo nadador quando um voltar e sair
	// Channel so pode receber um outro valor quando esse channerl for esvaziado
	// duas funçoes trabalhando em paralelo
	// A de cima atribui valor, a de baixo solta o valor
}
