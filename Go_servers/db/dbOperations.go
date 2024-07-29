/*
Database handler
*/
package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/postgrest-go"
	"github.com/supabase-community/supabase-go"
)

// Use this for query on "works" table
type WorkQuery struct {
	CreatedAt string `json:"CreatedAt"`
	ID        int    `json:"ID"`
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	WorkID int    `json:"WorkID"`
}
// Use this to send works to html in a slice []work
type Work struct {
    Path   string `json:"Path"`
    Title     string `json:"Title"`
    WorkID int    `json:"WorkID"`
}

// My "data base"
var worksMap = map[string][]Work{
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
func WorksDB() *map[string][]Work {
	return &worksMap
}


func EditTitle(workID int, title string) (bool,error){
	works_map := WorksDB()

	if len((*works_map)["works"]) < workID || workID < 0{
		return false, errors.New("WorkID does not exist")
	}
	// Modifiying directly the work title
	workIndex := workID - 1
	(*works_map)["works"][workIndex].Title = title
	return true, nil
}

func DeleteWork(workID int) error{
	works_map:= WorksDB()
	if len((*works_map)["works"]) < workID || workID < 0{
		return errors.New("WorkID does not exist")
	}
	works:= (*works_map)["works"]
	fmt.Println("Works length idexes ", len(works))
	workIndex := workID - 1
	fmt.Println("work idx: ", workIndex)
	(*works_map)["works"] = append(works[:workIndex], works[workIndex + 1:]...)
	for i := range (*works_map)["works"] {
		(*works_map)["works"][i].WorkID = i + 1
	}

	fmt.Println("My mapp: ", *works_map)
	return nil
}


// Create a variable to hold the result
// var result interface{}

// func InsertWork(supa *supabase.Client) {
// 	newWork := struct {
// 		Title   string `json:"title"`
// 		PicPath string `json:"picPath"`
// 	}{
// 		Title:   "New Title",
// 		PicPath: "/path/to/picture",
// 	}

// 	// Build the query
// 	query := supa.DB.From("works").Insert(newWork)
// 	err := query.Execute(&result)
// 	if err != nil {
// 		log.Printf("Error inserting work: %v", err)
// 		return
// 	}

// 	// Handle the successful insertion
// 	log.Println("Inserted work:", result)
// }

/*
	Get all "works" from database.
	Returns a slice []work that contains all the works.
	Example: works[0].Title or .WorkID or .Path(url)
*/
func AllWorks() []Work{
	
	var worksQuery []WorkQuery
	var works []Work
	// prepare godot to rerad from .env
	if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	dbURL := os.Getenv("DB_URL")
	dbKey:= os.Getenv("DB_KEY")

	supaClient, err := supabase.NewClient(dbURL, dbKey, nil) // Maybe need to init this on another place...
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	//  Query "works" table
	result, _, err := supaClient.From("works").
		Select("*", "", false).
		Order("WorkID", &postgrest.OrderOpts{Ascending: true}).
		Execute() // true for descending order.Execute()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
    if err := json.Unmarshal(result, &worksQuery); err != nil {
        log.Fatalf("Error unmarshalling result: %v", err)
    }

	// Open storage(bucket) client and prepare 'works = []work' to send data to front
    storage := supaClient.Storage
	for i,_:= range worksQuery{
    	filePath := worksQuery[i].Path
    	title := worksQuery[i].Title
		workID:= worksQuery[i].WorkID
		publicURL := storage.GetPublicUrl("works", filePath)
		works = append(works, Work{Path: publicURL.SignedURL, Title: title, WorkID: workID})
}
	return works
}