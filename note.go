package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func createNote(cfg *Config, title, templateName string) (string, error) {
	ts := time.Now()
	var filename string
	if cfg.Savemode == "daily" {
		filename = ts.Format("2006-01-02") + ".md"
	} else {
		filename = ts.Format("2006-01-02_15-04") + ".md"
	}

	path := filepath.Join(cfg.NotesDir, filename)
	var content string
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	if templateName == "" {
		templateName = "default.md"
	}
	templatePath := filepath.Join(
		configDir,
		"nanta",
		"templates",
		templateName,
	)
	tmpl, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	content = string(tmpl)
	content = strings.ReplaceAll(
		content,
		"{{.Timestamp}}",
		ts.Format("2006-01-02 15:04"),
	)
	content = strings.ReplaceAll(content, "{{.Title}}", title)

	if cfg.Savemode == "daily" && fileExists(path) {
		f, _ := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		f.WriteString("\n\n" + content)
		f.Close()
	} else {
		os.MkdirAll(cfg.NotesDir, 0755)
		os.WriteFile(path, []byte(content), 0644)
	}

	return path, nil
}

func titleOrTimestamp(title string, ts time.Time) string {
	if title != "" {
		return title
	}
	return ts.Format("2006-01-02 15:04")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func openFile(modifier, path string) {
	// Generic, takes any command as arg, for example <cat, nvim, glow>
	cmd := exec.Command(modifier, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to open:", err)
	}
}

func findLastNote(dir string) (string, error) {
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
		return "", fmt.Errorf("no markdown notes found")
	}
	return lastFile, nil
}
