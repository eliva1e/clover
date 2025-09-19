package main

import (
	"fmt"
	"log"
	"net/http"
)

var cfg Config

func main() {
	cfg = LoadConfig()

	http.HandleFunc("GET /", profileHandler)
	http.HandleFunc("GET /{symlink}", symlinkHandler)

	log.Printf("Starting Clover on %s", cfg.Address)
	err := http.ListenAndServe(cfg.Address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "profile")
}

func symlinkHandler(w http.ResponseWriter, r *http.Request) {
	symlink := r.PathValue("symlink")

	for _, link := range cfg.Links {
		if link.Symlink == symlink {
			http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
			return
		}
	}

	http.NotFound(w, r)
}
