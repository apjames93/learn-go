package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	x := []int{0, 1, 2, 3, 4}
	fmt.Println(x)

	firstTwoValues := x[:2]
	lastTwoValues := x[3:]

	// delete value 2 at postion 2 from x slice of ints
	x = append(firstTwoValues, lastTwoValues...)
	fmt.Println(x)
}
