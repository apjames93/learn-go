package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)

	b, s := bar()
	fmt.Println(b, s)
}

func foo() int {
	return 69
}

func bar() (int, string) {
	return 999, "cool"
}
