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

// Exercise

type Player struct {
	name string
	health, maxHealth uint
	energy, maxEnergy uint
}

func (player *Player) addHealth(amount uint) {
	player.health += amount
	if player.health > player.maxHealth {
		player.maxHealth = player.health
	}
	fmt.Println(player.name, "Add", amount, "health ->", player.health)
}

func (player *Player) applyDamage(amount uint) {
	if player.health - amount > player.health {
		player.health = 0
	} else {
		player.health -= amount
	}
	fmt.Println(player.name, "Damage", amount, "->", player.health)
}

func (player *Player) addEnergy(amount uint) {
	player.energy += amount
	if player.energy > player.maxEnergy {
		player.maxEnergy = player.energy
	}
	fmt.Println(player.name, "Add", amount, "energy ->", player.energy)
}

func (player *Player) consumeEnergy(amount uint) {
	if player.energy - amount > player.energy {
		player.energy = 0
	} else {
		player.energy -= amount
	}
	fmt.Println(player.name, "Consume", amount, "->", player.energy)
}

func main()  {
	lot := ParkingLot{space: make([]Space, 4)}

	fmt.Println("Initial:", lot)
	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After occupied:", lot)
	
	lot.vacateSpace(2)
	fmt.Println("After vacate:", lot)

	// Exercise
	player := Player{
		name: "knight",
		health: 100,
		maxHealth: 100,
		energy: 500,
		maxEnergy: 500,
	}

	player.applyDamage(99)
	player.addHealth(10)
	player.consumeEnergy(20)
	player.addEnergy(10)

	player.consumeEnergy(9999)
}