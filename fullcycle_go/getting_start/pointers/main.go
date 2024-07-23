package main

import "fmt"

func xpto(a *int) int {
	*a = 100
	fmt.Println(a)
	return *a
}

func main() {
	// Apontamento e enderecamento de memoria
	// abre uma caixa e guarda o x dentro dessa caixa
	// & - aponta para memoria
	// _ * pega o valor
	x := 10

	// &x = 0xc00000a098 - Local na memoria
	//  x = 10
	fmt.Println(&x)

	y := &x
	fmt.Println(y) // 0xc00000a098

	// Dereferencing - *
	fmt.Println(*y) // 10

	// Ponteiro da memoria, que esta o x
	*y = 20
	fmt.Println(x) // 20

	// Vai guardar definidamente o valor de z tem que ser um int e nao pode ser mudado
	var z *int = &x
	fmt.Println(z)  // 0xc00000a098
	fmt.Println(*z) // 0xc00000a098

	// XPTO Func
	b := 10

	// B pegou endereco da memoria de a que era 100
	xpto(&b)       // 0xc00000e80
	fmt.Println(b) // 100

}
