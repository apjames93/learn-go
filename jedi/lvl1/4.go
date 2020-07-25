package main

import "fmt"

type myNewType int

var x myNewType

func main() {
	fmt.Println(x)        // zero value of int 0
	fmt.Printf("%T\n", x) // main.myNewType

	x = 10
	fmt.Println(x)
}
