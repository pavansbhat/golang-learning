package loop

import "fmt"

func Looper() int {
	i := 1

	for i <= 3 {
		fmt.Println(i, ": for loop with a single condition")
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j, ": Standard for loop")
	}

	for i := range 3 {
		fmt.Println(i, ": range loop")
	}

	for {
		fmt.Println(": for loop like the while loop, but not really a while loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n, ": for loop with continue example")
	}

	return i
}
