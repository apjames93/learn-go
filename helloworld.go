package main

import "fmt"

func main() {
	fmt.Println("hello world")
	foo()
	loop()
	bar()
}

func foo() {
	fmt.Println("in foo you fool")
}

func loop() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func bar() {
	fmt.Println("It has been done!")
}
