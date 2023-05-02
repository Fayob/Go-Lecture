package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook Chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func PrepareDishes(dishes []Preparer)  {
	fmt.Println("Prepare Dishes")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("--Dish: %v--\n", dish)
		dish.PrepareDish()
	}
	fmt.Println()
}


func main()  {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Salad taste")}
	PrepareDishes(dishes)

}