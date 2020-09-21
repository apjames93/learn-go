package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(ii...))
	fmt.Println(bar(ii))
}

func foo(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0

	for _, v := range xi {
		total += v
	}
	return total
}
