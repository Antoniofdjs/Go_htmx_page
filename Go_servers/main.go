package main

/*
Main server
	^Routes are kept here
	^Handlers are short named = <method>Hand
		example work.GetHand ====> work package, GET Method Handler
	^Html is being served from the htmlTemplates

*/

import (
	"Go_servers/handlers/about"
	contacts "Go_servers/handlers/contact"
	"Go_servers/handlers/galleries"
	"Go_servers/handlers/work"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// Use the router to serve static files
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Routes
	r.HandleFunc("/", landingHandler).Methods("GET")
	r.HandleFunc("/about", about.GetHand).Methods("GET")
	r.HandleFunc("/work", work.GetHand).Methods("GET")
	r.HandleFunc("/work", work.PostHand).Methods("POST")
	r.HandleFunc("/work/editor", work.PutHandEditor).Methods("PUT")
	r.HandleFunc("/work/editor", work.GetHandEditor).Methods("GET")
	r.HandleFunc("/work/editor", work.PostHandEditor).Methods("POST")
	r.HandleFunc("/work/del", work.DelHandEditor).Methods("POST")
	r.HandleFunc("/work/component", work.FectchComponent).Methods("POST")
	r.HandleFunc("/contact", contacts.GetHand).Methods("GET")
	r.HandleFunc("/contact", contacts.PostHand).Methods("POST")
	r.HandleFunc("/gallery/{title}", galleries.Gallery).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}
