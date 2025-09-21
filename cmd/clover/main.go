package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eliva1e/clover/internal/assets"
	"github.com/eliva1e/clover/internal/config"
	"github.com/eliva1e/clover/internal/middleware"
)

var cfg *config.Config
var tmpl *template.Template

func main() {
	cfg = config.LoadConfig("config.json")

	var err error
	tmpl, err = template.New("profile").Parse(assets.ProfileTemplate)
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}

	http.HandleFunc("GET /", profileHandler)
	http.HandleFunc("GET /{symlink}", symlinkHandler)

	if cfg.EnableTls {
		go func() {
			log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				target := "https://" + r.Host + r.URL.RequestURI()
				http.Redirect(w, r, target, http.StatusMovedPermanently)
			})))
		}()

		log.Printf("Starting Clover on port 443 (w/ TLS)")
		log.Fatal(http.ListenAndServeTLS(
			":443",
			"server.crt",
			"server.key",
			middleware.LoggingMiddleware(http.DefaultServeMux),
		))
	} else {
		log.Printf("Starting Clover on port 80")
		log.Fatal(http.ListenAndServe(":80", middleware.LoggingMiddleware(http.DefaultServeMux)))
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "profile", cfg); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

func symlinkHandler(w http.ResponseWriter, r *http.Request) {
	symlink := r.PathValue("symlink")

	for _, obj := range cfg.Objects {
		if obj.Symlink == symlink {
			http.Redirect(w, r, obj.Url, http.StatusSeeOther)
			return
		}
	}

	http.NotFound(w, r)
}
