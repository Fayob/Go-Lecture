package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main()  {
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println("time.Date", time.Date(2023, 5, 5, 10, 40, 0, 0, time.UTC)) 
	fmt.Println("time.Date", time.Date(time.Now().Year(), time.May, time.Now().Day(), 10, 40, 0, 0, time.UTC)) 
	str := "qrwertryrui"
	fmt.Println(strings.Count(str, "r"))
	fmt.Println(string(str[0])) 
	strArr := strings.Split(str, "")
	fmt.Println(strArr[len(strArr)-1])
	rubyDate := time.RubyDate
	fmt.Println("rubyDate:", rubyDate)

	sli := []string{"a", "b", "c"}
	fmt.Println(append(sli[:1], sli[1]))
	fmt.Println(strings.Contains(strings.Join(sli, ""), "a"))
	fmt.Println(strconv.Atoi("123"))
	total := 123
	st, _ := strconv.Atoi(string(strconv.Itoa(total)[1]))
	// fmt.Printf("%T", strconv.Atoi(st))
	fmt.Printf("%T", st)
}