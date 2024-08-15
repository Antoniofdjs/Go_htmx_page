/*All base struct models and storage for structs*/
package models

// Use this for query on "works" table
type Work struct {
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Description string `json:"Description"`
	Position int    `json:"Position"`
}


/* Local storage for works*/
var WorksStorage []Work
