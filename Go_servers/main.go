package main

/*
Main server
	^Routes are kept here
	^Handlers are short named = <mehtod>Hand
		example work.GetHand ====> work package, GET Method Handler
	^Html is being served from the htmlTemplates

*/

import (
	"Go_servers/handlers/about"
	contacts "Go_servers/handlers/contact"
	"Go_servers/handlers/work"
	"html/template"
	"log"
	"net/http"
)

func main() {
	
	landingHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl := template.Must(template.ParseFiles("htmlTemplates/index.html"))
			tmpl.Execute(w, nil)
		}
	}
	
	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Routes
	http.HandleFunc("/", landingHandler)
	
	http.HandleFunc("GET /about", about.GetHand)

	http.HandleFunc("GET /work", work.GetHand)
	http.HandleFunc("POST /work", work.PostHand)

	http.HandleFunc("PUT /work/editor", work.PutHandEditor)
	http.HandleFunc("GET /work/editor", work.GetHandEditor)
	http.HandleFunc("POST /work/editor", work.PostHandEditor)
	http.HandleFunc("POST /work/del", work.DelHandEditor)

	http.HandleFunc("POST /work/component", work.FectchComponent)

	http.HandleFunc("GET /contact", contacts.GetHand)
	http.HandleFunc("POST /contact", contacts.PostHand)
	

	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}