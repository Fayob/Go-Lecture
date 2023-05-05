package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main()  {
	jsonStream := `{"Name":"Jim", "Age":12, "Location":"Lagos"}{"Name":"John", "Age":20, "Location":"Texas"}`

	reader := strings.NewReader(jsonStream)

	writer := os.Stdout

	decoder := json.NewDecoder(reader)
	encoder := json.NewEncoder(writer)

	for {
		var v map[string]interface{}

		err := decoder.Decode(&v)
		if err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k=="Age" {
				delete(v,k)
			}
		}
		
		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
			return
		}
	}
}