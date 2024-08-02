package main

import "fmt"

func main() {
	// flags - Semaforo - Vai rodar apenas no sinal verde

	channel := make(chan int)
	ok := make(chan bool) // * flag

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		ok <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		ok <- true
	}()

	// Precisa de algo para esvaziar pois está cheio
	go func() {
		<-ok // Esvaziando como semáforo
		<-ok // Sinal verde, uma flag
		close(channel)
	}()

	// Show channel
	for number := range channel {
		fmt.Println(number)
	}
}
