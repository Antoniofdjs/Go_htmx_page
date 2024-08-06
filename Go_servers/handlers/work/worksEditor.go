package work

import (
	"Go_servers/db"
	"Go_servers/handlers/editors"
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	// "time"
)

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
		WorkID:   r.FormValue("WorkID"),
		Component: r.FormValue("Component"),
	}

	// Search for the component and call handler
	tmplFunc, exists := editors.ComponentsHandlers[data.Component]
	if !exists {
		fmt.Println("\n\nOption not found", data.Option)
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	// Get the template and render it
	tmpl := tmplFunc(data.WorkID, templateFs)
	if tmpl == nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	fmt.Println("\n\nFound template")
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
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

	WorkID := r.FormValue("WorkID")
	if WorkID == "" {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	fmt.Println("EDIT my work id is: ",WorkID)
	title := r.FormValue("inputTitle")

	PicBytes, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		defer PicBytes.Close() // Prevent resource leak
	}

	if PicBytes == nil && title != "" {
		updated, err:= db.EditTitle(WorkID, title)
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
	workID := r.FormValue("WorkID")
	fmt.Println("workId in before delete",workID )

	err := db.DeleteWork(workID)
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