package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Exercise

type Time struct {
	hour, minute, second int
}

type  TimeParseError struct {
	msg string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

// input is a string of time in this format: 14:07:33
func ParseTime(input string) (Time, error)  {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time component", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
		} 
		minutes, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minutes: %v", err), input}
		} 
		seconds, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing seconds: %v", err), input}
		} 
		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hours out of range, should be between 0 and 23", fmt.Sprintf("%v", hour)}
		}
		if minutes > 59 || minutes < 0 {
			return Time{}, &TimeParseError{"minutes out of range, should be between 0 and 59", fmt.Sprintf("%v", minutes)}
		}
		if seconds > 59 || seconds < 0 {
			return Time{}, &TimeParseError{"seconds out of range, should be between 0 and 59", fmt.Sprintf("%v", seconds)}
		}

		return Time{hour, minutes, seconds}, nil
	}
}

func main()  {
	
}