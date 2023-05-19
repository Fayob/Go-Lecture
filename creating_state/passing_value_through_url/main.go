package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	value := req.FormValue("q")
	io.WriteString(w, "You search for => " + value)
}

// http://localhost:8080/?q=dog
// http://localhost:8080/?q=who%20is%20the%20president%20of%20nigeria