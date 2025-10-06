package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/eliva1e/clover/internal/assets"
	"github.com/eliva1e/clover/internal/config"
)

func exitLog(format string, a ...any) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		exitLog("Usage: %v <path-to-config>\n", os.Args[0])
	}

	configPath := os.Args[1]

	if configPath == "" {
		exitLog("Usage: %v <path-to-config>\n", os.Args[0])
	}

	cfg := config.LoadConfig(configPath)

	profileTmpl, err := template.New("profile").Parse(assets.ProfileTemplate)
	if err != nil {
		exitLog("failed to parse templates: %v", err)
	}

	redirectTmpl, err := template.New("redirect").Parse(assets.RedirectTemplate)
	if err != nil {
		exitLog("failed to parse templates: %v", err)
	}

	var profileHtml bytes.Buffer
	if err := profileTmpl.ExecuteTemplate(&profileHtml, "profile", cfg); err != nil {
		exitLog("failed to execute template: %v", err)
	}

	if err = os.MkdirAll("dist", os.ModePerm); err != nil {
		exitLog("failed to create directory: %v", err)
	}

	if err = os.WriteFile("dist/index.html", profileHtml.Bytes(), 0644); err != nil {
		exitLog("failed to write file: %v", err)
	}

	for _, obj := range cfg.Objects {
		if obj.Type == "button" {
			if err = os.MkdirAll("dist/"+obj.Symlink, os.ModePerm); err != nil {
				exitLog("failed to create directory: %v", err)
			}

			var redirectHtml bytes.Buffer
			if err := redirectTmpl.ExecuteTemplate(&redirectHtml, "redirect", obj); err != nil {
				exitLog("failed to execute template: %v", err)
			}

			if err = os.WriteFile(
				"dist/"+obj.Symlink+"/index.html",
				redirectHtml.Bytes(),
				0644,
			); err != nil {
				exitLog("failed to write file: %v", err)
			}
		}
	}

	fmt.Println("☘️ Output generated to ./dist!")
}
