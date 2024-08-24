/*
Database handler
*/
package db

import (
	"Go_servers/models"
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
// type Work struct {
// 	CreatedAt string `json:"CreatedAt"`
// 	ID        int    `json:"ID"`
// 	Path   string `json:"Path"`
// 	Title     string `json:"Title"`
// 	Position int    `json:"Position"`
// }
// Use this to send works to html
type WorkForFront struct {
    Path   string `json:"Path"`
    Title     string `json:"Title"`
    Position int    `json:"Position"`
}

/*
	Delete picture in 'works' bucket when the work will be deleted.
*/
func DeletePicture(picName string) error{
	supaClient:= InitDB()
	var paths []string
	paths = append(paths, picName)
	storage := supaClient.Storage
	response, err:= storage.RemoveFile("works",paths)
	if err!=nil{
		fmt.Printf("Error deleting picture: %v\n", err)
		fmt.Println("response", response)
		return err
	}
	return nil
}

func AddPicture(picName string, picBytes []byte) error{
	supaClient:= InitDB()
	// Insert picture to bucket as jpeg content headers 
	content:="image/jpeg"
	fileOption := storage_go.FileOptions{
		ContentType: &content,
	}

	picReader := bytes.NewReader(picBytes)
	res, err := supaClient.Storage.UploadFile("works",picName, picReader, fileOption)
	fmt.Println("Response", res)
	fmt.Println("Error: ",err)
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


func Login(email string, pwd string) (string,error){
	supaClient:= InitDB()
	fmt.Println("Email: ",email)
	fmt.Println("Password: ",pwd)
	authSession , err:= supaClient.SignInWithEmailPassword(email, pwd)
	if err!=nil{
		fmt.Println("Error", err)
		return "", err
	}
	token:= authSession.AccessToken
	return token, nil
}

/*
	Insert new work. Picture is also sent here. 
*/
func InsertWork(newTitle string, position string, description string,picName string, picBytes []byte) error {
	var newWork models.Work
	var works []models.Work
	supaClient:= InitDB()

	// Count the number of rows in the "works" table before insertion
	workRowsQuery, totalWorks, err := supaClient.From("works").Select("Position,Title,Path,Description", "exact", false).Execute()
	if err != nil {
		return err
	}
	positionToInsert, err:= strconv.Atoi(position) // int version of work id
	if err!= nil{
		return err
	}

	//New work to insert
	newWork = models.Work{
		Title: newTitle,
		Path: picName,
		Position: positionToInsert,
		Description: description,
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
				_,_,err = supaClient.From("works").Update(work, "", "").Eq("Path",work.Path).Execute()
				if err!=nil{
					return err
				}
			}
		}
		// Insert new work
		_, _, err = supaClient.From("works").Insert(newWork, true, "","", "").Execute()
		if err!=nil{
			return err
		}
		
	}
	// Insert picture to bucket as jpeg content headers 
	content:="image/jpeg"
	fileOption := storage_go.FileOptions{
		ContentType: &content,
	}

	picReader := bytes.NewReader(picBytes)
	response, err:= supaClient.Storage.UploadFile("works",newWork.Path, picReader, fileOption)
	fmt.Println("Response", response)
	fmt.Println("Error: ",err)

	// Update local storage
	models.WorksStorage = AllWorks()
	return nil
}

/*
	Edit tile of a work.
*/ 
func EditTitle(position string, newTitle string, newDescription string, newPath string) (bool, error) {
    supaClient := InitDB()
	var workTodeletePic []models.Work
	positionInt, err := strconv.Atoi(position)
	if err!=nil{
		return false, fmt.Errorf("error in ATOI: %w", err)
	}
	fmt.Println("Description: ", newDescription)
	fmt.Println("EDITING WITH DB")
    // Perform the update on databse
    // Option 1(if), picture was not changed, option 2(else) picture was changed
	if newPath == ""{
	_, _, err = supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription}, "", "").Eq("Position", position).Execute()
    if err != nil {
        return false, fmt.Errorf("error updating record: %w", err)
    }}else{
		fmt.Println("Work position to delete pic: ", position)
		workToDeleteQuery,_,_ := supaClient.From("works").Select("Path,Title,Description,Position","",false).Eq("Position", position).Execute()
		json.Unmarshal(workToDeleteQuery, &workTodeletePic)
		picNameToDelete := workTodeletePic[0].Path
		_, _, err = supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription, "Path":newPath}, "", "").Eq("Position", position).Execute()
		if err != nil {
    return false, fmt.Errorf("error updating record: %w", err)
	}
	fmt.Println("Pic to delete: ", picNameToDelete)
	DeletePicture(picNameToDelete)
	}
	// Update Local Storage
	models.WorksStorage[positionInt - 1].Title = newTitle
	models.WorksStorage[positionInt - 1].Description = newDescription
    return true, nil
}

/*
	Delete a work
*/ 
func DeleteWork(position string) error {
    supaClient := InitDB()
	var works []models.Work
	var updatedWorkList []models.Work
	var err error = nil
	var worksToDelete []models.Work
	
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
		_,_,err = supaClient.From("works").Update(workToDelete, "", "").Eq("Path", workToDelete.Path).Execute()
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
		_,_,err = supaClient.From("works").Update(work, "", "").Eq("Path",work.Path).Execute()
		if err!=nil{
			return err
		}}
	}
	// Finally Delete specified work with the position = 0 and its picture from the bucket
	err = DeletePicture(workToDelete.Path)
	if err!= nil{
		fmt.Println("Error trying to delete picture in bucket: ", err)
		return err
	}

	fmt.Println("Works udpated, proceeding to delete")
	_, _, err = supaClient.From("works").Delete("","").Eq("Position", "0").Execute()
	if err!=nil{
		return err
	}

	models.WorksStorage = AllWorks()
	return err
}

/*
	Get all "works" from database.
	Returns a slice []work that contains all the works.
	Example: works[0].Title or .Position or .Path
*/
func AllWorks() []models.Work{
	
	var works []models.Work
	supaClient:= InitDB()

	//  Query "works" table
	result, _, err := supaClient.From("works").
		Select("Position,Title,Path,Description", "", false).
		Order("Position", &postgrest.OrderOpts{Ascending: true}).
		Execute() // true for descending order.Execute()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
    if err := json.Unmarshal(result, &works); err != nil {
        log.Fatalf("Error unmarshalling result: %v", err)
    }

	// Open bucket to get picture urls
    storage := supaClient.Storage
	for i:= range works{
		publicURL := storage.GetPublicUrl("works", works[i].Path)
		works[i].Path = publicURL.SignedURL
}
	return works
}
