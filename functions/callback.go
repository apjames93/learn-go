package main

import "fmt"

// https://golang.org/ref/spec#Defer_statements
func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)
	fmt.Println(s)

	// pass in sum to be called at the end of the even block
	callback := even(sum, ii...)
	fmt.Println(callback)

	callback2 := odd(sum, ii...)
	fmt.Println(callback2)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	// invokes callback passed in in this example that would be sum
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
