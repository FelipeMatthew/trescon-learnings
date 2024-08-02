package main

import (
	"encoding/json"
	"fmt"
)

// Unexported
type person struct {
	firstname  string
	lastname   string
	age        uint
	profession string
	salary     float32
}

func persona() {
	person1 := person{
		firstname:  "Felipe",
		lastname:   "Matthew",
		age:        20,
		profession: "developer",
		salary:     1500.50,
	}

	// Transformar em json
	result, _ := json.Marshal(person1)

	// fmt.Println(result) // Tabela asc
	fmt.Println(string(result))

	// Exported - Public -  Maiusculo
	// Unexported - private - Minusculo
}
