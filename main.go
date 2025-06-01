package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	title := flag.String("title", "", "Note title")
	templatePath := flag.String("template", "", "Path to template")
	openLast := flag.Bool("last", false, "Open last note")
	flag.Parse()

	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if *openLast {
		last, err := findLastNote(cfg.NotesDir)
		if err != nil {
			fmt.Println("Couldn't find last note:", err)
			os.Exit(1)
		}
		openInEditor(cfg.Editor, last)
		return
	}

	path, err := createNote(cfg, *title, *templatePath)
	if err != nil {
		fmt.Println("Error creating note:", err)
		os.Exit(1)
	}

	openInEditor(cfg.Editor, path)
}
