package work

import (
	"Go_servers/db"
	editorComponents "Go_servers/handlers/editors"
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	// "time"
)

type DataComponents struct{
	Position  string
	BelowPosition string
	Component string
}

/*
	Get html template for the '/editor' view
*/ 
func GetHandEditor(w http.ResponseWriter, r *http.Request, editorFs embed.FS) {
	tmpl := template.Must(template.ParseFS(editorFs,"htmlTemplates/editorTemplates/workEditor.html"))
	works := db.AllWorks()
	tmpl.Execute(w, works)
}

/*
	Get components for the editor, this includes the views of the buttons clicked and 'Buttons Editor Component'
*/ 
func GetEditorComponents(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	// Read the body of the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Could not read body")
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Print the body content
	fmt.Println("Body:", string(body))
	r.Body = io.NopCloser(io.MultiReader(bytes.NewReader(body)))

	// Extract values from request
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	data := RequestData{
		Position:   r.FormValue("Position"),
		Component: r.FormValue("Component"),
	}
	belowPositionInt, _:= strconv.Atoi(data.Position) // validate error later here......
	belowPosition:= strconv.Itoa(belowPositionInt + 1)
	
	componentData := DataComponents{
		Position: data.Position,
		BelowPosition: belowPosition,
	}

	// Search for the component and call handler
	tmplFunc, exists := editorComponents.ComponentsHandlers[data.Component]
	if !exists {
		fmt.Println("\n\nOption not found", data.Option)
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	// Get the template and render it
	tmpl := tmplFunc(data.Position, templateFs)
	if tmpl == nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	fmt.Println("\n\nFound template")
	err = tmpl.Execute(w, componentData)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}


/*
	Handler for the POST request of '/editor'.
	Here either 'Insert Above' or 'Insert Below'  is being triggered to insert new work.
*/ 
func PostHandEditor(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	fmt.Println("Post activated")
	err:= r.ParseMultipartForm(100<<20)
	if err!=nil{
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get picture bytes and pic name
	file,fileHeader,err := r.FormFile("picture")
	if err!= nil{
		http.Error(w, "Could not access file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileName := fileHeader.Filename
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Could not read file", http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	position := r.FormValue("Position")
	err = db.InsertWork(title, position,fileName, fileBytes)
	if err!= nil{
		http.Error(w, "Unable to insert new work", http.StatusInternalServerError)
	}
	tmpl , err:= template.ParseFS(templateFs,"htmlTemplates/reloads/worksSectionSucces.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return
	}
	works:= db.AllWorks()
	tmpl.Execute(w, works)
}

/*
Handle any edits for the "works" table.
Will call the database to order db operations related to PUT.
*/
func PutHandEditor(w http.ResponseWriter, r *http.Request, templateFs embed.FS) {
	fmt.Println("PUT EDITOR")
	// time.Sleep(2 * time.Second)

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	Position := r.FormValue("Position")
	if Position == "" {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	fmt.Println("EDIT my work id is: ",Position)
	title := r.FormValue("inputTitle")

	PicBytes, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		defer PicBytes.Close() // Prevent resource leak
	}

	if PicBytes == nil && title != "" {
		updated, err:= db.EditTitle(Position, title)
		if !updated{
			fmt.Println(err)
			http.Error(w, "Error updating title", http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFS(templateFs,"htmlTemplates/reloads/worksSectionSucces.html")
		if err != nil {
			log.Printf("Error parsing template: %v", err)
			return
		}
		works:= db.AllWorks()
		tmpl.Execute(w, works)
	} else if PicBytes != nil && title == "" {
		fmt.Println("Change Picture")
	} else {
		fmt.Println("No data sent")
	}
}

func DelHandEditor(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	fmt.Println("Delete Activated")
	option := r.FormValue("Component")
	fmt.Println("option component",option )
	position := r.FormValue("Position")
	fmt.Println("Position in before delete",position )

	err := db.DeleteWork(position)
	if err!=nil{
		http.Error(w, "Unable to delete work", http.StatusBadRequest)
		fmt.Printf("Error: %v\n", err)
		return
	}
	// Render template:
	tmpl, err := template.ParseFS(templateFs,"htmlTemplates/reloads/worksSectionSucces.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return
	}
	works:= db.AllWorks()
	tmpl.Execute(w, works)
}