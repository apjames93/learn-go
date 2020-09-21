package main

import "fmt"

// https://golang.org/doc/effective_go.html#allocation_make
func main() {
	x := make([]int, 10, 100)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[9] = 999

	// will error for index out of range
	// x[10] = 69

	// append will make slice go to 11
	x = append(x, 3030)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := make([]int, 11, 12)

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	y = append(y, 10)
	fmt.Printf("the len is %v and the cap is %v\n", len(y), cap(y))

	// append will make a new array once the cap is passed and will double the size of the underlying array
	y = append(y, 10)
	fmt.Printf("the len is %v and the cap is %v\n", len(y), cap(y))

}

// if you know the size of the slice declare it so you don't need to create the array in the run time
