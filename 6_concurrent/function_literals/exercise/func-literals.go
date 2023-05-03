package main

import (
	"fmt"
	"unicode"
)

// Exercise

type LineCallback func(line string)

func linesIterator(lines []string, callback LineCallback)  {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main()  {
	// Exercise
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"11 spaces,",
		"and 4 punctuation marks",
		"in these lines of text!",
	}

	letters := 0
	digits := 0
	punctuation := 0
	spaces := 0

	lineFunc := func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters += 1
			}
			if unicode.IsDigit(r) {
				digits += 1
			}
			if unicode.IsPunct(r) {
				punctuation += 1
			}
			if unicode.IsSpace(r) {
				spaces += 1
			}
		}
	}

	linesIterator(lines, lineFunc)
	fmt.Println(letters, "letters")
	fmt.Println(digits, "digits")
	fmt.Println(punctuation, "punctuation")
	fmt.Println(spaces, "spaces")

}