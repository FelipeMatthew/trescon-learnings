package main

import "fmt"

func funcName(a int) int {
	return a * a
}

// Define o retorno como um valor na variavel
func variableReturn(name string) (variable string) {
	variable = name
	return
}

// Retorna mais de dois valores
func moreReturn(a, b string) (string, string) {
	return a, b
}

// Voce pode passar um array de inteiros nesse x, pode ser um ou mais valores
func variadicFunc(x ...int) int {
	var res int

	// famoso for each
	// Index, value
	for _, value := range x {
		res += value
	}

	return res
}

// Funcao que vai retornar uma funcao que no final vai ser um valor inteiro
func funcInsideFunc() func() int {
	x := 10

	return func() int {
		return x * x
	}
}

func main() {
	x := funcName(10)
	name := variableReturn("felpola")
	firstname, lastname := moreReturn("felipe", "matthew")
	funcOutside := funcInsideFunc()
	funcInside := funcOutside()

	fmt.Println(x)
	fmt.Println(name)
	fmt.Println(firstname, lastname)
	fmt.Println(variadicFunc(1, 3, 5, 6, 7))
	fmt.Println(funcInside)

	z := 10

	// Funcao anonima
	add := func() int {
		z += 2
		return z
	}
	fmt.Println(add())
}
