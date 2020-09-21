package main

import "fmt"

func main() {
	c := clouser()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

func clouser() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
