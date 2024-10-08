package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"Go_htmx_page/db"
	"Go_htmx_page/handlers/about"
	contacts "Go_htmx_page/handlers/contact"
	editor "Go_htmx_page/handlers/editors"
	"Go_htmx_page/handlers/galleries"
	"Go_htmx_page/handlers/user"
	"Go_htmx_page/handlers/work"
	"Go_htmx_page/models"
	storageInits "Go_htmx_page/storageInit"

	"github.com/joho/godotenv"
)

// Embed all HTML files
//go:embed htmlTemplates
var templatesFS embed.FS

// Embed all files in the static directory
//// go:embed static/*
// var staticFS embed.FS

// Route wrappers
func initStorageMiddleware(nextHandler http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if models.GalleriesStorage == nil{
			storageInits.InitGalleries()
			models.DeleteGalleryItemTempStorage = make(map[string][]int)
		}
		if models.WorksMapStorage == nil{
			storageInits.InitWorksStorage()
		}
		nextHandler.ServeHTTP(w, r)
	}
}

func authMiddleware(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		supaClient:= db.InitDB()
		tokenName:= os.Getenv("SPB_TOKEN_NAME")
		spbToken, err:= r.Cookie(tokenName)
		if err != nil{
			msg := "Unauthorized"
			w.Write([]byte(msg))
            return
        }

		authClient:= supaClient.Auth.WithToken(spbToken.Value)

		_, err = authClient.GetUser()
		if err != nil {
			msg := "Unauthorized"
			w.Write([]byte(msg))
            return
        }

		//  Init local storage
		if models.GalleriesStorage == nil{
			storageInits.InitGalleries()
			models.DeleteGalleryItemTempStorage = make(map[string][]int)
		}
		if models.WorksMapStorage == nil{
			storageInits.InitWorksStorage()
		}
        
		nextHandler.ServeHTTP(w, r)
    }
}
// End Route Wrappers


func main() {
        
	err:= godotenv.Load()
	if err != nil {
        log.Println("Error loading .env file, using OS env")
		
    }

	fmt.Println("SERVER LISTENING:")
	// db.AllWorks() // Testing db storage and buucker from supabase here

	landingHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFS(templatesFS,"htmlTemplates/index.html"))
	tmpl.Execute(w, nil)
	}

	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("GET /{$}", landingHandler) // Match only exactly '/' thanks to {$}
	http.HandleFunc("GET /about",func(w http.ResponseWriter, r *http.Request){about.GetHand(w, r, templatesFS)})

	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {contacts.GetHand(w, r, templatesFS)})
	http.HandleFunc("POST /contact", func(w http.ResponseWriter, r *http.Request) {contacts.PostHand(w, r, templatesFS)})

	//  Work Routes
	http.HandleFunc("GET /work", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){work.GetWorksView(w, r, templatesFS)}))
	http.HandleFunc("GET /work/{title}", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){galleries.Gallery(w, r)})) // Gallery
	http.HandleFunc("GET /image", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){galleries.GetModal(w, r)})) // get modal for image

	//  Work Editor Routes
	http.HandleFunc("GET /editor", authMiddleware(func(w http.ResponseWriter, r *http.Request){editor.GetHandEditor(w, r, templatesFS)}))
	http.HandleFunc("PUT /editor", func(w http.ResponseWriter, r *http.Request){editor.PutHandEditor(w, r, templatesFS)}) // auth
	http.HandleFunc("POST /editor", func(w http.ResponseWriter, r *http.Request){editor.PostHandEditor(w, r, templatesFS)}) // auth
	http.HandleFunc("POST /editor/del", func(w http.ResponseWriter, r *http.Request){editor.DelHandEditor(w, r)}) // auth

	// LoginRoutes
	http.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request){user.GetLoginTmpl(w, r, templatesFS)})
	http.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request){user.Login(w, r)})
	http.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request){user.Logout(w, r)})

	http.HandleFunc("GET /editor/components", func(w http.ResponseWriter, r *http.Request){editor.GetEditorComponents(w, r, templatesFS)})

	// Gallery Editor Routes
	http.HandleFunc("GET /editor/{title}", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){editor.GetEditorGallery(w, r, templatesFS)}))
	http.HandleFunc("POST /editor/{title}", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){editor.PostHandGalleryEditor(w, r)})) // Actual insert of gallery items
	http.HandleFunc("PUT /editor/{title}", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){editor.PutHandGalleryEditor(w, r)})) // Edit gallery items(delete pics)
	
	// Thes 2 routes maybe need to be defined in another handler and group them
	http.HandleFunc("POST /editor/gallery", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){editor.FileUploadTemporaryStorage(w, r)}))
	http.HandleFunc("GET /editor/update", func(w http.ResponseWriter, r *http.Request){editor.UpdateGalleryItems(w, r)})

	// Test route for editor
	http.HandleFunc("GET /test", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){editor.GetTestView(w, r, templatesFS)})) // this is for /editor!!!!


	// Start server
	port := os.Getenv("PORT")
	if port == "" {
    port = "8080" // Default to port 8080 for local development
	}
	log.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}