package scope

import "fmt"

var z string = "hello from another scope"

func RunZ() {
	fmt.Println(z)
}
