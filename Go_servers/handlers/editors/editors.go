package editors

import (
	"embed"
	"fmt"
	"html/template"
	"log"
)

var Handlers = map[string]func(picId string, templateFs embed.FS) *template.Template{
	"EditTitle":   EditTitle,
	"Delete": DeleteWork,
	"ChangePic": ChangePic,
	"InsertAbove": InsertAbove,
	"InsertBelow": InsertBelow,
}


/*
	Renders the html template for the buttom 'Edit Title' in the 'buttons-container'.
*/
func EditTitle(workID string, templateFS embed.FS) *template.Template {
	fmt.Println("My Pic Id is:", workID)

	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFS(templateFS, "htmlTemplates/workEditor/editTitleWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

/*
	Renders the html template for the buttom 'Delete' in the 'buttons-container'.
*/ 
func InsertAbove(picID string, templateFs embed.FS) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFS(templateFs,"htmlTemplates/workEditor/insertAboveWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func InsertBelow(picID string, templateFs embed.FS) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFS(templateFs,"htmlTemplates/workEditor/insertBelowWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func ChangePic(picID string, templateFS embed.FS) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFS(templateFS,"htmlTemplates/workEditor/changePictureWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func DeleteWork(picID string, templateFs embed.FS) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFS(templateFs,"htmlTemplates/workEditor/deleteWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}