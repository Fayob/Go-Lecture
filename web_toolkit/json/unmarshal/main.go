package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	Url           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	rcvd := `{ "Width":800, "Height": 600, "Title": "View from 15th floor", 
		"Thumbnail": { "Url": "http://www.example.com/image/27482648245", "Height": 125, "Width": 100},
		"Animated": false, "IDS": [116, 943, 234, 6789]}`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	fmt.Println(data)
	for i, v := range data.IDs {
		fmt.Println(i, v)
	}
	fmt.Println(data.Thumbnail.Url)
}
