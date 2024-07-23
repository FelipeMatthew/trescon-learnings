package main

import "fmt"

func main() {
	// array - [10]string{}
	// slice - []string {}

	sliceString := []string{
		"hello",
		"world",
		"from",
		"felipe",
	}

	fmt.Println(sliceString[0]) // Hello
	// Pega do 2 indices - 0, 1
	fmt.Println(sliceString[:2]) // [hello, world]

	// Pega do indice 1 ate o 3 = 1, 2
	fmt.Println(sliceString[1:3]) // [world, from]

	// Pega do segundo indice e vai ate o final - 2, 3 ...
	fmt.Println(sliceString[2:]) // [from, felipe]

	fmt.Println(len(sliceString))
	fmt.Println(cap(sliceString))
}
