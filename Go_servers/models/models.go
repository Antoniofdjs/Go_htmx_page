/*All base struct models and storage for structs*/
package models

// Use this for query on "works" table. Representation of the "works" table
type Work struct {
	Id int `json:"ID"`
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position int    `json:"Position"`
}
// Use this for query on "galleries" table. Representation of the "galleries" table
type GalleryItem struct{
	Path string `json:"Path"`
	Position int `json:"Position"`
	Work_ID int `json:"Work_ID"` // Reference key to table "works"
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

type FileTemp struct{
	FileName string
	FileBytes []byte
}

var FileTempStorage []FileTemp