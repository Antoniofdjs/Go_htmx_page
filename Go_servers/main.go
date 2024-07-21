package main

import (
	contacts "Go_servers/handlers/contact"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type PictureData struct {
    Title      string `json:"title"`
    Picture []byte `json:"picture"`
}

type Picture struct{
	Title string
	Path string
}

func main() {
	// Define a struct to hold the JSON data

	landingHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl := template.Must(template.ParseFiles("htmlTemplates/index.html"))
			tmpl.Execute(w, nil)
		}
	}

	// Handlers ----------------------------------------------------------------------------
	// work
	workGetHand := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("htmlTemplates/work.html"))
		pictures:= map[string][]Picture{
			"Pictures": {
				{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
				{Title:"FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
				{Title:"ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
				{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
				{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
				{Title:"FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
				{Title:"ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
				{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
				{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
				{Title:"FOREST - EL YUNQUE", Path: "../static/images/userWorks/forest.jpg"},
				{Title:"ICELAND - BLACK SANDS", Path: "../static/images/userWorks/iceland.jpg"},
				{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
			},
		}
		tmpl.Execute(w, pictures)
	}

	workHandEditor := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("htmlTemplates/editorTemplates/workEditor.html"))
		pictures:= map[string][]Picture{
			"Pictures": {
				{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
				{Title:"FOREST", Path: "../static/images/userWorks/forest.jpg"},
				{Title:"ICELAND", Path: "../static/images/userWorks/iceland.jpg"},
				{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
				{Title:"BEACH", Path: "../static/images/userWorks/beach.jpg"},
				{Title:"FOREST", Path: "../static/images/userWorks/forest.jpg"},
				{Title:"ICELAND", Path: "../static/images/userWorks/iceland.jpg"},
				{Title:"FOOD", Path: "../static/images/userWorks/food.jpg"},
			},
		}
		tmpl.Execute(w, pictures)
	}
	

	workPostHand := func(w http.ResponseWriter, r *http.Request) {
		var data PictureData
		fmt.Println("POST ACTIVATED")
		err := r.ParseMultipartForm(10 << 20) // 10 MB limit
        if err != nil {
            http.Error(w, "Failed to parse form", http.StatusInternalServerError)
            return
        }
		fmt.Println("AFTER PARSE")
		data.Title =  r.FormValue("title")
		fmt.Printf("Tile: %v",data.Title)
		
		//  Check file of picture
		PicBytes, headers, err := r.FormFile("picture")
		if err != nil {
            http.Error(w, "Failed to get file", http.StatusBadRequest)
            return
        }
        defer PicBytes.Close() //Like a file close in python, prevent leaks

		picName := headers.Filename
		path := fmt.Sprintf("static/images/userWorks/%s", picName)
		fmt.Printf("\nPic name: %v", path)

		//  Create path for picture, picFile is destination for the picture
		picFile, err := os.Create(path)
        if err != nil {
            http.Error(w, "Failed to create file", http.StatusInternalServerError)
            return
        }

		 // Send the pic bytes the 'picFile'
		 _, err = io.Copy(picFile, PicBytes)
		 if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		 }

		// htmlContent := fmt.Sprintf(`
        // <div class="flex w-full h-full flex-wrap pb-10 object-center">
        //     <div class="w-full p-1 md:p-2">
        //         <img
        //             alt="gallery"
        //             class="block h-full w-full rounded-lg object-cover object-center"
        //             src="%s"
        //         />
        //     </div>
        //     <h2 class="text-gray-600 w-full text-center mt-2 hover:text-amber-600 font-serif text-xl">%s</h2>
        // </div>`, data.PictureURL, data.Title)

		// w.Header().Set("Content-Type", "text/html")
        // w.Write([]byte(htmlContent))
	}

	// about
	aboutHand := func(w http.ResponseWriter, r *http.Request){
		tmpl := template.Must(template.ParseFiles("htmlTemplates/about.html"))
		tmpl.Execute(w, nil)
	}
	// ---------------------------------------------------------------------------------------
	
	
	// Serve output.css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", landingHandler)
	
	http.HandleFunc("GET /about", aboutHand)
	

	http.HandleFunc("GET /work", workGetHand)
	http.HandleFunc("GET /work/editor", workHandEditor)
	http.HandleFunc("POST /work", workPostHand)
	
	http.HandleFunc("GET /contact", contacts.ContactGetHand)
	http.HandleFunc("POST /contact", contacts.ContactPostHand)
	
	// Playa -> /playa {contentHTML}
	// Start server
	log.Fatal(http.ListenAndServe(":8000", nil))
}