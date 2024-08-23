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



func EditTitleComponent(workID string, templateFS embed.FS) bool {
	fmt.Println("My Pic Id is:", workID)
	return true
}


func InsertAboveComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}


func InsertBelowComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}


func ChangePicComponent(picID string, templateFS embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)
	// Parse the template file and handle errors gracefully

	return true
}

func DeleteWorkComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)

	return true
}

func ButtonsEditorComponent(picID string, templateFs embed.FS) bool{
	fmt.Println("My Pic Id is:", picID)

	return true
}