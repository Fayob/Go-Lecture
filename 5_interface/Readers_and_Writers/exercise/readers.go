package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello = "hello"
	CmdGoodbye = "bye"
)

func main()  {
	// My Solution
	// r := bufio.NewReader(os.Stdin)
	// numOfCmd := 0
	// for {
	// 	input, inputErr := r.ReadString(' ')
	// 	n := strings.TrimSpace(input)
	// 	if inputErr != nil {
	// 		fmt.Println("Error reading Stdin:", inputErr)
	// 	}
	// 	if n == "q" || n == "Q" {
	// 		fmt.Printf("You have used the command line %v times", numOfCmd)
	// 		return
	// 	}
	// 	if n == "" {
	// 		fmt.Println("Please enter a value")
	// 		continue
	// 	} else {
	// 		fmt.Println("\nYou entered a value")
	// 		numOfCmd += 1
	// 	}
	// }

	// Instructor's Solution
		scanner := bufio.NewScanner(os.Stdin)
		numlines := 0
		numCommands := 0
		for scanner.Scan() {
			if strings.ToUpper(scanner.Text()) == "Q" {
				break
			} else {
				text := strings.TrimSpace(scanner.Text())

				switch text {
				case CmdHello:
					numCommands += 1
					fmt.Println("command response: hi")
				case CmdGoodbye:
					numCommands += 1
					fmt.Println("command response: bye")
				}
				if text != "" {
					numlines += 1
				}
			}
		}
	fmt.Printf("You entered %v lines\n", numlines)
	fmt.Printf("You entered %v commands\n", numCommands)
}