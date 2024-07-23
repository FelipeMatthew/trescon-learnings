package main

import "fmt"

func main() {
	// Go pode criar arrays limitados, com tamanho definido
	// Array com a capacidade de 10 elementos
	var x [10]int

	x[1] = 140
	x[2] = 110
	x[3] = 150

	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(len(x))

	z := [10]int{2, 30, 10}
	fmt.Println(z)
}
