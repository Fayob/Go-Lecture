package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1
	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter){
	*old = new
	increment(counter)
}

// Exercise
const (
	Active = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag SecurityTag
}

// func activate(tag *SecurityTag)  {
// 	*tag = Active
// }
 
func deactivate(tag *SecurityTag)  {
	*tag = Inactive
}

func checkout(items []Item)  {
	fmt.Println("Checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[1].tag)
	}
}
 
func main() {
	counter := Counter{}

	hello := "Hello"
	world := "world!"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[1], "Go!", &counter)
	fmt.Println(phrase)

	// Exercise
	shirt := Item{"Shirt", Active}
	pants := Item{"Pants", Active}
	purse := Item{"Purse", Active}
	watch := Item{"Watch", Active}

	items := []Item{shirt, pants, purse, watch}
	fmt.Println("Initial", items)

	deactivate(&items[0].tag)
	fmt.Println("Item 0 deactivated", items)
	
	checkout(items)
	fmt.Println("checked out", items)
}