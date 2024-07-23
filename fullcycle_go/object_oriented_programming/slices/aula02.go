package main

import "fmt"

func aula02() {
	slice := make([]int, 2, 5)
	// slice = append(slice, 10, 40, 50, 30)

	// fmt.Println(slice)      // [0, 0, 10, 40, 50, 30]
	// fmt.Println(len(slice)) // 6
	// fmt.Println(cap(slice)) // 10s

	//  * for i := 0; i < 20; i++ {
	// 	* slice = append(slice, i)
	// 	* fmt.Println("Len: ", len(slice), " Cap: ", cap(slice))
	//  * }

	sliceTest := slice
	// * Quando a capacidade do array eh estourada o array copia, vai pegar outro endrecamento na memoria ao qual ele so vai ser uma mera copia do array original, nao vai ter mais relacao igual tinha anteriormente.

	slice = append(slice, 30, 50, 70, 3) // Excedeu a capacidade

	slice[0] = 10
	// Ele nao faz uma copia, ele pega o valor da memoria
	// Se alterar um altera os dois

	// Mesmo ponteiro - apontamento ainda eh para o mesmo array
	fmt.Println(slice)     // [10, 0, ...]
	fmt.Println(sliceTest) // [0, 0]
}
