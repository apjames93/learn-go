package main

import (
	"fmt"
	"os"
)

// https://golang.org/ref/spec#Defer_statements
func main() {
	defer foo() // will wait to run this until program is exiting
	// good practices for when you open a file to defer and close the file in the beginning of the func
	bar()

	// if working with a func that can return an error defer after checkiong for the error
	s, err := thingWithError()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer s()

}

func thingWithError() (func(), error) {
	return foo, nil
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
