package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main()  {
	dice, sides := 2, 6
	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0

		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", Die #", d, ":", rolled)
		}
		fmt.Println("Total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("Snake Eyes")
		case sum == 7:
			fmt.Println("Lucky 7")
		case sum % 2 == 0:
			fmt.Println("Even")
		case sum % 2 == 1:
			fmt.Println("Odd")
		}
	}
	fmt.Print(1, 2, 3)
	fmt.Print(true, false)
	fmt.Print("a", "b", "c")
}
