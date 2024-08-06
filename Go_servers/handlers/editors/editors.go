/*
Fecth all components for the editor
*/
package editors

import (
	"embed"
	"fmt"
	"html/template"
	"log"
)

/*
Map of handlers to fecth components for the editor
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
	Renders the html template for the buttom 'Edit Title' in the 'buttons-container'.
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
	Renders the html template for the buttom 'Insert Above' in the 'buttons-container'.
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
Fecthes Buttons Editor Component
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