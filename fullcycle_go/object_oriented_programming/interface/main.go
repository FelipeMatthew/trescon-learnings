package main

import "fmt"

// Interface vai forçar que todos tenham mesma assinatura
type Animal interface {
	start() string
}

type dog struct {
	name string
}

type cat struct {
	name string
}

// Method do objeto

func sound(a Animal) string {
	return a.start()
}

func main() {
	dog := dog{"dog"}

	fmt.Println(dog.sound())
}
