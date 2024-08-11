package galleries

import (
	"Go_servers/db"
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//WOORKKK HEREEEEEEEEEEEE __________________________________________________________
/*
	Handler to fecth the gallery of a work based on the title.
*/
func Gallery(w http.ResponseWriter, r *http.Request, templateFs embed.FS) {
	fmt.Println("Gallery activated")
	title := r.PathValue("title")

	fmt.Println("My title is: ", title)
	supaClient:= db.InitDB()
	workUniqueID,results,err:= supaClient.From("works").Select("ID", "exact", false).Eq("Title", title).Execute()
	if err!=nil{
		fmt.Println("Error ", err)
		http.Error(w, "Error in query", http.StatusInternalServerError)
		return
	}
	if results == 0{
		http.Error(w, "No results", http.StatusNotFound)
		return
	}
	fmt.Println("Work unique ID: ", workUniqueID)
	// Fecth from the galleries table here:
	// galleryPicsPaths, _, err:= supaClient.From("galleries").Select("Path","",false).Filter("workUniqueID","=",string(workUniqueID)).Execute()

	tmpl := template.Must(template.ParseFS(templateFs,"htmlTemplates/gallery.html"))
	tmpl.Execute(w, nil)
}
