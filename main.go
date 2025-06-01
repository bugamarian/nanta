package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	title := flag.String("title", "", "Note title")
	templatePath := flag.String("template", "", "Path to template")
	openLast := flag.Bool("last", false, "Open last note")
	configPath := flag.String("config", "", "Home directory")
	preview := flag.Bool(
		"preview",
		false,
		"Preview, as in preview in other tool",
	)
	flag.Parse()

	configFile := *configPath
	if *configPath == "" {
		xdgDir, err := os.UserConfigDir()
		if err != nil {
			log.Fatalf("Could not find config dir")
		}
		configAppName := "nanta"
		configName := "config.yaml"
		configFile = fmt.Sprintf("%s/%s/%s", xdgDir, configAppName, configName)

	}
	cfg, err := LoadConfig(configFile)
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
		if *preview {
			openFile(cfg.Previewer, last)
			return
		}
		openFile(cfg.Editor, last)
		return
	}

	path, err := createNote(cfg, *title, *templatePath)
	if err != nil {
		fmt.Println("Error creating note:", err)
		os.Exit(1)
	}
	if !*preview {
		openFile(cfg.Editor, path)
	}
}
