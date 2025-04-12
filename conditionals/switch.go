package conditionals

import (
	"fmt"
	"time"
)

func SwitchChapter() string {
	i := 2
	switch i {
	case 1:
		fmt.Println("Case 1 - basic switchExpression")
	case 2:
		fmt.Println("Case 2- basic switchExpression")
	case 3:
		fmt.Println("Case 3 - basic switchExpression")
	}

	// Comma separated multi expression in switch statements

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend - comma seperated case expression example")
	default:
		fmt.Println("Weekday - default case example")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Morning time - switch without expression example i.e if/else alternative")
	default:
		fmt.Println("post noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Int")
		default:
			fmt.Println("Dont know the type! %T\n", t)
		}
	}

	whatAmI(false)
	whatAmI(1)
	whatAmI("hey")

	return "Switch statements"
}
