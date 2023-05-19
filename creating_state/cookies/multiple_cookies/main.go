package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request)  {
	http.SetCookie(w, &http.Cookie{
		Name: "my-cookie",
		Value: "my cookie value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println("my-cookie", err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", cookie1)
	}

	cookie2, err := req.Cookie("general")
	if err != nil {
		log.Println("general", err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2:", cookie2)
	}

	cookie3, err := req.Cookie("specific")
	if err != nil {
		log.Println("specific", err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3:", cookie3)
	}
}

func abundance(w http.ResponseWriter, req * http.Request)  {
	http.SetCookie(w, &http.Cookie{
		Name: "general",
		Value: "general value",
	})
	
	http.SetCookie(w, &http.Cookie{
		Name: "specific",
		Value: "specific value",
	})

	fmt.Fprintln(w, "specific and general cookies has been set")
}