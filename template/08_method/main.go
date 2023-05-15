package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type person struct {
	Name string
	Age int
}

var tmp *template.Template

func (p person) SomeProcessing() int {
	return 7
}

func (p person) DoubleAge(x int) int {
	return x * 2
}

func init()  {
	filePrefix, err := filepath.Abs("template/08_method")
	if err != nil {
		log.Fatalln(err)
	}
	tmp = template.Must(template.ParseFiles(filePrefix + "/tmp.gohtml"))
}

func main()  {
	person := person{"James", 45}

	err := tmp.Execute(os.Stdout, person)
	if err != nil {
		log.Fatalln(err)
	}
}