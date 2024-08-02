package main

import "fmt"

func main() {
	// Deadlock - channel tem que rodar dentro de uma go routine

	channel := make(chan int)

	go func() {
		channel <- 10
	}()

	// channel <- 10 - Gera um deadlock

	fmt.Println(<-channel)
}
