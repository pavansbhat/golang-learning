package constants

import (
	"fmt"
	"math"
)

const Version = "1.0.0"

func Constantsfunc() (int, int) {
	const n = 50000000

	const d = 3e25 / n
	fmt.Println(d, "d")

	fmt.Println(int64(d), "converted d")

	fmt.Println(math.Sin(n))

	return n, d
}
