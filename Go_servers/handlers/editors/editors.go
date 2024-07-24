package editors

import (
	"fmt"
	"html/template"
	"log"
)

var Handlers = map[string]func(picId string) *template.Template{
	"EditTitle":   EditTitle,
	"Delete": DeleteWork,
	"ChangePic": ChangePic,
	"InsertAbove": InsertAbove,
	"InsertBelow": InsertBelow,
}


/*
	Renders the html template for the buttom 'Edit Title' in the 'buttons-container'.
*/
func EditTitle(picID string) *template.Template {
	fmt.Println("My Pic Id is:", picID)

	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/editTitleWork.html")
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
func InsertAbove(picID string) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/insertAboveWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func InsertBelow(picID string) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/insertBelowWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func ChangePic(picID string) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/changePictureWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}

func DeleteWork(picID string) *template.Template{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully
	fmt.Println("Parsing Template: ")
	tmpl, err := template.ParseFiles("htmlTemplates/components/deleteWork.html")
	if err != nil {
		// Log the error and return nil to indicate failure
		log.Printf("Error parsing template: %v", err)
		return nil
	}
	return tmpl
}