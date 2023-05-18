package main

import (
	"log"
	"net/http"
)

func main()  {
	// When serving a static file, it needs to be named index.html
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("./serve_files/static_file"))))
}