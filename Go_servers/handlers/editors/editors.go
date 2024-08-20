/*
Components for the editor view, handlers, maps, etc...
*/
package editorComponents

import (
	"embed"
	"fmt"
)

/*
Map of handlers to fecth html components for the editor. This does not execute or render them.
*/
var ComponentsHandlers = map[string]func(workId string, templateFs embed.FS) bool{
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
// func EditTitleComponent(workID string, templateFS embed.FS) *template.Template {
// 	fmt.Println("My Pic Id is:", workID)

// 	// Parse the template file and handle errors gracefully
// 	fmt.Println("Parsing Template: ")
// 	tmpl, err := template.ParseFS(templateFS, "htmlTemplates/workEditor/editTitleWork.html")
// 	if err != nil {
// 		// Log the error and return nil to indicate failure
// 		log.Printf("Error parsing template: %v", err)
// 		return nil
// 	}
// 	return tmpl
// }
/*
	Fecth the html template for the buttom 'Edit Title' in the 'buttons-container'.
*/
func EditTitleComponent(workID string, templateFS embed.FS) bool {
	fmt.Println("My Pic Id is:", workID)
	return true
}

/*
	Fecth the html template for the buttom 'Insert Above' in the 'buttons-container'.
*/ 
func InsertAboveComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}

/*
	Fecth the html template for the buttom 'Insert Below' in the 'buttons-container'.
*/ 
func InsertBelowComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}

/*
	Fecth the html template for the buttom 'Change Picture' in the 'buttons-container'.
*/ 
func ChangePicComponent(picID string, templateFS embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}

/*
	Fecth the html template for the buttom 'Delete' in the 'buttons-container'.
*/ 
func DeleteWorkComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}

/*
Fecth the html Buttons Editor Component
*/
func ButtonsEditorComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)

	return true
}