/*All base struct models and storage for structs*/
package models

// Use this for query on "works" table
type Work struct {
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position int    `json:"Position"`
}

// Need this struct to send strings for front end HTML content, its the same as work but only strings
type WorkFrontEnd struct{
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position string    `json:"Position"`
}

/* Local storage for works*/
var WorksStorage []Work
