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
	servicesHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("htmlTemplates/services.html"))
		tmpl.Execute(w, nil)}

	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("GET /contact", contacts.ContactGetHand)
	http.HandleFunc("/", landingHandler)
	
	http.HandleFunc("GET /services", servicesHandler)
	
	http.HandleFunc("POST /contact", contacts.ContactPostHand)
	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
