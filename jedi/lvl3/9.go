package main

import "fmt"

func main() {
	favCar := "honda"
	switch favCar {
	case "subaru":
		fmt.Println("good subaru")
	case "toyota":
		fmt.Println("good toyota")
	case "honda":
		fmt.Println("good honda")
	case "ford":
		fmt.Println("bad")
	}
}
