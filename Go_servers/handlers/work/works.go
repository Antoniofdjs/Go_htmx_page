package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

// Structs for json request
type RequestData struct {
    PicID      string `json:"PicID"`
    EditTitle  string `json:"EditTitle"`
	Component string `json:"Component"`
}

type Picture struct{
	PicID uint16
	Title string
	Path string
}

type PictureData struct {
    Title      string `json:"title"`
    Picture []byte `json:"picture"`
}


func GetHand(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("htmlTemplates/work.html"))
	pictures := map[string][]Picture{
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
	pictures:= map[string][]Picture{
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
	tmpl.Execute(w, pictures)
}

func PostHandEditor(w http.ResponseWriter, r *http.Request){
	fmt.Println("Fetching EditorTitle component")
	// Parse URL-encoded form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}
	// Extract form values
	Data := RequestData{
		PicID:   r.FormValue("PicID"),
		EditTitle: r.FormValue("EditPic"),
	}
	fmt.Println("My Pic Id is:", Data.PicID)
	fmt.Println("My Pic Id is:", Data.EditTitle)
    

    // Load and execute the HTML template
    tmpl := template.Must(template.ParseFiles("htmlTemplates/components/workTitleForm.html"))
    tmpl.Execute(w, Data)
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