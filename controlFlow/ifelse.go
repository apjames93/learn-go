package main

import "fmt"

func main() {
	x := 39

	for i := 0; i < 50; i++ {
		x++
		if x == 40 {
			fmt.Println("value is 40")
		} else if x == 41 {
			fmt.Println("value is 41")
		} else if x == 42 {
			fmt.Println("value is 41")
		} else if x == 43 {
			fmt.Println("value is 41")
		} else if x == 44 {
			fmt.Println("value is 41")
		} else {
			fmt.Println("value is not 40")
		}
	}

}
