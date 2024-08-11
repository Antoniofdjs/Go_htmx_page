/*
Database handler
*/
package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/supabase-community/postgrest-go"
	storage_go "github.com/supabase-community/storage-go"
	"github.com/supabase-community/supabase-go"
)

// Use this for query on "works" table
type Work struct {
	CreatedAt string `json:"CreatedAt"`
	ID        int    `json:"ID"`
	Path   string `json:"Path"`
	Title     string `json:"Title"`
	Position int    `json:"Position"`
}
// Use this to send works to html
type WorkForFront struct {
    Path   string `json:"Path"`
    Title     string `json:"Title"`
    Position int    `json:"Position"`
}

/*
	Work method used to delete the picture in the bucket of the work.
*/
func (w Work) deletePicture() error{
	supaClient:= InitDB()
	var paths []string
	paths = append(paths, w.Path)
	storage := supaClient.Storage
	response, err:= storage.RemoveFile("works",paths)
	if err!=nil{
		fmt.Printf("Error deleting picture: %v\n", err)
		fmt.Println("response", response)
		return err
	}
	return nil
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

/*
	Insert new work. Picture is also sent here. 
*/
func InsertWork(newTitle string, position string, picName string, picBytes []byte) error {
	var newWork WorkForFront
	var works []Work
	supaClient:= InitDB()

	// Count the number of rows in the "works" table before insertion
	workRowsQuery, totalWorks, err := supaClient.From("works").Select("*", "exact", false).Execute()
	if err != nil {
		return err
	}
	positionToInsert, err:= strconv.Atoi(position) // int version of work id
	if err!= nil{
		return err
	}

	//New work to insert
	newWork = WorkForFront{
		Title: newTitle,
		Path: picName,
		Position: positionToInsert,
	}
	// Check first if Work is going be inserted in last position... else
	if positionToInsert == int(totalWorks) + 1{
		_, _, err = supaClient.From("works").Insert(newWork, true, "","", "").Execute()
		if err!=nil{
			return err
		}
	}else{
		err:= json.Unmarshal(workRowsQuery, &works)
		if err!=nil{
			fmt.Printf("Error unmarshalling result: %v\n", err)
			return err
		}

		// Increase by one any work id matching or higher than the new work coming
		//This will shift all work id's correctly and update them 
		// COnsider a faster way to do this, maybe a function in the db
		for _, work := range works{
			if work.Position >= positionToInsert{
				work.Position += 1
				_,_,err = supaClient.From("works").Update(work, "", "").Eq("ID", strconv.Itoa(work.ID)).Execute()
				if err!=nil{
					return err
				}
			}
		}
		// Finally insert new work
		_, _, err = supaClient.From("works").Insert(newWork, true, "","", "").Execute()
		if err!=nil{
			return err
		}
		
	}
	// Finally insert picture to bucket with jpeg content 
	content:="image/jpeg"
	fileOption := storage_go.FileOptions{
		ContentType: &content,
	}

	picReader := bytes.NewReader(picBytes)
	response, err:= supaClient.Storage.UploadFile("works",newWork.Path, picReader, fileOption)
	fmt.Println("Response", response)
	fmt.Println("Error: ",err)
	return nil
}

/*
	Edit tile of a work.
*/ 
func EditTitle(position string, newTitle string) (bool, error) {
    supaClient := InitDB()
	fmt.Println("EDITING WITH DB")
    // Perform the update directly
    _, _, err := supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,}, "", "").Eq("Position", position).Execute()
    if err != nil {
        return false, fmt.Errorf("error updating record: %w", err)
    }

    return true, nil
}

/*
	Delete a work
*/ 
func DeleteWork(position string) error {
    supaClient := InitDB()
	var works []Work
	var updatedWorkList []Work
	var err error = nil
	var worksToDelete []Work
	
    fmt.Println("DELETING FROM DB")
	
	// Count the number of rows in the "works" table before deletion
	workRowsQuery, totalWorks, err := supaClient.From("works").Select("*", "exact", false).Execute()
	if err != nil {
		return err
	}

	// Print the count result
	fmt.Printf("Count Result: %v\n", totalWorks)
	positionToDelete, err:= strconv.Atoi(position)
	if err!= nil{
		return err
	}
	// Fecth work to delete
	workQuery, _, err := supaClient.From("works").Select("*","", false).Eq("Position", position).Execute()
	if err!=nil{
		return err
	}
	err =json.Unmarshal(workQuery,&worksToDelete)
	if err!= nil{
		fmt.Printf("Error unmarshalling result: %v", err)
		return err
	}
	workToDelete := worksToDelete[0]

	// Delete last work, no need to change position of other works... else
	if positionToDelete == int(totalWorks){
		fmt.Println("Last work being deleted")
		workToDelete.Position = 0
		_,_,err = supaClient.From("works").Update(workToDelete, "", "").Eq("ID", strconv.Itoa(workToDelete.ID)).Execute()
		if err!=nil{
			return err
		}
	}else{
		if err := json.Unmarshal(workRowsQuery, &works); err != nil{
			log.Fatalf("Error unmarshalling result: %v", err)
			return err
		}

		// Update works with new Positions
		for _, work := range works{
			if work.Position == positionToDelete{
				work.Position = 0 // assing 0 to the work we will delete later
			}
			// Reduce 'Position' by one for all works after ID selected for delete
			if work.Position > positionToDelete{
				work.Position = work.Position - 1
			}
			updatedWorkList = append(updatedWorkList, work)
		}
		// Update works
		fmt.Printf("updated works %v\n", updatedWorkList)
		for _, work := range updatedWorkList{
			fmt.Println("WORK ID: ", work.Position)
		_,_,err = supaClient.From("works").Update(work, "", "").Eq("ID", strconv.Itoa(work.ID)).Execute()
		if err!=nil{
			return err
		}}
	}
	// Finally Delete specified work with the position = 0 and its picture from the bucket
	err = workToDelete.deletePicture()
	if err!= nil{
		fmt.Println("Error trying to delete picture in bucket: ", err)
		return err
	}
	fmt.Println("Works udpated, proceeding to delete")
	_, _, err = supaClient.From("works").Delete("","").Eq("Position", "0").Execute()
	if err!=nil{
		return err
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
	Example: works[0].Title or .Position or .Path
*/
func AllWorks() []WorkForFront{
	
	var (
		worksQuery []Work
		works []WorkForFront
	)
	supaClient:= InitDB()

	//  Query "works" table
	result, _, err := supaClient.From("works").
		Select("*", "", false).
		Order("Position", &postgrest.OrderOpts{Ascending: true}).
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
		position:= worksQuery[i].Position
		publicURL := storage.GetPublicUrl("works", filePath)
		works = append(works, WorkForFront{Path: publicURL.SignedURL, Title: title, Position: position})
}
	return works
}