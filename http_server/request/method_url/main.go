package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL					*url.URL
		Submissions url.Values
	}{
		req.Method,
		req.URL,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	filePrefix, err := filepath.Abs("http_server/request/method_url")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseFiles(filePrefix + "/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}