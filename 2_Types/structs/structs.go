package main

import "fmt"

// Demo
type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

// Exercise
type Coordinate struct {
	x, y int
}

type Rectangle struct {
	a, b Coordinate
}

func width(rect Rectangle) int  {
	return (rect.b.x - rect.a.x)
}
func length(rect Rectangle) int  {
	return (rect.a.y - rect.b.y)
}
func area(rect Rectangle) int  {
	return length(rect) * width(rect)
}
func perimeter(rect Rectangle) int  {
	return (length(rect) * 2) + (width(rect) * 2)
}
func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimeter is", perimeter(rect))
}

// type Rectangle struct {
// 	field string
// 	length, breadth int
// }

// func area(length, breadth int) int  {
// 	return length * breadth
// }

// func perimeter(length, breadth int) int {
// 	return (2 * length) + (2 * breadth)
// }

func main() {
	// Demo
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 2
	fmt.Println(heidi)

	var (
		bill = Passenger{ Name: "Bill", TicketNumber: 3}
		ella = Passenger{ Name: "Ella", TicketNumber: 4}
	)
	fmt.Println(bill, ella)

	casey.Boarded = true
	bill.Boarded = true

	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	if bill.Boarded {
		fmt.Println(bill.Name, "has boarded the bus")
	} else {
		fmt.Println(bill.Name, "has not boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{ heidi }
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

	// Exercise
	rect := Rectangle{ a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect)

	rect.a.y *= 2
	rect.b.x *= 2

	printInfo(rect)
	// areaRec := Rectangle{ field: "Area", length: 8, breadth: 4}
	// perimeterRec := Rectangle{ field: "Perimeter", length: 14, breadth: 12}

	// areaResult := area(areaRec.length, areaRec.breadth)
	// fmt.Println("Area is", areaResult)
	// perimeterResult := perimeter(perimeterRec.length, perimeterRec.breadth)
	// fmt.Println("Perimeter", perimeterResult)
}
