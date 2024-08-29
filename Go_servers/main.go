package main

import (
	"Go_servers/db"
	"Go_servers/handlers/about"
	contacts "Go_servers/handlers/contact"
	"Go_servers/handlers/galleries"
	"Go_servers/handlers/user"
	"Go_servers/handlers/work"
	"Go_servers/models"
	storageInits "Go_servers/storageInit"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
		}
		if models.WorksStorage == nil || len(models.WorksStorage) == 0{
			models.WorksStorage = db.AllWorks()
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
		}
		if models.WorksStorage == nil || len(models.WorksStorage) == 0{
			models.WorksStorage = db.AllWorks()
		}
        
		nextHandler.ServeHTTP(w, r)
    }
}
// End Route Wrappers


func main() {
        
	err:= godotenv.Load()
	if err != nil {
        log.Fatalf("Error loading .env file")
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
	
	// Routes
	http.HandleFunc("GET /{$}", landingHandler) // Match only exactly '/' thanks to {$}
        
	http.HandleFunc("GET /about",func(w http.ResponseWriter, r *http.Request){about.GetHand(w, r, templatesFS)})
        
	http.HandleFunc("GET /contact", func(w http.ResponseWriter, r *http.Request) {contacts.GetHand(w, r, templatesFS)})
	http.HandleFunc("POST /contact", func(w http.ResponseWriter, r *http.Request) {contacts.PostHand(w, r, templatesFS)})

	http.HandleFunc("GET /work", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){work.GetWorksView(w, r, templatesFS)}))
	http.HandleFunc("GET /work/{title}", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){galleries.Gallery(w, r)}))
	
	http.HandleFunc("GET /editor", authMiddleware(func(w http.ResponseWriter, r *http.Request){work.GetHandEditor(w, r, templatesFS)}))
	http.HandleFunc("PUT /editor", func(w http.ResponseWriter, r *http.Request){work.PutHandEditor(w, r, templatesFS)})
	http.HandleFunc("POST /editor", func(w http.ResponseWriter, r *http.Request){work.PostHandEditor(w, r, templatesFS)})
	http.HandleFunc("POST /editor/del", func(w http.ResponseWriter, r *http.Request){work.DelHandEditor(w, r)})


	http.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request){user.GetLoginTmpl(w, r, templatesFS)})
	http.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request){user.Login(w, r)})
	http.HandleFunc("GET /logout", func(w http.ResponseWriter, r *http.Request){user.Logout(w, r)})

	http.HandleFunc("GET /editor/components", func(w http.ResponseWriter, r *http.Request){work.GetEditorComponents(w, r, templatesFS)})
	http.HandleFunc("GET /test", initStorageMiddleware(func(w http.ResponseWriter, r *http.Request){work.GetTestView(w, r, templatesFS)}))



	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}