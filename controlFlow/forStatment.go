package main

import "fmt"

// https://golang.org/doc/effective_go.html#for
func main() {
	// for statment with a single condition
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")

	// for with break
	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}

	// for statement with for clause
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}
}
