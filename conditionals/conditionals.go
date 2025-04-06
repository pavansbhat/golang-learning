package conditionals

import "fmt"

func Conditionals() string {
	if 7%2 == 0 {
		fmt.Println(": if condition")
	} else {
		fmt.Println(": else condition!")
	}

	if 8%4 == 0 {
		fmt.Println(": if condition without an else condition!")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println(": Example of || condition, Same way you can have && also!")
	}

	if num := 9; num < 0 {
		fmt.Println(": num can be declared in the same line after if, and it will be available in that same if, else if or else branch!")
	} else if num < 10 {
		fmt.Println(": inside the else if block")
	} else {
		fmt.Println(": Inside else block!")
	}

	return "Conditionals chapter!"
}
