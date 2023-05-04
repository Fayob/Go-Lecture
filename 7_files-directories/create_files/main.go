package main

import (
	"log"
	"os"
)

func main()  {
	createFile, err := os.Create("create.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(createFile)
	}

	createFile.Close()
}