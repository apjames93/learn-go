package main

import "fmt"

// main operator https://golang.org/ref/spec#Types
func main() {
	// DECLARE the VARIABLE with the IDENTIFIER "a" is of TYPE INT and ASSIGN 42
	var a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// DECLARE the VARIABLE with the IDENTIFIER "b" is of TYPE STRING and ASSIGN "cool dude"
	var b = "cool dude"
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = "this would break because it is the wrong TYPE VARIABLE a is unable to be of TYPE STRING"

	// string literal https://golang.org/ref/spec#String_literals
	var c = `some new info "wow so new"`
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
