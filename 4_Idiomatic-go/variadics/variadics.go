package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main()  {
	a := []int{1,2,3}
	b := []int{4,5,6}
	c := []int{7,8,9}

	all := append(a, b...)

	answer := sum(all...)
	fmt.Println("Without c", answer)

	all = append(all, c...)
	answer = sum(all...)
	fmt.Println("With c", answer)

	answer = sum(10,11,12,13,14,15,16,17,18,19)
	fmt.Println("Using manual values", answer)
}