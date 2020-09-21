package main

import "fmt"

func main() {
	g := func(n int) int {
		return n * n
	}

	thing := foo(g, 4)
	fmt.Println(thing)
}

func foo(f func(n int) int, num int) int {
	n := f(num)
	n++
	return n
}
