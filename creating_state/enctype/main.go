package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tpl *template.Template

func init()  {
	filePrefix, err := filepath.Abs("creating_state/enctype")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := tpl.Execute(w, body)
	if err != nil {
		http.Error(w, "File not found", 404)
		log.Fatalln(err)
	}
}
