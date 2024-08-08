package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	// while
	x := 0
	for x == 10 {
		fmt.Println(x)
		x++
	}

	// Infinity loop

	for { // true
		x++

		if x == 50 {
			continue
		}

		if x == 100 {
			break
		}
	}
}
