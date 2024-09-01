package user

import (
	"Go_servers/db"
	templates "Go_servers/templ"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Get Login html template.
// func GetLoginTmpl(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
// 	tmpl, err:= template.ParseFS(templateFs, "htmlTemplates/login.html")
// 	if err != nil {
// 		log.Printf("Error parsing template: %v", err)
// 		return
// 	}
// 	tmpl.Execute(w,nil)
// }
func GetLoginTmpl(w http.ResponseWriter, r *http.Request, templateFs embed.FS){
	err:= templates.ShowLogin().Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	err:= godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	tokenName:= os.Getenv("SPB_TOKEN_NAME")
	r.ParseForm()
	email := r.FormValue("email")
	pwd := r.FormValue("pwd")
	
	tokenValue , err :=db.Login(email, pwd)
	if err!=nil{
		errorMessage := fmt.Sprintf(`<p id="message-login" class="block w-full p-3 text-center rounded-sm" style="color: red;">%v</p>`, err)
		w.Write([]byte(errorMessage))
		return
	}

	// Set the expiration time to 1 minute from now
	expirationTime := time.Now().Add(30 * time.Minute)
	cookieSupaBase := http.Cookie{
        Name:     tokenName,
        Value:    tokenValue, // Value from supaBase
        Expires:  expirationTime,
        HttpOnly: true,  // Ensures the cookie is only accessible via HTTP(S)
        Secure:   false, // Set to false if serving over HTTP (not HTTPS)
        Path:     "/",   // Scope of the cookie
    }

    http.SetCookie(w, &cookieSupaBase)
	fmt.Println("Logged in succes")
	w.Header().Set("HX-Redirect", "/editor")
	w.WriteHeader(http.StatusOK)
}


func Logout(w http.ResponseWriter, r *http.Request) {
    
	err:= godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	tokenName:= os.Getenv("SPB_TOKEN_NAME")

	expiredCookie := http.Cookie{
        Name:     tokenName,
        Value:    "",
        Expires:  time.Unix(0, 0), // Expiration date set in the past
        HttpOnly: true,
        Secure:   false, // Set to false if serving over HTTP
        Path:     "/",   // Scope of the cookie
    }

    // Set the expired cookie, which will delete it from the browser
    http.SetCookie(w, &expiredCookie)
	  // Optionally, send a response to the frontend
	  w.Write([]byte("LOGGED OUT"))
}