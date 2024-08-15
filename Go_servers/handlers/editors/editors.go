/*
Components for the editor view, handlers, maps, etc...
*/
package editorComponents

import (
	"embed"
	"fmt"
	"html/template"
	"log"
)

/*
Map of handlers to fecth html components for the editor. This does not execute or render them.
*/
var ComponentsHandlers = map[string]func(workId string, templateFs embed.FS) *template.Template{
	"EditTitle":   EditTitleComponent,
	"Delete": DeleteWorkComponent,
	"ChangePic": ChangePicComponent,
	"InsertAbove": InsertAboveComponent,
	"InsertBelow": InsertBelowComponent,
	"ButtonsEditor": ButtonsEditorComponent,
}


/*
	Fecth the html template for the buttom 'Edit Title' in the 'buttons-container'.
*/
func EditTitleComponent(workID string, templateFS embed.FS) *template.Template {
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
	Fecth the html template for the buttom 'Insert Above' in the 'buttons-container'.
*/ 
func InsertAboveComponent(picID string, templateFs embed.FS) *template.Template{
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

/*
	Fecth the html template for the buttom 'Insert Below' in the 'buttons-container'.
*/ 
func InsertBelowComponent(picID string, templateFs embed.FS) *template.Template{
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

/*
	Fecth the html template for the buttom 'Change Picture' in the 'buttons-container'.
*/ 
func ChangePicComponent(picID string, templateFS embed.FS) *template.Template{
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

/*
	Fecth the html template for the buttom 'Delete' in the 'buttons-container'.
*/ 
func DeleteWorkComponent(picID string, templateFs embed.FS) *template.Template{
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

/*
Fecth the html Buttons Editor Component
*/
func ButtonsEditorComponent(picID string, templateFs embed.FS) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	tmpl,err:= template.ParseFS(templateFs,"htmlTemplates/workEditor/buttonsEditor.html")
	if err!=nil{
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}