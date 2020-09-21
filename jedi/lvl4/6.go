package main

import "fmt"

func main() {
	x := make([]string, 50, 50)

	x = []string{
		"Alaska",
		"Alabama",
		"Arkansas",
		"American Samoa",
		"Arizona",
		"California",
		"Colorado",
		"Connecticut",
		"District of Columbia",
		"Delaware",
		"Florida",
		"Georgia",
		"Guam",
		"Hawaii",
		"Iowa",
		"Idaho",
		"Illinois",
		"Indiana",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Massachusetts",
		"Maryland",
		"Maine",
		"Michigan",
		"Minnesota",
		"Missouri",
		"Mississippi",
		"Montana",
		"North Carolina",
		" North Dakota",
		"Nebraska",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"Nevada",
		"New York",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Puerto Rico",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Virginia",
		"Virgin Islands",
		"Vermont",
		"Washington",
		"Wisconsin",
		"West Virginia",
		"Wyoming",
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
