package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	space []Space
}

func occupySpace(lot *ParkingLot, spaceNum int)  {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = true
}

func (lot *ParkingLot) vacateSpace(spaceNum int) {
	lot.space[spaceNum-1].occupied = false
}


func main()  {
	lot := ParkingLot{space: make([]Space, 4)}

	fmt.Println("Initial:", lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied:", lot)
	
	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)
}