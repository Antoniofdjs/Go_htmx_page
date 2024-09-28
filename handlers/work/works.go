/*
'/work related stuff....
*/
package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"embed"
	"net/http"
	"strconv"

	"github.com/Antoniofdjs/Go_htmx_page/models"
	templates "github.com/Antoniofdjs/Go_htmx_page/templ"
)

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
// func GetWorksView(w http.ResponseWriter, r *http.Request, fileEmbed embed.FS) {
//     // Parse the HTML template
//     tmpl := template.Must(template.ParseFS(fileEmbed, "htmlTemplates/work.html"))

// 	// Check if local data is valid if not, get data from DB
// 	if models.WorksStorage == nil || len(models.WorksStorage) == 0{
// 		fmt.Println("Fecthing data from database")
// 		models.WorksStorage = db.AllWorks()
// 	}

//     tmpl.Execute(w, models.WorksStorage)
// }
func GetWorksView(w http.ResponseWriter, r *http.Request, fileEmbed embed.FS) {
	var works []models.WorkFrontEnd

	for _, w:= range models.WorksStorage{		
		positionString := strconv.Itoa(w.Position)
		workStringsOnly := models.WorkFrontEnd{
			Title : w.Title,
			Path : w.Path,
			Description : w.Description,
			Position : positionString,
		}

		works = append(works, workStringsOnly)
	}

	templates.ShowWorks(works).Render(r.Context(), w)
}