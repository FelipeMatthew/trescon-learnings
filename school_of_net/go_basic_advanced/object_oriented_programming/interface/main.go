package main

import "fmt"

// Interface vai forçar que todos tenham mesma assinatura
type Animal interface {
	sound() string
}

type dog struct {
	name string
}

type cat struct {
	name string
}

// Method do objeto
// Vai definir que o objeto faz parte por conta do método
// não existe em go - implements - ela ja é atribuida automaticamente
func (d dog) sound() string {
	return "this sound came from: " + d.name
}

func animalSound(a Animal) string {
	return a.sound()
}

func main() {
	dog := dog{"dog"}

	fmt.Println(dog.sound())
}
