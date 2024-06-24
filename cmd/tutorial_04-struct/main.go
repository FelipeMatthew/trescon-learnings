package main

import (
	"fmt"
)

type marriedWith struct {
	name string
	lastName string
}

type person struct {
	name     string
	lastName string
	age      int8
	marriedWith
	int
}

func main() {
	myPerson := person{"felipe", "matthew", 30, marriedWith{"tatiane", "lima"}, 3}

	// Another method to create a type
	car := struct{
		name string
		year int8
	}{"felipe", 30}

	fmt.Println(car)

	myPerson.name, myPerson.lastName = "tatiane", "lima"

	fmt.Println(myPerson)

	var myConcatNames myInterface
	fmt.Printf("The name concatened is: %v", myConcatNames)


}

type myInterface interface {
	concactNames()
}

func (e person) concactNames() string {
	concact := e.name + " " + e.lastName
	return concact
}