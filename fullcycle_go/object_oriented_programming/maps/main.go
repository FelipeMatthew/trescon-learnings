package main

import "fmt"

func main() {
	// Parecido com objetos
	m := make(map[string]int) // cuida para tratar os dados do array

	m["a"] = 20
	m["altura"] = 180
	m["peso"] = 90

	delete(m, "peso") // 0 | " "

	// blankidentifier
	// value, exist - true or false
	_, exists := m["peso"]

	value, exist := m["idade"]

	fmt.Println(value, exist) // 20 true

	fmt.Println(exists) // false

	fmt.Println(m)

	// emptyMap := map[string]int{}
	// fmt.Println(emptyMap)

	newMap := map[int]int{ // Vai fazer direto
		50: 20,
		02: 20,
	}

	fmt.Println(newMap)

	if valor, existe := m["a"]; existe {
		fmt.Println(valor)
	} else {
		fmt.Print(false)
	}

}
