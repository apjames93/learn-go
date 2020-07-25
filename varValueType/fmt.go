package main

import "fmt"

// https://godoc.org/fmt
func main() {
	var y = 69
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	fmt.Printf("%T\n%b\n%x", y, y, y)
}
