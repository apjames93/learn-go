package main

import "fmt"

// https://golang.org/doc/effective_go.html#for
func main() {
	// for init: condition: post
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
