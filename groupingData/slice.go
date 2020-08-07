package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	// x := type{valuse} // composit literal https://golang.org/doc/effective_go.html#composite_literals
	x := []int{6, 9, 6, 9, 6, 9}
	fmt.Println(x)
}

// a SLICE allows you to group together values of the same type
