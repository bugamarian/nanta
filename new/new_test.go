package new

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestCreateNote(t *testing.T) {
	tmpDir := t.TempDir()
	tmplDir := filepath.Join(tmpDir, "templates")
	os.MkdirAll(tmplDir, 0755)
	tmpl := `# Title: {{.Title}}
	Time: {{.Timestamp}}

	---

	- [] xxxx
	- [] xxxx
	- [] xxxx
	`
	tmplPath := filepath.Join(tmplDir, "note.tmpl")
	err := os.WriteFile(
		tmplPath,
		[]byte(tmpl),
		0644,
	)
	if err != nil {
		t.Fatalf("Failed to write template: %v", err)
	}

	viper.Set("notes_dir", tmpDir)
	viper.Set("template", tmplPath)
	viper.Set("title", "Test Note")
	viper.Set(
		"modifier",
		"cat",
	) // This indicates that the modifier might not be a good idea

	CreateNote()

	files, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("Failed to read note directory: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("Expected at least one note to be created")
	}

	// Validate note content
	content, err := os.ReadFile(filepath.Join(tmpDir, files[0].Name()))
	if err != nil {
		t.Fatalf("Failed to read note content: %v", err)
	}
	if !strings.Contains(string(content), "Test Note") {
		t.Error("Note content missing title")
	}
}
