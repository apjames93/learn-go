package main

import "fmt"

// https://golang.org/ref/spec#Function_types
func main() {
	fmt.Println("foo")

	f := func() {
		fmt.Println("func expression")
	}

	f()
}
