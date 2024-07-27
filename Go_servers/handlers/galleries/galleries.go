package galleries

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
	Renders the html template for the buttom 'Edit Title' in the 'buttons-container'.
*/
func Gallery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gallery activated")
	title := r.PathValue("title")

	fmt.Println("My title is: ", title)
	tmpl := template.Must(template.ParseFiles("./htmlTemplates/gallery.html"))
	tmpl.Execute(w, nil)
}
