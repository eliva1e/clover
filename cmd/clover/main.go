package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eliva1e/clover/internal"
	"github.com/eliva1e/clover/internal/assets"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var cfg internal.Config
var tmpl *template.Template

func main() {
	cfg = internal.LoadConfig("config.json")

	var err error
	tmpl, err = template.New("profile").Parse(assets.ProfileTemplate)
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", profileHandler)
	r.Get("/{symlink}", symlinkHandler)

	log.Printf("Starting Clover on %s", cfg.Address)
	log.Fatal(http.ListenAndServe(cfg.Address, r))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "profile", cfg); err != nil {
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
