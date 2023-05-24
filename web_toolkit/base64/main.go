package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main()  {
	s := "Here is the text i wanna encode and decode for this file"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("An error occur", err)
	}
	fmt.Println(string(bs))
}