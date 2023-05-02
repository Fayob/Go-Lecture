package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func mainA()  {
	r := bufio.NewReader(os.Stdin)
	sum := 0
	for {
		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)
		if n == "" {
			fmt.Println("Please input values")
			continue
		}
		num, convErr := strconv.Atoi(n)
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}
		if inputErr == io.EOF {
			break
		}
		if inputErr != nil {
			fmt.Println("Error reading Stdin:", inputErr)
		}
	}

	fmt.Printf("\nsum: %v", sum)

	// Exercise
	
}