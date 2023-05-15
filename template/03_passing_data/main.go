package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var temp *template.Template

func init() {
	filePrefix, err := filepath.Abs("template/03_passing_data/")
	if err != nil {
		log.Fatalln(err)
	}
	temp = template.Must(template.ParseFiles(filePrefix + "/tpl.gohtml"))
}

func main() {
	err := temp.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
