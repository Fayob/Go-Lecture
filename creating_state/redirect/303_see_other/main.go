package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var tpl *template.Template

func init() {
	fileDir, err := filepath.Abs("creating_state/redirect/303_see_other")
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseGlob(fileDir + "/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

// func bar(w http.ResponseWriter, req *http.Request)  {
// 	fmt.Println("Your request method at bar:", req.Method)
// 	// process form submission here
// 	w.Header().Set("Location", "/")
// 	w.WriteHeader(http.StatusSeeOther)
// }

// Another of using see other (303) redirect
func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}