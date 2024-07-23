package main

import "fmt"

// Go trabalha orientado a objeto com struct
type Car struct {
	Name  string
	Year  int
	Color string
}

// Atachando a funcao a struct
func (c Car) info() string {
	return fmt.Sprintf("name: %s\nyear: %d\ncolor: %s\n", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{
		Name:  "carinho",
		Year:  20,
		Color: "azul",
	}

	fmt.Println(car1.info())
}
