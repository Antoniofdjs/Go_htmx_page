package work

import (
	"Go_servers/db"
	editorComponents "Go_servers/handlers/editors"
	"Go_servers/models"
	templates "Go_servers/templ"
	"bytes"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/a-h/templ"
)

type DataComponents struct{
	Position  string
	BelowPosition string
	Component string
	Title string
	Description string
}

func GetTestView(w http.ResponseWriter, r *http.Request, editorFs embed.FS) {
	var works []models.WorkFrontEnd

	// Check if local data is valid if not, get data from DB
	if models.WorksStorage == nil || len(models.WorksStorage) == 0{
		fmt.Println("Fecthing data from database")
		models.WorksStorage = db.AllWorks()
	}

	// Change to strings all values of the work struct, in this case Position since its an int
	// Every value must be string for the html .templ
	for _, work := range models.WorksStorage{
		positionString := strconv.Itoa(work.Position)
		workStringsOnly := models.WorkFrontEnd{
			Title : work.Title,
			Path : work.Path,
			Description : work.Description,
			Position : positionString,
		}
		
		fmt.Println("Description BACK: ",work.Description)
		fmt.Println("Description FRONT: ",workStringsOnly.Description)
		works = append(works, workStringsOnly)
	}

	templates.ShowEditor(works, true).Render(r.Context(), w)
}

/*
	Get html template for the '/editor' view
*/ 
func GetHandEditor(w http.ResponseWriter, r *http.Request, editorFs embed.FS) {
	var works []models.WorkFrontEnd

	// Check if local data is valid if not, get data from DB
	if models.WorksStorage == nil || len(models.WorksStorage) == 0{
		fmt.Println("Fecthing data from database")
		models.WorksStorage = db.AllWorks()
	}

	// Change to strings all values of the work struct, in this case Position since its an int
	// Every value must be string for the html .templ
	for _, work := range models.WorksStorage{
		positionString := strconv.Itoa(work.Position)
		workStringsOnly := models.WorkFrontEnd{
			Title : work.Title,
			Path : work.Path,
			Description : work.Description,
			Position : positionString,
		}
		fmt.Println("Description: ",workStringsOnly.Description)
		works = append(works, workStringsOnly)
	}

	templates.ShowEditor(works, true).Render(r.Context(), w)
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
		Title: r.FormValue("Title"),
		Description: r.FormValue("Description"),
	}
	belowPositionInt, _:= strconv.Atoi(data.Position) // validate error later here......
	belowPosition:= strconv.Itoa(belowPositionInt + 1)
	
	componentData := DataComponents{
		Position: data.Position,
		BelowPosition: belowPosition,
		Title: data.Title,
		Description: data.Description,
	}
	// Search for the component and call handler
	fmt.Println("Fectgin component: ", data.Component)
	fmt.Println("Title is: ", data.Title)
	fmt.Println("Position is: ", data.Position)
	fmt.Println("Decription is: ", data.Description)

	_, exists := editorComponents.ComponentsHandlers[data.Component]
	if !exists {
		fmt.Println("\n\nOption not found", data.Option)
		http.Error(w, "Invalid option", http.StatusBadRequest)
		return
	}
	
	editorComponents := map[string]func(models.WorkFrontEnd) templ.Component{
		"EditTitle": func(work models.WorkFrontEnd) templ.Component {
			return templates.ButtonView("Edit", work)
		},
		"ButtonsEditor": func(work models.WorkFrontEnd) templ.Component {
			return templates.ButtonsContainer(work)
		},
		"Delete": func(work models.WorkFrontEnd) templ.Component {
			return templates.ButtonView("Delete", work)
		},
		"InsertAbove": func(work models.WorkFrontEnd) templ.Component {
			return templates.ButtonView("Insert", work)
		},
		"InsertBelow": func(work models.WorkFrontEnd) templ.Component {
			return templates.ButtonView("Insert", work)
		},
	}

	work := models.WorkFrontEnd{
		Title: componentData.Title,
		Position: componentData.Position,
		Description: componentData.Description,
	}
	fmt.Println("Below Position: ", componentData.BelowPosition)
	compFunc:= editorComponents[data.Component]

	templComponent := compFunc(work)
	templComponent.Render(r.Context(), w)
	
	// Get the template and render it
	// tmpl := templHandler(data.Position, templateFs)
	// if tmpl == nil {
	// 	http.Error(w, "Template error", http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Println("\n\nFound template")
	// err = tmpl.Execute(w, componentData)
	// if err != nil {
	// 	http.Error(w, "Unable to render template", http.StatusInternalServerError)
	// }
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

	title := r.FormValue("Title")
	description := r.FormValue("Description")
	position := r.FormValue("Position")
	fmt.Printf("Title: %s, Description: %s, Position: %s \n",title, description, position)

	err = db.InsertWork(title, position,fileName, fileBytes)
	if err!= nil{
		http.Error(w, "Unable to insert new work", http.StatusInternalServerError)
		return
	}
	models.WorksStorage = db.AllWorks()
	w.Header().Set("HX-Redirect", "/test")
	w.WriteHeader(http.StatusOK)
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
	description := r.FormValue("inputDescription")
	PicBytes, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		defer PicBytes.Close() // Prevent resource leak
	}

	if PicBytes == nil && title != "" {
		updated, err:= db.EditTitle(Position, title, description)
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
		// Check if local data is valid if not, get data from DB
		if models.WorksStorage == nil || len(models.WorksStorage) == 0{
			fmt.Println("Fecthing data from database")
			models.WorksStorage = db.AllWorks()
		}
		tmpl.Execute(w, models.WorksStorage)
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