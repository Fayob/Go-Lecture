package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

func main()  {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request)  {
	fileDir, err := filepath.Abs("/serve_files/exercise")
	if err != nil {
		log.Println(err)
	}
	tpl, err := template.ParseFiles(fileDir + "dog.gohtml")
	if err != nil {
		log.Println(err)
	}
	tpl.Execute(w, nil)
}