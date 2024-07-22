package work

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

type Picture struct{
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
			{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg"},
			{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
			{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
			{Title: "FOOD", Path: "../static/images/userWorks/food.jpg"},
			{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg"},
			{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
			{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
			{Title: "FOOD", Path: "../static/images/userWorks/food.jpg"},
			{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg"},
			{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
			{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
			{Title: "FOOD", Path: "../static/images/userWorks/food.jpg"},
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
			{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
			{Title:"FOREST", Path: "../static/images/userWorks/forest.jpg"},
			{Title:"ICELAND", Path: "../static/images/userWorks/iceland.jpg"},
			{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
			{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
			{Title:"FOREST", Path: "../static/images/userWorks/forest.jpg"},
			{Title:"ICELAND", Path: "../static/images/userWorks/iceland.jpg"},
			{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
		},
	}
	tmpl.Execute(w, pictures)
}