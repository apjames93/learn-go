package main

import "fmt"

// https://golang.org/ref/spec#Defer_statements
func main() {
	defer foo() // will wait to run this until program is exiting
	// good practices for when you open a file to defer and close the file in the beginning of the func
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
