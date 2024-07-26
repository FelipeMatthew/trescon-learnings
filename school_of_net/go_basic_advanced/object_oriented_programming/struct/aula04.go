package main

import (
	"encoding/json"
	"fmt"
)

// Exported
type Person struct {
	Firstname  string `json:"-"` // * Não vai incluir ele no json
	Lastname   string `json:"lastname"`
	Age        uint
	Profession string
	Salary     float32
}

func testing() {
	person1 := Person{
		Firstname:  "Felipe",
		Lastname:   "Matthew",
		Age:        20,
		Profession: "developer",
		Salary:     1500.50,
	}

	// Transformar em json
	result, _ := json.Marshal(person1)

	// fmt.Println(result) // Tabela asc
	fmt.Println(string(result))

	// Exported - Public -  Maiusculo
	// Unexported - private - Minusculo
}
