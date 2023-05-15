package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

var temp *template.Template

// func init()  {
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	temp = template.Must(template.ParseFiles(filePrefix + "/tpl.gohtml"))
// }

func main() {
	// using ParseFiles
	filePrefix, err := filepath.Abs("template/02_parsefile/")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err := template.ParseFiles(filePrefix + "/one.gohtml")
	if err != nil {
		fmt.Println(err)
	}

	nf, err := os.Create(filePrefix + "/newfile.html")
	if err != nil {
		fmt.Println(err)
	}

	err = tpl.Execute(nf, nil)

	if err != nil {
		fmt.Println(err)
	}

	// Using ParseGlob
	temp, err = template.ParseGlob(filePrefix + "/*.gohtml")

	if err != nil {
		fmt.Println(err)
	}

	err = temp.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = temp.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	err = temp.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		fmt.Println(err)
	}
}
