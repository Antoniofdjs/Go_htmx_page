package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/handlers/editors"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
)


var WorksMap = map[string][]Picture{
	"Pictures": {
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg", PicID: 1},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", PicID: 2},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", PicID: 3},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", PicID: 4},
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg",PicID: 5},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", PicID: 6},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", PicID: 7},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", PicID: 8},
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg", PicID: 9},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", PicID: 10},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", PicID: 11},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", PicID: 12},
	},
}

type RequestData struct {
    PicID      string `json:"PicID"`
    Option  string `json:"Option"`
	Component string `json:"Component"`
}

type Picture struct{
	PicID int
	Title string
	Path string
}

type PictureData struct {
    Title      string `json:"title"`
    Picture []byte `json:"picture"`
}


func GetHand(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("htmlTemplates/work.html"))
	pictures := WorksMap

	tmpl.Execute(w, pictures)
}
func PostHand(w http.ResponseWriter, r *http.Request) {
	var data PictureData
	
	fmt.Println("POST ACTIVATED")
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}
	fmt.Println("AFTER PARSE")
	data.Title =  r.FormValue("title")
	fmt.Printf("Tile: %v",data.Title)
	
	//  Check file of picture
	PicBytes, headers, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer PicBytes.Close() //Like a file close in python, prevent leaks

	picName := headers.Filename
	path := fmt.Sprintf("static/images/userWorks/%s", picName)
	fmt.Printf("\nPic name: %v", path)

	//  Create path for picture, picFile is destination for the picture
	picFile, err := os.Create(path)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}

	 // Send the pic bytes the 'picFile'
	 _, err = io.Copy(picFile, PicBytes)
	 if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	 }

}

func GetHandEditor(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("htmlTemplates/editorTemplates/workEditor.html"))
	pictures := WorksMap
	tmpl.Execute(w, pictures)
}

// Working here i need to know what component was activated edit title, delete pic, etc
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
		PicID:   r.FormValue("PicID"),
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
	tmpl := tmplFunc(Data.PicID)
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

//  WORKING HEREEE
func FectchComponent(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Fecthing back buttonsEditor")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	Data := RequestData{
		PicID:   r.FormValue("PicID"),
		Component: r.FormValue("Component"),
	}
	fmt.Println("Accesing values")
	fmt.Println(Data.PicID)
	fmt.Println(Data.Component)

	tmpl := template.Must(template.ParseFiles("htmlTemplates/components/buttonsEditor.html"))
	tmpl.Execute(w, Data)
}

// Working here-------------------------------------------------------------------------------------------------------
// PUTS ---------- maybe put into other folders this------------------------------------------------------------------

func PutHandEditor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT EDITOR")
	// time.Sleep(2 * time.Second)
	
	
	err:= r.ParseForm()
	if err!= nil{
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return 
	}

	PicID := r.FormValue("PicID")
	if PicID == ""{
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	PicIDint,_ := strconv.Atoi(PicID)
	Title := r.FormValue("Title")

	PicBytes, _, err := r.FormFile("picture")
	if err != nil {
		fmt.Println("No picture detected")
	} else {
		defer PicBytes.Close() // Prevent resource leak
	}
	
	if PicBytes==nil && Title!=""{
		fmt.Println("EDIT TITLE FOR PIC")
		for key,works := range WorksMap{
			for work := range works{
				if works[work].PicID == PicIDint{
					fmt.Println("Found a match for the id")
					WorksMap[key][work].Title = Title
					// http://localhost:8000/work/editor
					//  need to redirect url
					http.Redirect(w, r, "/work/editor", http.StatusSeeOther)
				}
			}
		}
	}else if  PicBytes!=nil && Title==""{
		fmt.Println("Change Picture")
	}else{
		fmt.Println("No data sent")
	}
}