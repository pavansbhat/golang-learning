package variables

import (
	"fmt"
)

const Version = "1.0.0"

func VariablesFunc() (string, int, int, int, string) {
	a := "initial"
	fmt.Println(a, "var a")

	var b, c int = 1, 2
	fmt.Println(b, c, "var b, c")

	x := true
	fmt.Println(x, "var x")

	var y int
	fmt.Println(y, "var y")

	d := "Hello"
	fmt.Println(d, "var d")

	return a, b, c, y, d
}
