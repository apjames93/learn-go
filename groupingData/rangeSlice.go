package main

import "fmt"

// https://golang.org/doc/effective_go.html#slices
func main() {
	xi := []int{6, 9, 6, 9, 6, 9}
	fmt.Println(xi)

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
