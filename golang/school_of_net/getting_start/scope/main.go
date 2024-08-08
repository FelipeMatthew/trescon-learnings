package scope

import "fmt"

var y int = 10

func main() {
	x := 10
	fmt.Println(x)
	showY()
	showZ()
}

func showY() {
	fmt.Println(y)
}

func showZ() {
	fmt.Println(z)
}

// Go run *.go
