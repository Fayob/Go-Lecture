package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is", sum)
	for i := 1; i < 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement", sum)
	}

	// Exercise
	for i := 1; i <= 50; i++ {
		if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}