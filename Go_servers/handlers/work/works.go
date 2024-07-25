package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/db"
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

type RequestData struct {
    WorkID      string `json:"PicID"`
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
	pictures := db.PicturesDB()

	tmpl.Execute(w, *pictures)
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

//  WORKING HEREEE
func FectchComponent(w http.ResponseWriter, r *http.Request){
	
	fmt.Println("Fecthing back buttonsEditor")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	Data := RequestData{
		WorkID:   r.FormValue("PicID"),
		Component: r.FormValue("Component"),
	}
	fmt.Println("Accesing values")
	fmt.Println(Data.WorkID)
	fmt.Println(Data.Component)

	tmpl := template.Must(template.ParseFiles("htmlTemplates/components/buttonsEditor.html"))
	tmpl.Execute(w, Data)
}
