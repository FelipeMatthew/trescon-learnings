package main

import "fmt"

func main() {
	a := 10

	if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("a = 10")
	}

	b := true

	if !b { // false
		fmt.Println(false)
	}

	// Atribui o valor de x e se o b for verdadeiro vai imprimir o valor de x
	if x := "hii"; b {
		fmt.Println(x)
	}
}
