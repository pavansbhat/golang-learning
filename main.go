package main

import (
	"fmt"

	"example.com/golang-learning/conditionals"
	"example.com/golang-learning/intro"
	loop "example.com/golang-learning/loops"
	"example.com/golang-learning/variables"
	"example.com/golang-learning/variables/constants"
)

func main() {
	hello := intro.HelloWorld()
	aVal, bVal, cVal, yVal, dVal := variables.VariablesFunc()
	n, d := constants.Constantsfunc()
	loops := loop.Looper()
	conditional := conditionals.Conditionals()

	fmt.Println(hello, " : Intro")
	fmt.Println("Returned from variables : ", aVal, bVal, cVal, yVal, dVal)
	fmt.Println(n, d, ": Constants chapter")
	fmt.Println(loops, ": For loop")
	fmt.Println(conditional, ": Conditionals chapter!")
}
