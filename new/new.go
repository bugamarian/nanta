package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/spf13/viper"

	"github.com/bugamarian/nanta/common"
)

type Template struct {
	Timestamp string
	Title     string
}

func CreateNote() {
	notesDir := viper.GetString("notes_dir")
	templateFile := viper.GetString("template")
	title := viper.GetString("title")
	modifier := viper.GetString("modifier")

	now := time.Now()
	timestamp := now.Format("2006-01-02")

	err := os.MkdirAll(notesDir, 0755)
	if err != nil {
		log.Fatalf("Failed to create notes directory: %s", err)
	}

	notePath := filepath.Join(notesDir, fmt.Sprintf("%s.md", string(timestamp)))
	data := Template{
		Timestamp: now.Format("2006-01-02 15:04"),
		Title:     title,
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf(
			"Failed to parse the template: %s, file: %s",
			err,
			templateFile,
		)
	}

	noteFile, err := os.Create(notePath)
	if err != nil {
		log.Fatalf("Failed to create note: %s", err)
	}

	defer noteFile.Close()

	err = tmpl.Execute(noteFile, data)
	if err != nil {
		log.Fatalf("Failer to render note: %s", err)
	}

	common.OpenFile(modifier, notePath)
}
