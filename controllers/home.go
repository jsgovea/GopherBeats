package controllers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/_layout.html", "views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home page",
		Message: "Hello world from Go!",
	}

	t.ExecuteTemplate(w, "layout", data)
}
