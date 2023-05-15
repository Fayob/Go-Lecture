package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main()  {
	filePrefix, err := filepath.Abs("template/05_composite/slice")
	if err != nil {
		log.Fatalln(err)
	}

	names := []string{"James", "John", "Mary", "Paul", "Peter"}

	temp, err := template.ParseGlob(filePrefix + "/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "tpl.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}