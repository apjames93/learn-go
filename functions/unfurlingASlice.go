package main

import "fmt"

// https://golang.org/ref/spec#Function_declarations
func main() {
	xi := []int{1, 2, 3, 4, 5}
	// unfurlSlice(xi) passes values as slice ([1, 2, 3, 4, 5])
	unfurlSlice(xi...) // pass individual slice values (1, 2, 3, 4, 5) instead of
}

func unfurlSlice(x ...int) {
	// x is a slice of values [1, 2, 3, 4, 5]
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
