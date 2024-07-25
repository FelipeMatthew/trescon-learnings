package main

import "fmt"

func main() {
	// Criando um canal
	// Make cria slice, map, channel
	msg := make(chan string)

	// Processo vai rodar em paralelo com o sistema
	go func() {
		// Atribuir valor em um channel
		msg <- "Hello world"
	}()

	result := <-msg

	fmt.Println(result)
}
