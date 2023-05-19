package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

var tpl *template.Template

func init()  {
	filePrefix, err := filepath.Abs("creating_state/through_form/another_example")
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
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := tpl.Execute(w, person{f,l,s})
	if err != nil {
		http.Error(w, "File not found", 404)
		log.Fatalln(err)
	}
}
