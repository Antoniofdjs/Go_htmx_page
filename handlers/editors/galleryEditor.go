package editor

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"Go_htmx_page/db"
	"Go_htmx_page/models"
	storageInits "Go_htmx_page/storageInit"
	templates "Go_htmx_page/templates"
)


func GetEditorGallery(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	// var works []models.Work
	var workFront models.WorkFrontEnd
	galleryItemsFront := []models.GalleryItemFrontEnd{}

	fmt.Println("\n\nGallery activated")
	title := r.PathValue("title")
	fmt.Println("Fecthing gallery for: ", title)
	// Reset temp delete storage
	for key:= range models.DeleteGalleryItemTempStorage{
		delete(models.DeleteGalleryItemTempStorage, key)
	}
	fmt.Println("Deleted all keys from local storage")

	works := models.WorksMapStorage
	work, exists := works[title]
	if !exists{
		fmt.Println("Now work found, no such gallery: ", title)
		http.Error(w, "No work found with the given title", http.StatusNotFound)
		return
	}

	workFront = models.WorkFrontEnd{
			Title : work.Title,
			Position : strconv.Itoa(work.Position),
			Description : work.Description,
			Path: work.Path,
		}
	
		gallery, exists:= models.GalleriesStorage[work.Id]
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
	opacity:= r.FormValue("Opacity") // if opacity is true, item marked for delete
	picUrl:= r.FormValue("PicUrl")
	position:= r.FormValue("Position") // position of gallery item
	positionInt,_:= strconv.Atoi(position)
	workTitle:= r.FormValue("WorkTitle")
	fmt.Println("Opacity: ",opacity)
	fmt.Println("Url: ", picUrl)

	_, exists := models.DeleteGalleryItemTempStorage[workTitle]
	if !exists{
		fmt.Printf("Key for %s doesnt exist, making one",workTitle)
		models.DeleteGalleryItemTempStorage[workTitle] = []int{} // create key and empty []int
	}

	if opacity == "true"{
	models.DeleteGalleryItemTempStorage[workTitle] = append(models.DeleteGalleryItemTempStorage[workTitle], positionInt) // add position for future delete confirm
	fmt.Println("Appeding to map position: ",positionInt)
	} 

	if opacity == "false"{
	var newStorageList []int
	for _, positionItem:= range models.DeleteGalleryItemTempStorage[workTitle]{
		if positionItem!= positionInt{
			newStorageList = append(newStorageList, positionItem) 	// skip the position sent, no longer will be deleted this item
		}
	}
	models.DeleteGalleryItemTempStorage[workTitle] = newStorageList
	fmt.Println("Removing from map position: ",positionInt)
	}

	templates.UpdatePicStatus(opacity, picUrl, position, workTitle).Render(r.Context(), w)
}

// Uploading multiple images from gallery into a local storage , await confirm from user. Pics are uplaoded in the PostHandGalleryEditor
func FileUploadTemporaryStorage(w http.ResponseWriter, r *http.Request){
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

	templates.FilesSelectedContainer(fileNames, workTitle).Render(r.Context(), w) // sending the names to front before confirming submit of files

}

func PutHandGalleryEditor(w http.ResponseWriter, r *http.Request){ // Working as the delete for items in the gallery, missing edits for title, description
	
	var galleryKey int
	var workIdstring string
	var updatedGalleryItems []models.GalleryItem
	var j int = 0
	var newposition int = 1

	workTitle := r.PathValue("title")
	fmt.Println("PUT Gallery Activated: Deleting items from gallery")
	fmt.Println("Deleting gallery items positions: ")
	fmt.Println("Title to delete pics from: ", workTitle)
	
	// Check if work exists before editing the gallery
	work, exists:= models.WorksMapStorage[workTitle]
	if !exists{
		fmt.Println("Now work found, no such gallery: ", workTitle)
		http.Error(w, "No work found with the given title", http.StatusNotFound)
		return
	}
	galleryKey = work.Id
	workIdstring = strconv.Itoa(galleryKey)

	galleryItems:= models.GalleriesStorage[galleryKey] // get items for the specific gallery of the work
	deletePositions, exists := models.DeleteGalleryItemTempStorage[workTitle] // get items marked for delete from temporary delete storage
	if !exists || len(models.DeleteGalleryItemTempStorage[workTitle]) == 0{
		fmt.Println("No changes detected (Testing case: No pictures marked for delete)")
		return
	}
	sort.Ints(deletePositions)

	for _, galleryItem := range galleryItems{
		if j >= len(deletePositions) || galleryItem.Position != deletePositions[j]{ // item appended to updated local gallery
			fmt.Println("Appending to updated storage", galleryItem.Position)
			galleryItem.Position = newposition
			newposition +=1
			updatedGalleryItems = append(updatedGalleryItems, galleryItem) 
			continue
		}else{ 																		// else delete item
			fmt.Println("Deleting item", deletePositions[j])
			positionToDelete := strconv.Itoa(deletePositions[j])
			args:= strings.Split(galleryItem.Path, "/")
			picName := args[len(args)-1]

			err:= db.DeleteGalleryItem(workIdstring,positionToDelete,picName)
			if err!=nil{
				fmt.Println("Total works deleted: ", j)
				for key:= range models.DeleteGalleryItemTempStorage{
					delete(models.DeleteGalleryItemTempStorage, key)
				}
				fmt.Println("Deleted all keys from local storage")
				return 
			}
			j+= 1 // next item on the list of deletes
		}
	}
	
	models.GalleriesStorage[galleryKey] = updatedGalleryItems
	db.UpdateGalleryPositions(models.GalleriesStorage[galleryKey]) // shift positions in database after deletes
	fmt.Println("Finished deleting")

	// Reset Local Delete Storage 
	for key:= range models.DeleteGalleryItemTempStorage{
		delete(models.DeleteGalleryItemTempStorage, key)
	}
	fmt.Println("Deleted all keys from local DeleteGalleryItemTempStorage")

	redirectroute := fmt.Sprintf("/editor/%s", workTitle)
	w.Header().Set("HX-Redirect", redirectroute)
}

// Uploading multiple images from gallery
func PostHandGalleryEditor(w http.ResponseWriter, r *http.Request){
	var workIdKey int
	title := r.PathValue("title")

	fmt.Println("Upload Confirm Pictures activated")

	work, exists:= models.WorksMapStorage[title]
	if !exists{
		fmt.Println("Now work found, no such gallery: ", title)
		http.Error(w, "No work found with the given title", http.StatusNotFound)
		return
	}
	workIdKey = work.Id
	// make a map to search pic names faster when the new ones come in
	galleryItems := models.GalleriesStorage[workIdKey]
	galleryPaths:= make(map[string]bool)

	for _, item := range galleryItems{
		args := strings.Split(item.Path, "/")
		picName:= args[len(args) - 1]
		galleryPaths[picName] = true
	}

	filesInserted := 0
	for _, file := range models.FileTempStorage{
		fileName := file.FileName
		_, exists:= galleryPaths[file.FileName]
		if exists{
			fileName = fmt.Sprintf("Duplicate-%s", fileName)
		}
		err := db.InsertGalleryItem(workIdKey, fileName, file.FileBytes)
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
