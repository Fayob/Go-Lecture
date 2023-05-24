package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	FName string
	LName string
	Items []string
}

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>Example of json marshal and encoding</h1>
	</body>
	</html>`
	w.Write([]byte(s))
}

func marshal(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	json, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		fmt.Println(err)
	}
}