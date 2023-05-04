package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func createFile(filename, text string)  {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	defer file.Close()

	file.WriteString(text)
}

func readFile(filename string)  {
	fmt.Println("Reading file .................")
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panicf("Failed reading data from file: %s", err)
	}

	fmt.Printf("\nFile name: %s", filename)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
}

func main()  {
	fmt.Println("Enter file name:")
	var filename string
	fmt.Scanln(&filename)

	fmt.Println("Enter Text:")
	inputReader := bufio.NewReader(os.Stdin)

	input, _ := inputReader.ReadString('\n')

	createFile(filename, input)

	readFile(filename)
}