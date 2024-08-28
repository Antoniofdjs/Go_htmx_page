package work

import (
	"Go_servers/db"
	"Go_servers/models"
	templates "Go_servers/templ"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type DataComponents struct{
	Position  string
	BelowPosition string
	Component string
	Title string
	Description string
}


//  This needs to replace the 'GET /editor' view
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
	var belowPosition string = ""

	// Extract values from request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	//  Data from http request
	requestData := RequestData{
		Position:   r.FormValue("Position"),
		Component: r.FormValue("Component"),
		Title: r.FormValue("Title"),
		Description: r.FormValue("Description"),
	}

	// Change postion to the next value for the InsertBelow request
	if requestData.Component == "InsertBelow"{
		belowPositionInt, _:= strconv.Atoi(requestData.Position)
		belowPosition = strconv.Itoa(belowPositionInt + 1)
		fmt.Println("Below position is: ", belowPosition)
	}
	
	fmt.Println("Fetching component: ", requestData.Component)
	fmt.Println("Title is: ", requestData.Title)
	fmt.Println("Position is: ", requestData.Position)
	fmt.Println("Decription is: ", requestData.Description)

	work := models.WorkFrontEnd{
		Title: requestData.Title,
		Position: requestData.Position,
		Description: requestData.Description,
		PositionBelow: belowPosition,
	}

	if requestData.Component == "ButtonsEditor"{
		templates.ButtonsContainer(work).Render(r.Context(), w)
	}else{
	templates.ButtonView(requestData.Component, work).Render(r.Context(), w)
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

	title := r.FormValue("Title")
	description := r.FormValue("Description")
	position := r.FormValue("Position")
	// insertBelow:= r.FormValue("InsertBelow")
	fmt.Printf("Title: %s, Description: %s, Position: %s \n",title, description, position)

	err = db.InsertWork(title, position, description, fileName, fileBytes)
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
	var fileName string = ""

	fmt.Println("PUT EDITOR")
	// time.Sleep(2 * time.Second)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	Position := r.FormValue("Position")
	title := r.FormValue("Title")
	description := r.FormValue("Description")

	if Position == "" {
		fmt.Println("Error parsing position")
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	//  Get file from form and then read bytes and get name of the picture.
	file, fileHeader, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		fmt.Println("Picture detected")
		defer file.Close() // Prevent resource leak
	}
	
	// Call db to add new picture
	if file != nil{
		fileName = fileHeader.Filename
		picBytes , err:= io.ReadAll(file)
		if err!=nil{
		http.Error(w,"Error reading picture bytes",http.StatusInternalServerError)
		return
		}
		err = db.AddPicture(fileName, picBytes)
		if err!=nil{
		http.Error(w,"Error changing picture on database",http.StatusInternalServerError)
		return
	}
	}
	fmt.Println("New Picture Name: ", fileName)
	// Edit the work object on the db
	updated, err:= db.EditTitle(Position, title, description, fileName)
	if !updated{
		fmt.Println(err)
		http.Error(w, "Error updating title", http.StatusBadRequest)
		return
	}
	// Update local storage
	models.WorksStorage = db.AllWorks()
	w.Header().Set("HX-Redirect", "/test")
	w.WriteHeader(http.StatusOK)
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