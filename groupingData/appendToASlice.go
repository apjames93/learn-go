package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	x := []int{0, 1, 2, 3, 4}
	fmt.Println(x)

	// you can append an unlimited amount of arguments of the same time to a slice
	x = append(x, 5, 6, 7, 8, 9)

	y := []int{10, 11, 12, 13, 14}

	// will append y int values to the x slice of int
	x = append(x, y...) // same thing append(x, y[0], y[1], y[2], y[3], y[4])
}
