package main

import "fmt"

func main() {
	a := "felipe"

	switch a {
	case "felipe":
		fmt.Println("hi felipe")
	case "jose":
		fmt.Println("hi jose")
	case "arlindo":
		fmt.Println("hi arlindo")
	case "felps":
		fmt.Println("hi felps")
	default:
		fmt.Println("User not founded")
	}
}
