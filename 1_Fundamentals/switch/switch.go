package main

import "fmt"

func price() int  {
	return 15
}

const (
	Economy = 10
	Business = 20
	FirstClass = 30
)

func main()  {
	switch p := price(); {
	case p < 10:
		fmt.Println("Cheap Item")
	case p < 20:
		fmt.Println("Moderate Item")
	default:
		fmt.Println("Expensive Item")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy Seating")
	case Business:
		fmt.Println("Business Seating")
	case FirstClass:
		fmt.Println("FirstClass Seating")
	default:
		fmt.Println("Default Seating")
	}

	// Exercise
	
	switch age := 12; {
	case age == 0:
		fmt.Println("newborn")
	case age < 4:
		fmt.Println("toddler")
	case age < 13:
		fmt.Println("child")
	case age < 18:
		fmt.Println("teenager")
	default:
		fmt.Println("adult")
	}
}