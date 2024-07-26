package main

import "fmt"

// Fun out inverso de fun in
// Ao inves de fazer um funil ele vai espalhar go routines
func main() {
	chann := generate(2, 10, 40)

	// Disttribuindo cannal
	d1 := divide(chann)
	d2 := divide(chann)

	fmt.Println(<-d1, <-d2)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()

	return channel
}
