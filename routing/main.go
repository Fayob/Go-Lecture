package main

import (
	"io"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "dog dog dog")
}

func cat(res http.ResponseWriter, req *http.Request)  {
	io.WriteString(res, "cat cat cat")
}

func main()  {
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)
}