package main

import "fmt"

func printSlice(title string, slice []string)  {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

// Exercise
type Part string

func showLine(line []Part)  {
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Println(part)
	}
}

func main()  {
	route := []string{"Grocery", "Department Store", "Salon"}
	printSlice("Route 1", route)

	route = append(route, "Home")
	printSlice("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "Visited")
	fmt.Println(route[1], "Visited")

	route = route[2:]
	printSlice("Remaining locations", route)

	// Exercise
	assemblyLine := make([]Part, 3)

	assemblyLine[0] = "Pipe"
	assemblyLine[1] = "Bolt"
	assemblyLine[2] = "Sheet"

	fmt.Println("3 parts:")
	showLine(assemblyLine)

	assemblyLine = append(assemblyLine, "Washer", "Wheel")
	fmt.Println("\nAdded two parts:")
	showLine(assemblyLine)

	assemblyLine = assemblyLine[3:]
	fmt.Print("\nSliced:")
	showLine(assemblyLine)
}