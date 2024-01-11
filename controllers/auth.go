package controllers

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/_layout.html", "views/auth/login.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.ExecuteTemplate(w, "layout", nil)
}
