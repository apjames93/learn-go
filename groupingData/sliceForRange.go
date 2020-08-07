package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	x := []int{6, 9, 6, 9, 6, 9}

	// print length of slice
	fmt.Println(len(x))

	fmt.Println(x)
	// get value by index position
	fmt.Println(x[3])

	for index, value := range x {
		fmt.Println(index, value)
	}
}
