package main

import (
	"fmt"
	"unicode/utf8"
)

func foo(){
	fmt.Println("hi im foo")
	return
}

func main() {
	// fmt.Println("Hello world")

	var intNum int = 32767 // The limit of int
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = `hello 
	world`

	var anotherString string = "concat" + " " + "strings"
	fmt.Println(myString, anotherString)

	var lengthOf string = "testing"
	fmt.Println(utf8.RuneCountInString(lengthOf))

	var boolean bool = true
	fmt.Println(boolean)

	// Another ways to declare
	// Default value of numbers is 0 and bool is false


	var strangerFunction string = foo()
	fmt.Println(strangerFunction)

	var1, var2 := 1, 3
	fmt.Println(var1, var2)

	const declaringConst string = "i cant change my value"

	
	
}