package main

import (
	"log"
	"os"
)

func main()  {
	err := os.Mkdir("creating_directory", 0755)

	if err != nil {
		log.Fatal(err)
	}
}