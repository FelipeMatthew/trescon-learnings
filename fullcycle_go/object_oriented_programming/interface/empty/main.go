package main

import "fmt"

type Name []interface{}

func (n *Name) load() {
	*n = Name{
		"felipe",
		"josé",
		"maria",
		"luna",
	}
}

func (n Name) print() {
	fmt.Println(n)
}

// Interface vazia aceita tudo
func main() {
	var names Name

	names.load()
	names.print()
}
