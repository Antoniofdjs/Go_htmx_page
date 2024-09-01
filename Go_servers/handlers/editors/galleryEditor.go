package editor

import (
	"Go_servers/models"
	templates "Go_servers/templ"
	"embed"
	"fmt"
	"net/http"
	"strconv"
)


func GetEditorGallery(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	gallery := models.GalleriesStorage[67] // 67 is the key and work id fro the mountains gallery at the moment
	work := models.WorksStorage[2]
	galleryItemsFront := []models.GalleryItemFrontEnd{}
	workFront := models.WorkFrontEnd{
		Title: work.Title,
		Path: work.Path,
		Position: strconv.Itoa(work.Position),
		Description: work.Description,
	}

	for _, item := range gallery{
		itemFront := models.GalleryItemFrontEnd{
			Path: item.Path,
			Position: strconv.Itoa(item.Position),
		}
		galleryItemsFront = append(galleryItemsFront, itemFront)
	}

	templates.ShowEditorGallery(workFront, galleryItemsFront).Render(r.Context(), w)
}

func UpdateGalleryItems(w http.ResponseWriter, r *http.Request){
	fmt.Println("UPDATE ELEMENT ACTIVATED")
	r.ParseForm()
	opacity:= r.FormValue("Opacity")
	picUrl:= r.FormValue("PicUrl")
	position:= r.FormValue("Position")
	fmt.Println("UPDATE ACTIVATED")
	fmt.Println("Opacity: ",opacity)
	fmt.Println("Url: ", picUrl)
	templates.UpdatePicStatus(opacity, picUrl, position).Render(r.Context(), w)
}

// Uploading multiple images from gallery
func FileUploadGallery(w http.ResponseWriter, r *http.Request){
	templates.FilesSelectedContainer().Render(r.Context(), w)

}
