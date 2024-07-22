package about

import (
	"html/template"
	"net/http"
)

func GetHand(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("htmlTemplates/about.html"))
	tmpl.Execute(w, nil)
}