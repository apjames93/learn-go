package main

import "fmt"

func main() {
	fun := func() {
		fmt.Println("anonymous func")
	}

	fun()
}
