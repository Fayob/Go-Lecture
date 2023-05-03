package main

import (
	"fmt"
)

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WareHouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s SpamMail) Ship() Shipping {
	return Air
}

func (s SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WareHouseAutomator)  {
	fmt.Printf("Covey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

func mainA()  {
	mail := SpamMail{4000}
	automate(&mail)

	// waste := ToxicWaste{300}
	// automate(&waste)
}