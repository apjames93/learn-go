package main

import "fmt"

// https://golang.org/doc/effective_go.html#arrays
func main() {
	// arrays are zero base index
	// the length is part of its type
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	fmt.Println(len(x))
}
