package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type sage struct {
	Name string
	Motto string
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

func main()  {
	filePrefix, err := filepath.Abs("template/05_composite/struct")
	if err != nil {
		log.Fatalln(err)
	}

	bob := sage{
		Name: "Bob",
		Motto: "Excellence is our trade mark",
	}

	janet := sage{
		Name: "Janet",
		Motto: "Far above ruby",
	}

	john := sage{
		Name: "John",
		Motto: "Righteousness is the key",
	}

	rick := sage{
		Name: "Rick",
		Motto: "Truth will set you free",
	}

	ryan := sage{
		Name: "Ryan",
		Motto: "Success is our result",
	}

	sages := []sage{bob, janet, john, rick, ryan}

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
	
	temp, err := template.ParseGlob(filePrefix + "/*")
	if err != nil {
		log.Fatalln(err)
	}
	
	err = temp.ExecuteTemplate(os.Stdout, "temp.gohtml", bob)
	if err != nil {
		log.Fatalln(err)
	}
	
	err = temp.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
	
	err = temp.ExecuteTemplate(os.Stdout, "tep.gohtml", allList)
	if err != nil {
		log.Fatalln(err)
	}
}