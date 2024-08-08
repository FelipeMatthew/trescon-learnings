package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// x - channel
	x := funnel(generateMsg("hi from my first channel"), generateMsg("hi from my second channel"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x) // Esvaziando o x
	}
	fmt.Println("Funil limpado...")
}

// Diversos canais que vao afunilar em uma unica funcao retornando um canal
// Voce precisa de tres funcoes executadas em paralelo, faz um funilamento dos resultados, em todos os canais, e vai apontar ele para um unico canal
// Afunilamento

// So vai receber informação
func generateMsg(s string) <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - value: %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()

	return channel
}

func funnel(channel_1, channel_2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		// For infinito
		for {
			channel <- <-channel_1
		}
	}()

	go func() {
		// For infinito
		for {
			channel <- <-channel_2
		}
	}()
	// Fez o funil e tudo esta sendo salvo no channel
	return channel
}
