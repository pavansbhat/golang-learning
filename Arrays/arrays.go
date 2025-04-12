package arrays

import "fmt"

func ArraysInGo() string {
	var a [5]int
	fmt.Println(a, ": Zero valued array of size 5")

	a[4] = 100
	fmt.Println(a, ": Array value is being set by accessing the index!")

	fmt.Println(len(a), ": Length of the a arrays can be accessed using builtin function len()")

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b, ": declare and Initialize [size]type{items comma separated!}")

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b, ": Have the compiler count the number of elements inside an array using `...`")

	b = [...]int{100, 3: 100, 200}
	fmt.Println(b, ": by using index and : we can have the elements in between be zeoroed")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD, ": Arrays are one dimensional, but you can create a two dimensional ones too")

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(twoD, ": You can declare and initialize the two deminsional array!")

	return "Arrays chapter!"
}
