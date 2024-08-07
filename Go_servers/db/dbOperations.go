/*
Database handler
*/
package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

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

/*
	Initialize client for supabase.
*/ 
func InitDB() *supabase.Client {
	fmt.Println("Initializing DATABASE:")
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbURL := os.Getenv("DB_URL")
	dbKey := os.Getenv("DB_KEY")

	// Initialize Supabase client
	supaClient, err := supabase.NewClient(dbURL, dbKey, nil)
	if err != nil {
		log.Fatalf("Error creating Supabase client: %v", err)
	}

	return supaClient
}


func InsertWork(title string, workID string, picName string) error {
	supaClient:= InitDB()
	var newWork Work
	// Count the number of rows in the "works" table before insertion
	_, totalWorks, err := supaClient.From("works").Select("*", "exact", false).Execute()
	if err != nil {
		return err
	}
	// Print the count result
	fmt.Printf("Count Result: %v\n", totalWorks)
	workIdToInsert, err:= strconv.Atoi(workID)
	if err!= nil{
		return err
	}
	// Work is going be inserted in last position
	if workIdToInsert == int(totalWorks) + 1{
		newWork = Work{
			Title: title,
			Path: picName,
			WorkID: workIdToInsert,
		}
		_, _, err = supaClient.From("works").Insert(newWork, true, "","", "").Execute()
		if err!=nil{
			return err
		}
	}
	return nil
}


/*
	Edit tile of a work.
*/ 
func EditTitle(workID string, newTitle string) (bool, error) {
    supaClient := InitDB()
	fmt.Println("EDITING WITH DB")
    // Perform the update directly
    _, _, err := supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,}, "", "").Eq("WorkID", workID).Execute()
    if err != nil {
        return false, fmt.Errorf("error updating record: %w", err)
    }

    return true, nil
}

/*
	Delete a work
*/ 
func DeleteWork(workID string) error {
    supaClient := InitDB()
    fmt.Println("DELETING FROM DB")
	var works []WorkQuery
	var updatedWorkList []WorkQuery
	var err error = nil

	// Count the number of rows in the "works" table before deletion
	worksRows, totalWorks, err := supaClient.From("works").Select("*", "exact", false).Execute()
	if err != nil {
		return err
	}
	// Print the count result
	fmt.Printf("Count Result: %v\n", totalWorks)
	workIdToDelete, err:= strconv.Atoi(workID)
	if err!= nil{
		return err
	}
	// Deleting last work, no need to change workId of other works... else
	if workIdToDelete == int(totalWorks){
		_, _, err = supaClient.From("works").Delete("","").Eq("WorkID", workID).Execute()
		if err!=nil{
			return err
		}
	}else{
		if err := json.Unmarshal(worksRows, &works); err != nil {
			log.Fatalf("Error unmarshalling result: %v", err)
			return err
		}
		// Update works with new workIDs
		for _, work := range works{
			// assing 0 to the work we will delete later
			if work.WorkID == workIdToDelete{
				work.WorkID = 0
			}
			// Reduce workId by one for all works after ID selected is deleted.
			if work.WorkID > workIdToDelete{
				work.WorkID = work.WorkID - 1
			}
			updatedWorkList = append(updatedWorkList, work)
		}
		// Update works
		fmt.Printf("updated works %v\n", updatedWorkList)
		for _, work := range updatedWorkList{
			fmt.Println("WORK ID: ", work.WorkID)
		_,_,err = supaClient.From("works").Update(work, "", "").Eq("ID", strconv.Itoa(work.ID)).Execute()
		if err!=nil{
			return err
		}}
		// Delete specified work
		fmt.Println("Works udpated, proceeding to delete")
		_, _, err = supaClient.From("works").Delete("","").Eq("WorkID", "0").Execute()
		if err!=nil{
			return err
		}
	}
	return err
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
	Example: works[0].Title or .WorkID or .Path
*/
func AllWorks() []Work{
	
	var (
		worksQuery []WorkQuery
		works []Work
	)
	supaClient:= InitDB()

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

	// Open storage(bucket) client to get picture url
    storage := supaClient.Storage
	for i:= range worksQuery{
    	filePath := worksQuery[i].Path
    	title := worksQuery[i].Title
		workID:= worksQuery[i].WorkID
		publicURL := storage.GetPublicUrl("works", filePath)
		works = append(works, Work{Path: publicURL.SignedURL, Title: title, WorkID: workID})
}
	return works
}