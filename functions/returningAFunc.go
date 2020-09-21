package main

import "fmt"

// https://golang.org/ref/spec#Function_types
func main() {
	x := bar()
	fmt.Printf("%T\n", x) // returns type of the func that is returned
	fmt.Println(x())      // execute the function to return the int value of 69

	fmt.Println(bar()()) // execute bar and func to return int value
}

func bar() func() int {
	return func() int {
		return 69
	}
}
