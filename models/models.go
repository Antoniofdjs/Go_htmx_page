/*
All base struct models and storage for structs
IMPORTANT local storages for Object.path contains the url for the picture
In the database, "Path" is the picture name, not the same
*/
package models

// Use this for query on "works" table. Representation of the "works" table
type Work struct {
	Id int `json:"ID"`
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position int    `json:"Position"`
}

var WorksMapStorage map[string]Work // Find works based on the title

// Use this for query on "galleries" table. Representation of the "galleries" table
type GalleryItem struct{
	Path string `json:"Path"`
	Position int `json:"Position"`
	Work_ID int `json:"Work_ID"` // Reference key to table "works"
	ID int `json:"id"`
}


// Need this structs to send strings for front end HTML content.
type WorkFrontEnd struct{
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position string    `json:"Position"`
	PositionBelow string `json:"PositionBelow"`
}

type GalleryItemFrontEnd struct{
	Path string `json:"Path"`
	Position string `json:"Position"`
}

/* Local storage for works. <work.Path> is storing the public bucket url, not the just the picture name*/
var WorksStorage []Work

/* Local storage for galleries assocciate with works_id and []Gallery. <Gallery.Path> is storing the public bucket url, not the the picture name*/
var GalleriesStorage map[int][]GalleryItem

// Contains picture name(file name) and the bytes for files that will be uploaded
type FileTemp struct{
	FileName string
	FileBytes []byte
}

// Local storage temp for uploaded picture files
var FileTempStorage []FileTemp

// Store the work title and positions of gallery items to delete
var DeleteGalleryItemTempStorage map[string][]int