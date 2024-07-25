package main

import "fmt"

func main() {
	// Criar um função para retornar um pipeline
	// Quando uma coisa consome a outra
	numbers := generate(2, 4, 6)
	result := divide(numbers)

	fmt.Println(<-result) // Valor do generate 2
	fmt.Println(<-result) // Valor do generate 4
	fmt.Println(<-result) // Valor do generate 6
}

// Cria uma função que vai retornar um channel que vai receber valores de numeros
// Vai gerar o channel
func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
	}()

	return channel
}
