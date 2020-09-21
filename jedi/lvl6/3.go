package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("bar")
}

func foo() {
	defer func() {
		fmt.Println("double defer!")
	}()
	fmt.Println("foo")
}
