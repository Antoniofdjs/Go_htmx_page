package storageInits

import (
	"Go_servers/db"
	"Go_servers/models"
	"fmt"
)

/* Initialize local storage Gallery Map with database data. Paths contain the urls fro the each gallery item */ 
func InitGalleries(){
	galleryItems:= db.AllGalleries()

	models.GalleriesStorage = make(map[int][]models.GalleryItem)
	for _, galleryItem := range galleryItems{
		fmt.Println("Gallery Item: ", galleryItem)
		fmt.Println("Appending item to local storage")
		fmt.Println("ITEM BEFORE MAP ", galleryItem.Path)
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