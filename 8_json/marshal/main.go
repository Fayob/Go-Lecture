package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string `json:"firstName"`
	Age      int
	Location string	`json:"Location,omitempty"`
}

func main() {
	firstPerson := Person{"John", 14, "Lagos"}
	
	firstPersonArray, err := json.Marshal(firstPerson)
	
	if err != nil {
		log.Fatalf("unable to encode")
	}
	
	fmt.Println(string(firstPersonArray))
	
	anotherPerson := Person{"Jim", 20, ""}

	anotherPersonArray, err := json.Marshal(anotherPerson)
	
	if err != nil {
		log.Fatalf("unable to encode")
	}
	
	fmt.Println(string(anotherPersonArray))
}
