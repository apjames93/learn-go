package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	x := []int{1, 2, 3, 4, 5}

	// get value by index
	fmt.Println(x[1])
	fmt.Println(x)

	// print value starting at position 1
	fmt.Println(x[1:])

	// print value starting at position 1 and ending at position 3
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}
