package galleries

import (
	"Go_servers/models"
	templates "Go_servers/templ"
	"fmt"
	"net/http"
	"strconv"
)


func Gallery(w http.ResponseWriter, r *http.Request) {

	var workFront models.WorkFrontEnd
	galleryItemsFront := []models.GalleryItemFrontEnd{}

	fmt.Println("\n\nGallery activated")
	title := r.PathValue("title")
	fmt.Println("Fecthing gallery for: ", title)

	// check if work title exists on works map
	works := models.WorksMapStorage
	work, exists := works[title]
	if !exists{
		fmt.Println("Now work found, no such gallery: ", title)
		http.Error(w, "No work found with the given title", http.StatusNotFound)
		return
	}

	// Perpare work for front, change to string all values
	workFront = models.WorkFrontEnd{
		Title : work.Title,
		Position : strconv.Itoa(work.Position),
		Description : work.Description,
		Path: work.Path,
	}
	
	// Retrieve all gallery items
	gallery, exists:= models.GalleriesStorage[work.Id]
	if exists{
	for _, item := range gallery{
		itemFront := models.GalleryItemFrontEnd{
			Path: item.Path,
			Position: strconv.Itoa(item.Position), // change positions to string
		}
		galleryItemsFront = append(galleryItemsFront, itemFront)
	}}

	templates.ShowGallery(workFront, false, galleryItemsFront).Render(r.Context(), w)
}

func GetModal(w http.ResponseWriter, r *http.Request){
	fmt.Println("Getting Modal: ")
	r.ParseForm()
	picPath := r.FormValue("Path")
	templates.ModalImage(picPath).Render(r.Context(), w)
}