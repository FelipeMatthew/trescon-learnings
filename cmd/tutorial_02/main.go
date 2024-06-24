package main

import (
	"errors"
	"fmt"
)
func main() {
	printValue := "Hello world"
	printMe(printValue)
	
	var numerator int = 10
	var denominator int = 5
	result, text, err := intDivisor(numerator, denominator)

	// && ||
	// If case
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if text == "" {
	// 	fmt.Printf("Invalid text")
	// } else {
	// 	fmt.Printf("The result is %v and the string text is: %v", result, text)
	// }
	
	// Switch case
	switch{
	case err != nil:
		fmt.Println(err.Error())
	case text == "":
		fmt.Printf("Invalid string input")
	default: 
		fmt.Printf("the result of integer is %v and the text is: %v", result, text)
	}
}
func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivisor(numerator, denominator int) (int, string, error) {
	var err error

	switch denominator{
		case 0:
			fmt.Printf("error") 
		case 1,2: 
			fmt.Printf("success") 
	}
	

	if denominator==0{
		err = errors.New("Cannot divide by zero")
		return 0, "", err
	}
	result := numerator/denominator
	return result, "im returning 2 values", err
}