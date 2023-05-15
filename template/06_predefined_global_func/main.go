package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Country struct {
	Name string
	City string
}

type Num struct {
	Score1, Score2 int
}

func main()  {
	filePrefix, err := filepath.Abs("template/06_predefined_global_func")
	if err != nil {
		log.Fatalln(err)
	}

	temp, err := template.ParseGlob(filePrefix + "/*")
	if err != nil {
		log.Fatalln(err)
	}

	Nigeria := Country{"Nigeria", "Lagos"}
	USA := Country{"USA", "Texas"}
	UK := Country{"UK", "London"}
	Canada := Country{"Canada", "ottawa"}
	empty := Country{"", "no city"}

	countriesList := []Country{Nigeria, USA, UK, Canada, empty}

	firstScore := Num{20,32}

	err = temp.ExecuteTemplate(os.Stdout, "temp.gohtml", countriesList)
	if err != nil {
		log.Fatalln(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "tpl.gohtml", firstScore)
	if err != nil {
		log.Fatalln(err)
	}
}