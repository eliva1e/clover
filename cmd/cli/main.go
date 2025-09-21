package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"

	_ "embed"

	"github.com/eliva1e/clover/internal/config"
	"github.com/eliva1e/clover/internal/assets"
)

func main() {
	configPath := flag.String("config", "", "Path to the Clover config file")
	flag.Parse()

	if *configPath == "" {
		fmt.Printf("Usage: %v -config <path-to-config>\n", os.Args[0])
		os.Exit(1)
	}

	cfg := config.LoadConfig(*configPath)

	tmpl, err := template.New("profile").Parse(assets.ProfileTemplate)
	if err != nil {
		fmt.Printf("failed to parse templates: %v", err)
	}

	var html bytes.Buffer
	if err := tmpl.ExecuteTemplate(&html, "profile", cfg); err != nil {
		fmt.Printf("failed to execute template: %v", err)
	}

	if err = os.MkdirAll("dist", os.ModePerm); err != nil {
		fmt.Printf("failed to create directory: %v", err)
	}

	htmlString := html.String()
	htmlString = strings.ReplaceAll(htmlString, "\n", "")

	if err = os.WriteFile("dist/index.html", []byte(htmlString), 0644); err != nil {
		fmt.Printf("failed to write file: %v", err)
	}

	fmt.Println("☘️ Output generated to ./dist!")
}
