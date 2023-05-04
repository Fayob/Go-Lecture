package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	newFile, err := os.Create("temp.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("create", newFile)
	}

	len, err := newFile.WriteString("This message was written using os package in Go")
	fmt.Println("len of characters =", len)  

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Stat("temp.txt")

	if os.IsNotExist(err) {
		log.Fatal(err)
		log.Fatal("File does not exist")
	}

	log.Println("stat", file)
	fmt.Println("Permission:", file.Mode())
	fmt.Printf("Size:%d\n", file.Size())
	fmt.Printf("Name:%s\n", file.Name())
	fmt.Printf("Modification Time:%s\n", file.ModTime())
}