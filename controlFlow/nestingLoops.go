package main

import "fmt"

func main() {
	// for init: condition: post
	for i := 0; i <= 10; i++ {
		fmt.Printf("outter loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t inner loop %d\n", j)
		}
	}
}
