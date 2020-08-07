package main

import "fmt"

func main() {
	bd := 1993
	for bd <= 2020 {
		if bd > 2020 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
