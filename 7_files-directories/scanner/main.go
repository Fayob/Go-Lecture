package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main()  {
	file, err := os.Open("temp.txt")

	log.Println(file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}