package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main()  {
	add := Operation(Add)
	fmt.Println(add.calculate(7, 3))
	

	subtract := Operation(Subtract)
	fmt.Println(subtract.calculate(10, 4))

	multiply := Operation(Multiply)
	fmt.Println(multiply.calculate(4, 4))
	
	divide := Operation(Divide)
	fmt.Println(divide.calculate(200, 2))
	
}