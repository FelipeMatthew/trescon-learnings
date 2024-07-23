package variables

import "fmt"

const value1 int = 233

const (
	A = ""
	b = ""
	c = 0
	d = ""
)

func readConsts() {
	fmt.Println(value1, b, c, d)
}
