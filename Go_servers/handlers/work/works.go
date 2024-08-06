/*
'/work related stuff...
*/
package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/db"
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

//  Currently being used for the json data received from the fecth of '/editor/component'
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

/*
	Get all works for the /work route
*/
func GetHand(w http.ResponseWriter, r *http.Request, fileEmbed embed.FS) {
	tmpl := template.Must(template.ParseFS(fileEmbed,"htmlTemplates/work.html"))
	works:= db.AllWorks()

	tmpl.Execute(w, works)
}

/*
	Currently being used to fetch the buttons editor component of /editor.
*/ 
func GetEditorComponents(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
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
