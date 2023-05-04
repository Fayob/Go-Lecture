package main

import (
	"log"
	"os"
)

func main()  {
	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(file)
	}

	file.Close()

	error := os.Rename("test.txt", "renamed.txt")

	if err != nil {
		log.Fatal(error)
	}
}