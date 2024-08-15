/*
'/work related stuff...
*/
package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/db"
	"Go_servers/models"
	"embed"
	"fmt"
	"net/http"
	"text/template"
)

//  Currently being used for the json data received from the fecth of '/editor/component'
type RequestData struct {
    Position      string `json:"PicID"`
    Option  string `json:"Option"`
	Component string `json:"Component"`
	Title string `json:"Title"`
	Description string `json:"Description"`
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
func GetWorksView(w http.ResponseWriter, r *http.Request, fileEmbed embed.FS) {
    // Parse the HTML template
    tmpl := template.Must(template.ParseFS(fileEmbed, "htmlTemplates/work.html"))

	// Check if local data is valid if not, get data from DB
	if models.WorksStorage == nil || len(models.WorksStorage) == 0{
		fmt.Println("Fecthing data from database")
		models.WorksStorage = db.AllWorks()
	}

    tmpl.Execute(w, models.WorksStorage)
}