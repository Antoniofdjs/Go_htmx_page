package editor

import (
	"Go_servers/db"
	"Go_servers/models"
	storageInits "Go_servers/storageInit"
	templates "Go_servers/templ"
	"embed"
	"fmt"
	"io"
	"net/http"
	"strconv"
)


func GetEditorGallery(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	var works []models.Work
	var workFront models.WorkFrontEnd
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
			workFront = models.WorkFrontEnd{
				Title : w.Title,
				Position : strconv.Itoa(w.Position),
				Description : w.Description,
				Path: w.Path,
			}
			break
		}
	}
	if workFront.Title == "" {
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
	fmt.Println("Upload Pictures Temps")
	
	err:= r.ParseMultipartForm(100<<20)
	if err!=nil{
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	workTitle:= r.FormValue("Title")
	// Get picture bytes and pic name
	fileHeaders:= r.MultipartForm.File["Pictures"]

	var fileNames []string
	models.FileTempStorage = []models.FileTemp{} // Reset local storage for files and bytes
	
	for _,fileHeader := range fileHeaders{
		fileNames = append(fileNames, fileHeader.Filename)
		fmt.Println("File: ",fileHeader.Filename)

		file, err:= fileHeader.Open()
		if err != nil {
            http.Error(w, "Could not open file", http.StatusInternalServerError)
            return
        }
        defer file.Close()	
		fileBytes, _:= io.ReadAll(file)
		
		tempFile:= models.FileTemp{
			FileName: fileHeader.Filename,
			FileBytes: fileBytes,
		}
		models.FileTempStorage = append(models.FileTempStorage, tempFile)
		fmt.Println("Appended a fiel to temp storage")
	}
	
	// I should save locally the data in a temp?

	templates.FilesSelectedContainer(fileNames, workTitle).Render(r.Context(), w) // sending the names to front before confirming submit of files

}

// Uploading multiple images from gallery
func UploadGalleryItems(w http.ResponseWriter, r *http.Request){
	fmt.Println("Upload Confirm Pictures activated")
	
	var works []models.Work
	var workTitleFound bool = false
	var workIdKey int

	title := r.PathValue("title")

	// Change work data to strings
	works = models.WorksStorage
	for _, w := range works{
		if w.Title == title{
			fmt.Println("Found work with title: ", title)
			workTitleFound = true
			workIdKey = w.Id
			break
		}
	}
	if !workTitleFound{
		fmt.Println("No work found")
		return
	}
	
	filesInserted := 0
	for _, file := range models.FileTempStorage{
			err := db.InsertGalleryItem(workIdKey, file.FileName, file.FileBytes)
			if err!=nil{
				fmt.Println("Error", err)
				fmt.Println("Total files inserted ", filesInserted)
				return
			}
			filesInserted += 1
	}
	storageInits.InitGalleries() // Finally init the gallery storage

	redirectPath:= fmt.Sprintf("/editor/%s", title)
	w.Header().Set("HX-Redirect", redirectPath)
	w.WriteHeader(http.StatusOK)
}
