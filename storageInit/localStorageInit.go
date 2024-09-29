package storageInits

import (
	"fmt"

	"Go_htmx_page/db"
	"Go_htmx_page/models"
)

/* Initialize local storage Gallery Map with database data. Paths contain the urls for each gallery item */ 
func InitGalleries(){
	galleryItems:= db.AllGalleries()

	models.GalleriesStorage = make(map[int][]models.GalleryItem)
	for _, galleryItem := range galleryItems{
		key:= galleryItem.Work_ID
		_, exists := models.GalleriesStorage[key] // Check if key doesnt exist on the local storage map
		if !exists{
			models.GalleriesStorage[key] = []models.GalleryItem{}
		}
		
		models.GalleriesStorage[key] = append(models.GalleriesStorage[key], galleryItem) // Append Item to the []GalleryItem
	}

	for keyID, gallery := range models.GalleriesStorage{
		fmt.Println("--------------------------------------------------")
		fmt.Println("Work ID KEY = ", keyID)
		for _, item := range gallery{
			fmt.Println("Pic Name: ", item.Path)
			fmt.Println("Pic Position: ", item.Position)
		}
	}
}

func InitWorksStorage(){
	fmt.Println("Initializing work map storage locally")
	works := db.AllWorks()
	models.WorksStorage = works // update the slice, this mantains the order of the pics

	if models.WorksMapStorage == nil{
		fmt.Println("Work Storage doesnt exists, creating storage")
		models.WorksMapStorage = make(map[string]models.Work)
	}else{
		fmt.Println("Map already exists, deleting all keys... and update is happening")
		for key := range models.WorksMapStorage {
			delete(models.WorksMapStorage, key)
		}
	}

	for _, work := range works{
		keyTitle:= work.Title
		_, exists := models.WorksMapStorage[keyTitle]
		if !exists{
			models.WorksMapStorage[keyTitle] = models.Work{
				Id : work.Id,
				Title: work.Title,
				Description: work.Description,
				Position: work.Position,
				Path: work.Path,
			}
		}
	}
	fmt.Println("Work Storage Map len ",len(models.WorksMapStorage))
}