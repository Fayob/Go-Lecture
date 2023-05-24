package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tpl *template.Template

func init()  {
	filePath, err := filepath.Abs("web_toolkit/ajax/example_01/template")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseGlob(filePath + "/*"))
}

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request)  {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func foo(w http.ResponseWriter, req *http.Request)  {
	s := "Here is some text from foo"
	fmt.Fprintln(w, s)
}