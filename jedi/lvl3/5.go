package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("when %v is divided by 4 the remainder (or modulus) %v\n", i, i%4)
	}
}
