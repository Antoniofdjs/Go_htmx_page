package about

import (
	"embed"
	"html/template"
	"net/http"
)

func GetHand(w http.ResponseWriter, r *http.Request, templateFS embed.FS){
	tmpl := template.Must(template.ParseFS(templateFS,"htmlTemplates/about.html"))
	tmpl.Execute(w, nil)
}