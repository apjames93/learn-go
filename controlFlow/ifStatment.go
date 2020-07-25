package main

import "fmt"

func main() {
	if true {
		fmt.Println("001") // will run
	}

	if false {
		fmt.Println("002") // wont run
	}

	if !true {
		fmt.Println("003") // wont run
	}

	if !false {
		fmt.Println("004") // will run
	}

	if 2 == 2 {
		fmt.Println("005") // will run
	}

	if 2 != 2 {
		fmt.Println("006") // wont run
	}

	// initialization statement in if statement
	// scope of x is limited to if block
	if x := 42; x == 2 {
		fmt.Println("007") // wont run
	}
}
