package main

import "fmt"

// Demo
func average(num1, num2, num3 int) float32 {
	return float32(num1 + num2 + num3) / 3
}

// Exercise
func accessGranted() {
	fmt.Println("Access Granted")
}

func accessDenied() {
	fmt.Println("Access Denied")
}

const (
	Monday = 0
	Tuesday = 1
	Wednesday = 2
	Thursday = 3
	Friday = 4
	Saturday = 5
	Sunday = 6
)

const (
	Admin = 10
	Manager = 20
	Contractor = 30
	Member = 40
	Guest = 50
)

func weekday(day int) bool {
	return day <= 4
}


func main() {
	// Demo

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz1 is higher than quiz2")
		} else if quiz1 < quiz2 {
		fmt.Println("Quiz1 is lesser than quiz2")
		} else {
		fmt.Println("Quiz1 is equal to quiz2")
	}

	if average(quiz1, quiz2, quiz3) > 9 {
		fmt.Println("Quiz average is greater than 9")
		} else {
		fmt.Println("Quiz average is not greater than 9")
	}

	// Exercise
	today, role := Monday, Guest

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !weekday(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}
}