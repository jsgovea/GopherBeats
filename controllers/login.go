package controllers

import (
	"fmt"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"html/template"
	"net/http"
	"os"
)

type LoginController struct{}

const redirectURI = ""

var dbHost = os.Getenv("SPOTIFY_CLIENT_ID")
var (
	auth  = spotifyauth.New(spotifyauth.WithRedirectURL("http://127.0.0.1:8080/callback/"), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	ch    = make(chan *oauth2.Token)
	state = "abc123"
)

func (l LoginController) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login controller")
	tpl, err := template.ParseFiles("views/_layout.html", "views/auth/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (l LoginController) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logging in...")
}
