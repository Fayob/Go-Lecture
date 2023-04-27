package main

import "fmt"

type Room struct {
	name string
	cleaned bool
}

func checkCleanliness(rooms [4]Room)  {
	for i := 0; i < len(rooms); i++ {
		// NOTE: Always make a copy of iteration (i), don't use it directly
		// Instead of using i directly you can create another variable like this j := i and make use of j 
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}

// Exercise
type Products struct {
	name string
	price int
}

func output(products []Products)  {
	fmt.Println("last item is", products[len(products) - 1])
	fmt.Println("Total number of item is", len(products))
	var sum int
	for i := 0; i < len(products); i++ {
		product := products[i].price
		sum += product
	}
	fmt.Println("Total cost is", sum)
}

func main()  {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	rooms[1].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)

	// Exercise
	products := [4]Products{
		{ name: "Jean", price: 12 },
		{ name: "Show", price: 10 },
		{ name: "Shirt", price: 14 },
	}
	output(products[:])
	
	products[3] = Products{ name: "Trouser", price: 20 }
	
	output(products[:])
}