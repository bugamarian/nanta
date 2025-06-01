package edit

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FindLastNote(dir string) (string, error) {
	var lastModTime time.Time
	var lastFile string

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		info, _ := os.Stat(path)
		if info.ModTime().After(lastModTime) {
			lastModTime = info.ModTime()
			lastFile = path
		}
		return nil
	})
	if lastFile == "" {
		return "", fmt.Errorf("No markdown documents found")
	}
	return lastFile, nil
}
