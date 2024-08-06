package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"Go_servers/handlers/about"
	contacts "Go_servers/handlers/contact"
	"Go_servers/handlers/galleries"
	"Go_servers/handlers/work"
)

// Embed all HTML files
//go:embed htmlTemplates
var templatesFS embed.FS

// Embed all files in the static directory
//// go:embed static/*
// var staticFS embed.FS

func main() {
        
	fmt.Println("SERVER LISTENING:")
	// db.AllWorks() // Testing db storage and buucker from supabase here

	landingHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFS(templatesFS,"htmlTemplates/index.html"))
	tmpl.Execute(w, nil)
	}

	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Routes
	http.HandleFunc("GET /{$}", landingHandler) // Match only exactly '/' thanks to {$}
        
	http.HandleFunc("GET /about",func(w http.ResponseWriter, r *http.Request){about.GetHand(w, r, templatesFS)})
        
	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {contacts.GetHand(w, r, templatesFS)})
	http.HandleFunc("POST /contact", func(w http.ResponseWriter, r *http.Request) {contacts.PostHand(w, r, templatesFS)})

	http.HandleFunc("GET /work", func(w http.ResponseWriter, r *http.Request){work.GetHand(w, r, templatesFS)})
	http.HandleFunc("GET /work/{title}", func(w http.ResponseWriter, r *http.Request){galleries.Gallery(w, r, templatesFS)})
	
	http.HandleFunc("GET /editor", func(w http.ResponseWriter, r *http.Request){work.GetHandEditor(w, r, templatesFS)})
	http.HandleFunc("PUT /editor", func(w http.ResponseWriter, r *http.Request){work.PutHandEditor(w, r, templatesFS)}) // work here
	http.HandleFunc("POST /editor", func(w http.ResponseWriter, r *http.Request){work.PostHandEditor(w, r, templatesFS)})
	http.HandleFunc("POST /editor/del", func(w http.ResponseWriter, r *http.Request){work.DelHandEditor(w, r, templatesFS)})
	
	http.HandleFunc("GET /editor/component", func(w http.ResponseWriter, r *http.Request){work.GetEditorComponents(w, r, templatesFS)})

	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}