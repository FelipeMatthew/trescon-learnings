package main

import "fmt"

func aula01() {
	// * Array mais dinamico
	// Faz um apontamento internamente para um array

	// Make - cria o slice na memoria
	// Tipo de array, quantos indices
	slice := make([]int, 2, 5) // len - 2 / cap - 5

	// _ := [4]int{2, 5, 3, 5}
	// Adicona mais um indice no array
	// * slice = append(slice, 3, 53, 43)

	// Nao consegue adicionar valor para ele na mao
	// slice[2] = 10 - Error

	// Ele adicona mas quando ultrapassa a capacidade do slice ele dobra de tamanho
	slice = append(slice, 3, 5, 7, 8, 9) // Cap - 10

	fmt.Println(slice)
	fmt.Println(len(slice)) // 7
	// Capacidade
	fmt.Println(cap(slice)) // 10
}
