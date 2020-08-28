package main

import "fmt"

func main() {
	x := []int{42, 43, 45, 46, 47, 48, 48, 49, 50, 51}
	fmt.Println(x[:5]) // starts at the 0 and ends at 5
	fmt.Println(x[5:]) // starts at 5 and ends at len of array
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
