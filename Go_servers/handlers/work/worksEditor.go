package work

import (
	"Go_servers/db"
	"Go_servers/handlers/editors"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"
	"time"
)


func GetHandEditor(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("htmlTemplates/editorTemplates/workEditor.html"))
	pictures := db.PicturesDB()
	tmpl.Execute(w, *pictures)
}

func PostHandEditor(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching EditorTitle component")
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

	Data := RequestData{
		WorkID:   r.FormValue("WorkID"),
		Option: r.FormValue("Option"),
	}

	// Check if option is in map of Editor
	tmplFunc, exists := editors.Handlers[Data.Option]
	if !exists {
		fmt.Println("\n\nOption not found", Data.Option)
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}

	// Get the template
	tmpl := tmplFunc(Data.WorkID)
	if tmpl == nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	fmt.Println("\n\nFound template")

	// Execute the template, this will get the component
	err = tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}


func PutHandEditor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT EDITOR")
	time.Sleep(2 * time.Second)

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
	workID, _ := strconv.Atoi(WorkID)
	title := r.FormValue("inputTitle")

	PicBytes, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		defer PicBytes.Close() // Prevent resource leak
	}

	if PicBytes == nil && title != "" {
		updated, err:= db.EditTitle(workID, title)
		if !updated{
			fmt.Println(err)
			http.Error(w, "Error updating title", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/work/editor", http.StatusSeeOther)
	} else if PicBytes != nil && title == "" {
		fmt.Println("Change Picture")
	} else {
		fmt.Println("No data sent")
	}
}