package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var temp *template.Template

func main() {
	filePrefix, err := filepath.Abs("template/04_using_variable/")
	if err != nil {
		log.Fatalln(err)
	}

	temp, err = template.ParseGlob(filePrefix + "/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "tpl.gohtml", "This is the data I'm passing into the template")
	if err != nil {
		log.Fatalln(err)
	}
}
