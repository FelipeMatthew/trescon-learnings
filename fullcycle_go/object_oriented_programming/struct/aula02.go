package main

import "fmt"

// Go trabalha orientado a objeto com struct
// * Herança
// * Overwriting
type Moto struct {
	Name  string
	Year  int
	Color string
}

type SuperMoto struct {
	Moto          // *
	Name   string // * overwriting
	CanFly bool
}

// type SuperMoto struct {
// 	Name   string
// 	Year   int
// 	Color  string
// 	CanFly bool
// }

// Atachando a funcao a struct
func (m Moto) information() string {
	return fmt.Sprintf("name: %s\nyear: %d\ncolor: %s\n", m.Name, m.Year, m.Color)
}

func motoca() {
	moto := Moto{
		Name:  "motinha",
		Year:  20,
		Color: "azul",
	}

	sMoto := SuperMoto{
		Moto: Moto{
			Name:  "motinha",
			Year:  20,
			Color: "azul",
		},
		CanFly: true,
		Name:   "Car overwrited",
	}

	fmt.Println(moto)
	fmt.Println(sMoto)
	fmt.Println(sMoto.information())
}
