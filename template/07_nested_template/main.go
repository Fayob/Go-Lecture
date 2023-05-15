package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var tpl *template.Template

func init()  {
	filePrefix, err := filepath.Abs("template/07_nested_template")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseGlob(filePrefix + "/*.gohtml"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}