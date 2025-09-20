package main

import (
	"log"
	"net/http"
	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var cfg Config
var tmpl *template.Template

func main() {
	cfg = LoadConfig()

	var err error
	tmpl, err = template.ParseFiles("templates/profile.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", profileHandler)
	r.Get("/{symlink}", symlinkHandler)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting Clover on %s", cfg.Address)
	log.Fatal(http.ListenAndServe(cfg.Address, r))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "profile.html", cfg); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

func symlinkHandler(w http.ResponseWriter, r *http.Request) {
	symlink := r.PathValue("symlink")

	for _, link := range cfg.Links {
		if link.Symlink == symlink {
			http.Redirect(w, r, link.Url, http.StatusSeeOther)
			return
		}
	}

	http.NotFound(w, r)
}
