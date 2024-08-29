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
	"net/url"
	"os"
	"strconv"
	"strings"

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
	var picNames []string
	picNames = append(picNames, picName)
	storage := supaClient.Storage
	response, err:= storage.RemoveFile("works",picNames)
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
	
	type workInsert struct{
		Path   string `json:"Path"`
		Title     string `json:"Title"`
		Description string `json:"Description"`
		Position int    `json:"Position"`
	}
	
	var works []models.Work
	var insertedWorks []models.Work
	var insertedWorkID int
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

	newInsertWork := workInsert{
		Title: newTitle,
		Path: picName,
		Position: positionToInsert,
		Description: description,
	}
	// Check first if Work is going be inserted in last position... else
	if positionToInsert == int(totalWorks) + 1{
		insertedQuery, _, err := supaClient.From("works").Insert(newInsertWork, true, "","", "").Execute() // Insert new Work
		if err!=nil{
			return err
		}
		json.Unmarshal(insertedQuery, &insertedWorks)
		insertedWorkID = insertedWorks[0].Id
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
				_,_,err = supaClient.From("works").
				Update(map[string]interface{}{"Position": work.Position,}, "", "").
				Eq("Path",work.Path).Execute()
				if err!=nil{
					return err
				}
			}
		}
		// Insert new work /// check response here....................................................
		insertedQuery, _, err := supaClient.From("works").Insert(newInsertWork, true, "","", "").Execute()
		if err!=nil{
			return err
		}
		json.Unmarshal(insertedQuery, &insertedWorks)
		insertedWorkID = insertedWorks[0].Id
	}

	// Insert picture to bucket as also create sub folder in galleries bucket
	content:="image/jpeg"
	fileOption := storage_go.FileOptions{
		ContentType: &content,
	}
	newKey := insertedWorkID
	models.GalleriesStorage[newKey]= []models.GalleryItem{} // Add to local storage map gallery, new key for the work
	picReader := bytes.NewReader(picBytes)
	_, _ = supaClient.Storage.UploadFile("works",newInsertWork.Path, picReader, fileOption)
	folderPath := fmt.Sprintf("%d/",insertedWorkID)
	response, _ := supaClient.Storage.UploadFile("galleries",folderPath, picReader, fileOption)
	fmt.Println("Response", response.Message)

	// Update local storage
	models.WorksStorage = AllWorks()
	return nil
}

/*
	Edit tile of a work.
*/ 
func EditTitle(position string, newTitle string, newDescription string, newPicName string) (bool, error) {
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
	if newPicName == ""{
	_, _, err = supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription}, "", "").Eq("Position", position).Execute()
    if err != nil {
        return false, fmt.Errorf("error updating record: %w", err)
    }}else{
		fmt.Println("Work position to delete pic: ", position)
		workToDeleteQuery,_,_ := supaClient.From("works").Select("Path,Title,Description,Position","",false).Eq("Position", position).Execute()
		json.Unmarshal(workToDeleteQuery, &workTodeletePic)
		picNameToDelete := workTodeletePic[0].Path
		_, _, err = supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription, "Path":newPicName}, "", "").Eq("Position", position).Execute()
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
func DeleteWork(positionToDel string) error {
    supaClient := InitDB()

	var updatedWorkList []models.Work
	var err error = nil
	var worksToDelete []models.Work
	
    fmt.Print("\nDELETE ACTIVATED\n\n")
	
	positionToDelete, err:= strconv.Atoi(positionToDel)
	if err!= nil{
		return err
	}

	// Fecth work that will be deleted from db
	workToDeleteQuery, _, err := supaClient.From("works").Select("*","", false).Eq("Position", positionToDel).Execute()
	if err!=nil{
		return err
	}
	err =json.Unmarshal(workToDeleteQuery,&worksToDelete)
	if err!= nil{
		fmt.Printf("Error unmarshalling result: %v", err)
		return err
	}	

	workToDelete := worksToDelete[0] // [0] will always be the the work to delete, only one result from query.
	workIdKey := workToDelete.Id
	workPathToDelete := workToDelete.Path

	// Updating local storage
	var newWorkStorage []models.Work
	for i:= range models.WorksStorage{
		if models.WorksStorage[i].Position == positionToDelete{
			continue // skip, this will be the deleted work
		}
		if models.WorksStorage[i].Position > positionToDelete {
			models.WorksStorage[i].Position -= 1 // shift down by 1, positions after the delete
		}
		newWorkStorage = append(newWorkStorage, models.WorksStorage[i])
	}
	models.WorksStorage = newWorkStorage

	// Delete work from database
	fmt.Printf("\n\nProceeding to delete from database:\n")
	_, _, err = supaClient.From("works").Delete("","").Eq("Position", positionToDel).Execute()
	if err!=nil{
		return err
	}

	// Update works on database
	fmt.Printf("\nUpdating works on database %v\n\n", updatedWorkList)
	for _, work := range models.WorksStorage{
		updatedWorkToInsert := work
		args:= strings.Split(updatedWorkToInsert.Path, "/")
		picNameEncodedURL := args[len(args)-1] // accesing file name only since local .path contains the url for the pic
		updatedWorkToInsert.Path, _= url.QueryUnescape(picNameEncodedURL) // picture name can contain white spaces from the url generated
		workIdstring := strconv.Itoa(work.Id)
		fmt.Println("WORK ID: ", work.Position)
	_,_,err = supaClient.From("works").Update(updatedWorkToInsert, "", "").Eq("ID",workIdstring).Execute()
	if err!=nil{
		return err
	}}

	// Delete picture associated to work
	fmt.Printf("Picture to delete: %s",workPathToDelete)
	err = DeletePicture(workPathToDelete)
	if err!= nil{
		fmt.Println("Error trying to delete picture in bucket: ", err)
		return err
	}

	// Update Galleries locally
	delete(models.GalleriesStorage, workIdKey)

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
	worksQuery, _, err := supaClient.From("works").
		Select("Position,Title,Path,Description,ID", "", false).
		Order("Position", &postgrest.OrderOpts{Ascending: true}).
		Execute() // true for descending order.Execute()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
    if err := json.Unmarshal(worksQuery, &works); err != nil {
        log.Fatalf("Error unmarshalling query result: %v", err)
    }

	// Open bucket to get picture urls
    storageBucket := supaClient.Storage
	for i:= range works{
		publicURL := storageBucket.GetPublicUrl("works", works[i].Path)
		works[i].Path = publicURL.SignedURL
}

return works
}
func AllGalleries() []models.GalleryItem{
	
	var galleryItems []models.GalleryItem
	supaClient:= InitDB()
	
	//  Query "galleries" table
	galleriesQuery, _, err := supaClient.From("galleries").
	Select("*", "", false).
	Order("Work_ID", &postgrest.OrderOpts{Ascending: true}).
	Execute()
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
    if err := json.Unmarshal(galleriesQuery, &galleryItems); err != nil {
		log.Fatalf("Error unmarshalling query result: %v", err)
    }
	
	storageBucket := supaClient.Storage
	for i := range galleryItems{
		subFolderName := strconv.Itoa(galleryItems[i].Work_ID)
		filePath := fmt.Sprintf("%s/%s", subFolderName, galleryItems[i].Path)
		publicUrl:= storageBucket.GetPublicUrl("galleries",filePath)
		galleryItems[i].Path = publicUrl.SignedURL
	}

	return galleryItems
}
	// Check if map exists
	// if models.GalleriesStorage == nil {
	// 	models.GalleriesStorage = make(map[int][]models.GalleryItem)
	// }
	// for _, galleryItem := range galleries{
	// 		fmt.Println("Gallery Item: ", galleryItem)
	// 		fmt.Println("Appending item to local storage")
	// 		key:= galleryItem.Work_ID

	// 		_, exists := models.GalleriesStorage[key] // Check if key doesnt exist on the local storage map
	// 		if !exists{
	// 			models.GalleriesStorage[key] = []models.GalleryItem{}
	// 		}
	// 		models.GalleriesStorage[key] = append(models.GalleriesStorage[key], galleryItem) // Append Item to the []GalleryItem
	// 	}

	// for keyID, gallery := range models.GalleriesStorage{
	// 	fmt.Println("--------------------------------------------------")
	// 	fmt.Println("Work ID KEY = ", keyID)
	// 	for _, item := range gallery{
	// 		fmt.Println("Pic Name: ", item.Path)
	// 		fmt.Println("Pic Position: ", item.Position)
	// 	}
	// }