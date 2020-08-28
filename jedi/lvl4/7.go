package main

import "fmt"

func main() {
	x := []string{"one", "two", "three"}
	x2 := []string{"four", "five", "six"}
	fmt.Println(x)
	fmt.Println(x2)

	xxs := [][]string{x, x2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("\t record", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: %v", j, val)
		}
	}
}
