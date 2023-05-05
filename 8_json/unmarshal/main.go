package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Individual struct {
	Name string
	Age int
	Location string
}

func main()  {
	j := []byte(`{"name":"Jim","age":20,"location":"Lagos"}`)

	var i Individual
	err := json.Unmarshal(j, &i)

	if err != nil {
		log.Fatal("Unable to Decode the JSON")
	}

	fmt.Println(i.Name, i.Age, i.Location)
}