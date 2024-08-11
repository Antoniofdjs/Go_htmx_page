/*
'/work related stuff...
*/
package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/db"
	"embed"
	"net/http"
	"text/template"
)

//  Currently being used for the json data received from the fecth of '/editor/component'
type RequestData struct {
    Position      string `json:"PicID"`
    Option  string `json:"Option"`
	Component string `json:"Component"`
	Title string `json:"Title"`
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
