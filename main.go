package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"music_parser.com/m/controllers"
)

func main() {
	fmt.Println("Running server on http://localhost:8080/")
	router := chi.NewRouter()

	// Config for static files
	worDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(worDir, "static"))
	FileServer(router, "/static", filesDir)

	// Routes
	router.Get("/", controllers.Home)
	router.Get("/login", controllers.Login)

	http.ListenAndServe(":8080", router)

}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("URL parameters are not allowed.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
