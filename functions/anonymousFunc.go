package main

import "fmt"

// https://golang.org/ref/spec#Defer_statements
func main() {
	foo()
	// anonymous fun runs
	func() {
		fmt.Println("anonymous func ran")
	}() // will invoke block

	func(x int) {
		fmt.Println("anonymous func with arguments", x)
	}(69)
}

func foo() {
	fmt.Println("foo ran")
}
