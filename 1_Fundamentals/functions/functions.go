package main

import "fmt"

// Demo functions

func double(x int) int {
	return x * x
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello there!!!")
}

// Exercise Functions

func greeting(name string) {
	fmt.Println("Welcome", name)
}

func message() string {
	return "Any message can be written here"
}

func sum(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func number() int {
	return 1
}

func twoNum() (int, int) {
	return 1, 2
}

func addition() int {
	var (
		firstNum = number()
		secondNum, thirdNum = twoNum()
	)
	return firstNum + secondNum + thirdNum
}

func main() {
	// Demo
	greet()

	square := double(6)
	fmt.Println("Square =", square)

	addOneToSquare := add(square, 1)
	fmt.Println("It should print 37", addOneToSquare)

	anotherOneToSquare := add(double(6), 1)
	fmt.Println("It should print 37 as well", anotherOneToSquare)

	// Exercise
	greeting("John")

	fmt.Println("message function", message())

	fmt.Println("sum function", sum(1, 2, 3))

	fmt.Println("number function", number())

	first, second := twoNum()
	fmt.Println("Two Number function", first, second)

	fmt.Println("Addition function", addition())

}
