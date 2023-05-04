package main

import (
	"log"
	"os"
)

func main()  {
	// delete file

	fileErr := os.Remove("text.txt")

	if fileErr != nil {
		log.Panicln(fileErr)
	}

	// delete folder if folder is empty

	folderErr := os.Remove("temp")

	if folderErr != nil {
		log.Panicln(folderErr)
	}
}