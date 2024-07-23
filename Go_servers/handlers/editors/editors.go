package editors

import (
	"fmt"
	"html/template"
	"log"
)

var Handlers = map[string]func(picId string) *template.Template{
	"EditTitle":   EditTitle,
	}

func EditTitle(picID string) *template.Template {
	fmt.Println("My Pic Id is:", picID)

	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/workTitleForm.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}
