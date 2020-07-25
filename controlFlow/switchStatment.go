package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case 2 == 4:
		fmt.Println("this should not print")
	case 3 == 3:
		fmt.Println("this should print")
		fallthrough
	case true:
		fmt.Println("this would not print if we didn't have fallthrough declared")
		fallthrough
	case false:
		fmt.Println("this would not print if we didn't have fallthrough declared")
	default:
		fmt.Println("this print if we had no fallthroughs and no true cases")
	}

	switch "thing" {
	case "other thing":
		fmt.Println("this should not print")
	case "more things":
		fmt.Println("this should not print")
	default:
		fmt.Println("default print if we had no case eval to true")
	}

	car := "civic"
	switch car {
	case "Honda", "civic":
		fmt.Println("you have a honda")
	case "bmw", "m3":
		fmt.Println("you have a bmw")
	default:
		fmt.Println("you do not have a car")
	}

}
