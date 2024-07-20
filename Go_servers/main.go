package main

import (
	contacts "Go_servers/routes"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Define a struct to hold the JSON data

	landingHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl := template.Must(template.ParseFiles("htmlTemplates/index.html"))
			tmpl.Execute(w, nil)
		}
	}

	// Handlers
	workHand := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("htmlTemplates/work.html"))
		tmpl.Execute(w, nil)
	}

	aboutHand := func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("htmlTemplates/about.html"))
		tmpl.Execute(w, nil)
	}
	
	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", landingHandler)
	
	http.HandleFunc("GET /about", aboutHand)
	
	http.HandleFunc("GET /work", workHand)
	
	http.HandleFunc("GET /contact", contacts.ContactGetHand)
	http.HandleFunc("POST /contact", contacts.ContactPostHand)
	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
