package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "$$,Go, Programming,**"
	string := "!!Go Programming Language??"
	str1 := "Go"
	str2 := "Go"

	trim := strings.Trim(string, "!")
	trimStr := strings.Trim(str, "*")
	trimSpace := strings.TrimSuffix(str, "**")
	split := strings.SplitAfter(str, ",")
	fmt.Println(trim)
	fmt.Println(trimStr)
	fmt.Println(trimSpace)
	fmt.Println(split, len(split))
	fmt.Println(strings.Join(split, ""))
	fmt.Println(strings.Contains("split", "s"))
	fmt.Println(strings.Replace("split", "i", "a", 1))
	fmt.Println(strings.Replace("agencies", "e", "*", -1))
	fmt.Println(strings.Index("agencies", "e"))
	fmt.Println(strings.LastIndex("agencies", "e"))

	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Compare("Drone", "Hello"))
	fmt.Println(strings.Compare("drone", "Drone"))
}