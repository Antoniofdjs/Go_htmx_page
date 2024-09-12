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
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

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
func DeletePicture(picNames []string, bucketName string) error{
	supaClient:= InitDB()
	storage := supaClient.Storage
	response, err:= storage.RemoveFile(bucketName,picNames)
	if err!=nil{
		fmt.Printf("Error deleting picture: %v\n", err)
		fmt.Println("response", response)
		return err
	}
	fmt.Println("Deleted pic ",picNames[0])
	return nil
}
/*
	Delete a subfolder in bucket
*/
func DeleteSubFolder(bucketName string, WorkID int) error{
	subFolderPath:= fmt.Sprintf("%d/", WorkID)
	supaClient:= InitDB()
	storage := supaClient.Storage
	options := storage_go.FileSearchOptions{}
	
	files, err:= storage.ListFiles(bucketName,subFolderPath, options)
	if err!=nil{
		fmt.Printf("Error deleting picture: %v\n", err)
		fmt.Println("response", files)
		return err
	}
	var fileNames []string

	// Extract all file names from sub folder
	for _, file:= range files{
		fmt.Println("File name pic to delete: ",file.Name)
		pathName:= fmt.Sprintf("%d/%s", WorkID, file.Name)
		fileNames = append(fileNames, pathName)
	}

	// If there are no files, return
	if len(fileNames) == 0 {
		fmt.Printf("No files found in subfolder: %d\n", WorkID)
		return nil
	}

	// Remove all files in the subfolder
	fmt.Println("\n\nRemoving all files from sub folder:")
	response, err := storage.RemoveFile(bucketName, fileNames)
	if err != nil {
		fmt.Printf("Error deleting files in subfolder: %v\n", err)
		fmt.Println("response", response)
		return err
	}

	fmt.Println("Deleted Sub Folder")
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


/* Insert new gallery item to table "galleries" */ 
func InsertGalleryItem(workID int, fileName string, fileBytes []byte) error{
	fmt.Println("INSERT ITEMS GALLERY TO DB ACTIVATED")
	supaClient := InitDB()

	positionForItem:= len(models.GalleriesStorage[workID]) + 1 // count how many items are already in the map for this work on local storage
	for _, item := range models.GalleriesStorage[workID]{
		fmt.Print("\nOld items: position- ", item.Position)
	}
	newGalleryItem := models.GalleryItem{ 
		Work_ID: workID,
		Path: fileName,
		Position: positionForItem,
		} 			// this newGalleryItem will be placed on the local storage map as a placeholder so that each next item can increase the count,
					//after all inserts are done, storageInits.InitGalleries will correctly fix the data
	
	type galleryInsert struct{
		Path string `json:"Path"`
		Position int `json:"Position"`
		Work_ID int `json:"Work_ID"` // Reference key to table "works"
	}
	newGalleryItemInsert := galleryInsert{
		Path: fileName,
		Position: positionForItem,
		Work_ID: workID,

	}
	
	fmt.Println("Inserting gallery item")
	response, _, err:= supaClient.From("galleries").Insert(newGalleryItemInsert, true, "", "","").Execute()
	if err!= nil{
		fmt.Println("Error inserting new gallery item")
		fmt.Printf("%s", err)
		return err
	}
	fmt.Println("Position of item : ",newGalleryItemInsert.Position)
	fmt.Println("\n\nInsert succes")
	fmt.Println("Response", response)

	// Insert picture to bucket as also create sub folder in galleries bucket
	content:="image/jpeg"
	fileOption := storage_go.FileOptions{
		ContentType: &content,
	}
	picReader := bytes.NewReader(fileBytes)
	folderPath := fmt.Sprintf("%d/%s",workID, fileName)
	storageBucket:= supaClient.Storage
	responseBucket , err:= storageBucket.UploadFile("galleries", folderPath,picReader, fileOption)
	if err!=nil{
		fmt.Println("Error inserting picture to bucket")
	}
	fmt.Println("Response Bucket", responseBucket.Message)
	models.GalleriesStorage[workID] = append(models.GalleriesStorage[workID], newGalleryItem) // newGalleryItem.path its not correct format, after this sotrageInits.InitGalleries() will be called to fix
	return err
}

func DeleteGalleryItem(workID string, position string, picName string) error{
	fmt.Println("DELETE ITEM GALLERY IN DB ACTIVATED")
	supaClient := InitDB()
	galleryValues:= map[string]string{
		"Work_ID": workID,
		"Position": position,
	}
	_,_, err := supaClient.From("galleries").Delete("", "").Match(galleryValues).Execute()
	if err!=nil{
		fmt.Println("Error deleting gallery item: ", err)
		return err
	}

	picPath := fmt.Sprintf("%s/%s", workID, picName)
	fmt.Println("Going to delete pciture: ",picPath)
	picPaths :=[]string{
		picPath,
	}
	err = DeletePicture(picPaths, "galleries")
	if err != nil{
		fmt.Println("Error deleting gallery pic from bucket")
		return err
	}
	
	return err
}


/*
	Insert new work. Picture is also sent here. 
*/
func InsertWork(newTitle string, position string, description string, picName string, picBytes []byte) error {
	
	type workInsert struct{
		Path   string `json:"Path"`
		Title     string `json:"Title"`
		Description string `json:"Description"`
		Position int    `json:"Position"`
	}
	
	var works []models.Work
	var insertedWorks []models.Work
	var insertedWorkID int

	// check if pic name doesnt exists
	for key := range models.WorksMapStorage{
		args:= strings.Split(models.WorksMapStorage[key].Path, "/")
		picNameEncodedURL := args[len(args)-1] // accesing file name only
		existingPicNAme, _:= url.QueryUnescape(picNameEncodedURL) // picture name can contain white spaces from the url generated.
		if existingPicNAme == picName{
			
			// Seed the random number generator to have unique nums(controlled randomness)
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    
			// Generate a random number
			randomNum := rng.Intn(100)
			picName = fmt.Sprintf("Cpy_%d_%s", randomNum, picName) // modify name of the pic since it already exists
		}
	}
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
	models.GalleriesStorage[newKey]= []models.GalleryItem{} // Add to local storage gallery, new key
	picReader := bytes.NewReader(picBytes)
	_, _ = supaClient.Storage.UploadFile("works",newInsertWork.Path, picReader, fileOption)
	folderPath := fmt.Sprintf("%d/",insertedWorkID)
	response, _ := supaClient.Storage.UploadFile("galleries",folderPath, picReader, fileOption)
	fmt.Println("Response", response.Message)

	return nil
}

/*
	Edit tile of a work.
*/ 
func EditWork(position string, newTitle string, newDescription string, newPicName string) (bool, error) {
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
	_, _, err := supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription}, "", "").Eq("Position", position).Execute()
    if err != nil {
        return false, fmt.Errorf("error updating record: %w", err)
    }}else{
		fmt.Println("Work position to delete pic: ", position)
		workToDeleteQuery,_,_ := supaClient.From("works").Select("Path,Title,Description,Position","",false).Eq("Position", position).Execute()
		json.Unmarshal(workToDeleteQuery, &workTodeletePic)
		picNameToDelete := workTodeletePic[0].Path
		_, _, err := supaClient.From("works").Update(map[string]interface{}{"Title": newTitle,"Description":newDescription, "Path":newPicName}, "", "").Eq("Position", position).Execute()
		if err != nil {
    return false, fmt.Errorf("error updating record: %w", err)
	}
	
	fmt.Println("Pic to delete: ", picNameToDelete)
	picPaths:= []string{
		picNameToDelete,
	}
	DeletePicture(picPaths,"works")
	}
	// Update Local Storage
	models.WorksStorage[positionInt - 1].Title = newTitle
	models.WorksStorage[positionInt - 1].Description = newDescription
    return true, nil
}

/*
	Delete a work from "works" table
*/ 
func DeleteWork(positionToDel string) error {
    supaClient := InitDB()

	var updatedWorkList []models.Work
	var err error = nil
	var worksToDelete []models.Work
	var workTitleToDelete string
	
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
			workTitleToDelete = models.WorksStorage[i].Title
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
		updatedWorkToInsert.Path, _= url.QueryUnescape(picNameEncodedURL) // picture name can contain white spaces from the url generated.
		workIdstring := strconv.Itoa(work.Id)
		fmt.Println("WORK ID: ", work.Position)
	_,_,err = supaClient.From("works").Update(updatedWorkToInsert, "", "").Eq("ID",workIdstring).Execute()
	if err!=nil{
		return err
	}}

	// Delete picture associated to work
	fmt.Printf("Picture to delete: %s",workPathToDelete)
	picPaths:= []string{
		workPathToDelete,
	}
	err = DeletePicture(picPaths, "works")
	if err!= nil{
		fmt.Println("Error trying to delete picture in bucket: ", err)
		return err
	}

	// Update Galleries locally and delete bucket pictures("works" table has delete cascade, will delete the gallery items on the table)
	// var galleryPaths = []string{}
	// galleryItems := models.GalleriesStorage[workIdKey]
	// for _, item := range galleryItems{
	// 	args:= strings.Split(item.Path, "/")
	// 	picName,_:= url.QueryUnescape(args[(len(args) -1 )]) // get name of pic translate special chars or white spaces
	// 	picPath:= fmt.Sprintf("%d/%s", workIdKey, picName)
	// DeletePicture(galleryPaths, "galleries") // Remove pics from bucket
	err = DeleteSubFolder("galleries", workIdKey)
	if err!= nil{
		fmt.Println("Error deleting pic from bucket")
		return err
	}

	// Elimate work from map of works anad update the positions on map
	delete(models.WorksMapStorage, workTitleToDelete)
	for _, udpatedWork := range models.WorksStorage{
		models.WorksMapStorage[udpatedWork.Title] = udpatedWork
	}
	delete(models.GalleriesStorage, workIdKey) // Eliminate gallery from local storage

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
		fmt.Println("INIT PIC NAME: ", works[i].Path)
		works[i].Path = publicURL.SignedURL
		fmt.Println("INIT PIC URL: ", works[i].Path)

}

return works
}
/* Get all galleries from "galleries" table*/ 
func AllGalleries() []models.GalleryItem{
	
	var galleryItems []models.GalleryItem
	supaClient:= InitDB()
	
	//  Query "galleries" table
	galleriesQuery, _, err := supaClient.From("galleries").
	Select("*", "", false).
	Order("Position", &postgrest.OrderOpts{Ascending: true}).
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

func UpdateGalleryPositions(galleryItems []models.GalleryItem){
	supaClient := InitDB()
	for _, item := range galleryItems{
		itemID := strconv.Itoa(item.ID)
		updatedItem := map[string]interface{}{"Position":item.Position}
		supaClient.From("galleries").Update(updatedItem, "", "").Eq("id", itemID).Execute()
	}
}