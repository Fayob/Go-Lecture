package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwith() int {
	sum := 0
	for _, ele := range b.amount {
		sum += int(ele)
	}
	return sum / len(b.amount)
}

func (c *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for _, ele := range c.temp {
		sum += int(ele)
	}
	return sum / len(c.temp)
}
func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for _, ele := range m.amount {
		sum += int(ele)
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main()  {
	bandwith := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dish := Dashboard{
		BandwidthUsage: bandwith,
		CpuTemp: temp,
		MemoryUsage: memory,
	}

	fmt.Printf("Average bandwidth usage: %v\n", dish.AverageBandwith())
	fmt.Printf("Average cpu temp usage: %v\n", dish.AverageCpuTemp())
	fmt.Printf("Average memory usage: %v\n", dish.AverageMemoryUsage())
}