package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/db"
	"embed"
	"fmt"
	"net/http"
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


func GetHand(w http.ResponseWriter, r *http.Request, fileEmbed embed.FS) {
	tmpl := template.Must(template.ParseFS(fileEmbed,"htmlTemplates/work.html"))
	pictures := db.WorksDB()

	tmpl.Execute(w, *pictures)
}

//  WORKING HEREEE
func FectchComponent(w http.ResponseWriter, r *http.Request, templateFs embed.FS){

	fmt.Println("Fecthing back buttonsEditor")

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	Data := RequestData{
		WorkID:   r.FormValue("WorkID"),
		Component: r.FormValue("Component"),
	}
	fmt.Println("Accesing values")
	fmt.Println(Data.WorkID)
	fmt.Println(Data.Component)

	tmpl := template.Must(template.ParseFS(templateFs,"htmlTemplates/workEditor/buttonsEditor.html"))
	tmpl.Execute(w, Data)
}
