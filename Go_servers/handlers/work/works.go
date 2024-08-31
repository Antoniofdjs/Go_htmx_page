/*
'/work related stuff...
*/
package work

//  WORKING WAAAAY BELOWW GO CHECK

import (
	"Go_servers/models"
	templates "Go_servers/templ"
	"embed"
	"net/http"
	"strconv"
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

	// Change to strings all values from the work struct, in this case Position sicne its an int
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

	templates.ShowWorks(works).Render(r.Context(), w)
}