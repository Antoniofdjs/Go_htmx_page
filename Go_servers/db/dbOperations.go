/*
DB storage simulating content, will be replaced with Postgres
*/
package db

import (
	"errors"
)

type work struct {
	WorkID int
	Title string
	Path  string
}

// My "data base"
var worksMap = map[string][]work{
	"works": {
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg", WorkID: 1},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", WorkID: 2},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", WorkID: 3},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", WorkID: 4},
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg", WorkID: 5},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", WorkID: 6},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", WorkID: 7},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", WorkID: 8},
		{Title: "BEACH", Path: "../static/images/userWorks/beach.jpg", WorkID: 9},
		{Title: "FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg", WorkID: 10},
		{Title: "ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg", WorkID: 11},
		{Title: "FOOD", Path: "../static/images/userWorks/food.jpg", WorkID: 12},
	},
}

/*
Returns a pointer to the map of the works. Acting as database table.
*/
func PicturesDB() *map[string][]work {
	return &worksMap
}


func EditTitle(workID int, title string) (bool,error){
	works_map := PicturesDB()

	if len((*works_map)["works"]) < workID || workID < 0{
		return false, errors.New("WorkID does not exist")
	}
	// Modifiying directly the work title
	workIndex := workID - 1
	(*works_map)["works"][workIndex].Title = title
	return true, nil
}