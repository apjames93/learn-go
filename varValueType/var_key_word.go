package main

import "fmt"

// global scope for package
// declare and assign = initilization
var y = 69 //  not best practices should keep scope as narrow as possibal

// main ..
func main() {
	fmt.Println(y)

	var x = 43
	fmt.Println(x)

	// Declare there is a variable with the identifier "z"
	// and that the variable wit the identifier "z" of type int
	// Assigns the zero value of type int to "z"
	// https://golang.org/ref/spec#The_zero_value
	var z int
	fmt.Println(z)

}
