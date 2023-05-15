package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main()  {
	filePrefix, err := filepath.Abs("template/05_composite/map")
	if err != nil {
		log.Fatalln(err)
	}

	sages := map[string]string{
		"JA": "James",
		"JO": "John",
		"MA": "Mary",
		"PA": "Paul",
		"PE": "Peter",
	}
	temp, err := template.ParseGlob(filePrefix + "/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "temp.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}