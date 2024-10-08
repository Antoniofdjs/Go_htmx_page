package contacts

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// Define a struct to hold the JSON data.

func GetHand(w http.ResponseWriter, r *http.Request, templateFs embed.FS) {
	tmpl := template.Must(template.ParseFS(templateFs,"htmlTemplates/contact.html"))
	tmpl.Execute(w, nil)
}

func PostHand(w http.ResponseWriter, r *http.Request, templateFS embed.FS) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	tmpl:= template.Must(template.ParseFS(templateFS,"htmlTemplates/contact_sent.html"))
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Print the form values for FUTURE IMPLEMENTATION HERE
	fmt.Printf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)

	// Send the html content for htmx
	tmpl.Execute(w, nil)
}
