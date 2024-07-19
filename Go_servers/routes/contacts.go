package contacts

import (
	"fmt"
	"html/template"
	"net/http"
)

// Define a struct to hold the JSON data


func ContactGetHand(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./htmlTemplates/contact.html"))
	tmpl.Execute(w, nil)
}

func ContactPostHand(w http.ResponseWriter, r *http.Request) {
	// tmpl := template.Must(template.ParseFiles("./htmlTemplates/services.html"))
	   // Parse the form values
	   err := r.ParseForm()
	   if err != nil {
		   http.Error(w, "Unable to parse form", http.StatusBadRequest)
		   return
	   }

	   // Access the form values
	   name := r.FormValue("name")
	   email := r.FormValue("email")
	   message := r.FormValue("message")
   
	   // Print the form values
	   fmt.Printf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)
   
	   // Set the redirect header for HTMX
	   w.Header().Set("HX-Redirect", "/")
   
	   // Respond to the client
	   w.Write([]byte("Form submitted successfully!"))
}
