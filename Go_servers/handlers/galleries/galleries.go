package galleries

import (
	"Go_servers/models"
	templates "Go_servers/templ"
	"fmt"
	"net/http"
	"strconv"
)

//WOORKKK HEREEEEEEEEEEEE __________________________________________________________
/*
	Handler to fecth the gallery of a work based on the title.
*/
// func Gallery(w http.ResponseWriter, r *http.Request, templateFs embed.FS) {
// 	fmt.Println("Gallery activated")
// 	title := r.PathValue("title")

// 	fmt.Println("My title is: ", title)
// 	supaClient:= db.InitDB()
// 	workUniqueID,results,err:= supaClient.From("works").Select("ID", "exact", false).Eq("Title", title).Execute()
// 	if err!=nil{
// 		fmt.Println("Error ", err)
// 		http.Error(w, "Error in query", http.StatusInternalServerError)
// 		return
// 	}
// 	if results == 0{
// 		http.Error(w, "No results", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Println("Work unique ID: ", workUniqueID)
// 	// Fecth from the galleries table here:
// 	// galleryPicsPaths, _, err:= supaClient.From("galleries").Select("Path","",false).Filter("workUniqueID","=",string(workUniqueID)).Execute()

// 	tmpl := template.Must(template.ParseFS(templateFs,"htmlTemplates/gallery.html"))
// 	tmpl.Execute(w, nil)
// }
func Gallery(w http.ResponseWriter, r *http.Request) {
	var works []models.Work
	var work models.WorkFrontEnd
	galleryItemsFront := []models.GalleryItemFrontEnd{}
	var workIdKey int

	fmt.Println("\n\nGallery activated")
	title := r.PathValue("title")
	fmt.Println("Fecthing gallery for: ", title)


	// Change work data to strings
	works = models.WorksStorage
	for _, w := range works{
		if w.Title == title{
			fmt.Println("Found work with title: ", title)
			workIdKey = w.Id
			work = models.WorkFrontEnd{
				Title : w.Title,
				Position : strconv.Itoa(w.Position),
				Description : w.Description,
				Path: w.Path,
			}
			break
		}
	}
	if work.Title == "" {
		http.Error(w, "No work found with the given title", http.StatusNotFound)
		return
	}

	// Change data to strings
	gallery, exists:= models.GalleriesStorage[workIdKey]
	if exists{
	for _, item := range gallery{
			itemFront := models.GalleryItemFrontEnd{
				Path: item.Path,
				Position: strconv.Itoa(item.Position),
			}
			galleryItemsFront = append(galleryItemsFront, itemFront)
		}}

	templates.ShowGallery(work, false, galleryItemsFront).Render(r.Context(), w)
}