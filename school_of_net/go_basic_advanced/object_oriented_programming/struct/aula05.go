package main

import (
	"encoding/json"
	"fmt"
)

// Exported
type Doctor struct {
	Firstname  string `json:"-"` // * Não vai incluir ele no json
	Lastname   string `json:"lastname"`
	Age        uint
	Profession string
	Salary     float32
}

func main() {
	// person1 := Doctor{
	// 	Firstname:  "Felipe",
	// 	Lastname:   "Matthew",
	// 	Age:        20,
	// 	Profession: "neuro cirurgião",
	// 	Salary:     13500.50,
	// }

	var person Doctor

	// Binding
	data := []byte(`{
		"firstname":"Felipe",
		"Lastname":"Matthew",
		"Age":20,
		"Profession":"neuro cirurgião",
		"Salary":13500.50,
	}`)

	json.Unmarshal(data, &person)
	fmt.Println(person.Firstname)
}
