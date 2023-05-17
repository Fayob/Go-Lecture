package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type hotdog int
var tpl *template.Template

func init()  {
	filePrefix, err := filepath.Abs("http_server/request/parse_form")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, req.Form)
}

func main()  {
	var d hotdog
	http.ListenAndServe(":8080", d)
}