package main

import "fmt"

func main()  {
	var myName = "Fayob"
	fmt.Println("My name is", myName)

	var name string = "Favour"
	fmt.Println("name =", name)

	userName := "admin"
	fmt.Println("Username is", userName)

	var sum int
	fmt.Println("The sum =", sum)

	part1, other := 1, 5
	fmt.Println("part 1 =", part1, "other is", other)

	part2, other := 2, 5
	fmt.Println("part 2 =", part2, "other is", other)

	sum = part1 + part2
	fmt.Println("The new sum =", sum)

	var (
		lessonName = "Variable"
		lessonType = "Demo"
	)

	fmt.Println("Lesson name is", lessonName)
	fmt.Println("Lesson type is", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)

	// Exercise

	var color string = "Yellow"
	fmt.Println("My favorite color is", color)

	birthYear, age := 2000, 23
	fmt.Println("Birth year is", birthYear, "age is", age)

	var (
		first = "first"
		last = "last"
	)
	fmt.Println(first, "and", last)

	var yourAge int
	yourAge = 24 * age
	fmt.Println("Your age is", yourAge)
}
