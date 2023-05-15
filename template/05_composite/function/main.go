package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type person struct {
	Name string
}

type car struct {
	Name string
}

type lists struct {
	Persons []person
	Cars []car
}

func init()  {
	filePrefix, err := filepath.Abs("template/05_composite/function")
	if err != nil {
		log.Fatalln(err)
	}

	tpl = template.Must(template.New("").Funcs(fm).ParseFiles(filePrefix + "/tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main()  {
	benz := car{"Benz"}
	toyota := car{"Toyota"}
	lexus := car{"Lexus"}
	honda := car{"Honda"}

	ade := person{"Ade"}
	ayo := person{"Ayo"}
	ola := person{"Ola"}
	akin := person{"Akin"}

	cars := []car{toyota, lexus, benz, honda}
	persons := []person{ayo, ola, ade, akin}

	allList := lists{Persons: persons, Cars: cars}

	err := tpl.Execute(os.Stdout, allList)
	if err != nil {
		log.Fatalln(err)
	}

}