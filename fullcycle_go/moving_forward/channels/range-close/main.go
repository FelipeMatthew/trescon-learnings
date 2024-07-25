package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	// Paralelo
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	// Paralelo
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		// Vai esperar os dois acima rodarem para finalizar
		waitGroup.Wait()
		close(channel) // Assim que tudo rodar, vai fechar os canais
	}()

	// Enquanto o cannal estiver rodando o range vai mantar o loop
	// Range esvazia automaticamente o channel para que o prox valor seja atribuido
	for number := range channel {
		fmt.Println(number)
	}

	time.Sleep(time.Second)
}
