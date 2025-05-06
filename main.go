package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

var slashCommentRegex = regexp.MustCompile(`(?s)^//\s*([^\n]+)\r?\n(.*)`)
var htmlCommentRegex = regexp.MustCompile(`(?s)^<!--\s*([^\n]+)\s*-->\r?\n(.*)`)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	targetDir := os.Args[1]
	info, err := os.Stat(targetDir)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Error: Direktori tujuan '%s' tidak ditemukan.", targetDir)
		}
		log.Fatalf("Error saat memeriksa direktori tujuan '%s': %v", targetDir, err)
	}
	if !info.IsDir() {
		log.Fatalf("Error: Path tujuan '%s' bukanlah sebuah direktori.", targetDir)
	}

	err = clipboard.Init()
	if err != nil {
		log.Fatalf("Gagal menginisialisasi clipboard: %v", err)
	}

	var previousClipboardContent []byte

	for {
		currentClipboardContent := clipboard.Read(clipboard.FmtText)

		if currentClipboardContent != nil && string(currentClipboardContent) != string(previousClipboardContent) {
			parseAndWrite(string(currentClipboardContent), targetDir)
			previousClipboardContent = currentClipboardContent
		}

		time.Sleep(2 * time.Second)
	}
}

func parseAndWrite(content string, baseDir string) {
	var relativePath string
	var fileContent string
	var matched bool

	if matches := slashCommentRegex.FindStringSubmatch(content); len(matches) == 3 {
		relativePath = strings.TrimSpace(matches[1])
		fileContent = matches[2]
		matched = true
	}

	if !matched {
		if matches := htmlCommentRegex.FindStringSubmatch(content); len(matches) == 3 {
			relativePath = strings.TrimSpace(matches[1])
			fileContent = matches[2]
			matched = true
		}
	}

	if matched {
		if strings.Contains(relativePath, "..") || filepath.IsAbs(relativePath) || relativePath == "" {
			log.Printf("Peringatan: Path relatif tidak valid atau kosong ('%s'), dilewati.", relativePath)
			return
		}

		fullPath := filepath.Clean(filepath.Join(baseDir, relativePath))
		parentDir := filepath.Dir(fullPath)
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			log.Printf("Error membuat direktori induk '%s': %v", parentDir, err)
			return
		}

		fileExists := false
		if _, err := os.Stat(fullPath); err == nil {
			fileExists = true
		} else if !errors.Is(err, os.ErrNotExist) {
			log.Printf("Error saat memeriksa status file '%s': %v", fullPath, err)
		}

		err := os.WriteFile(fullPath, []byte(fileContent), 0644)
		if err != nil {
			log.Printf("Error menulis file '%s': %v", fullPath, err)
		} else {
			if fileExists {
				fmt.Printf("diubah: %s\n", fullPath)
			} else {
				fmt.Printf("dibuat: %s\n", fullPath)
			}
		}
	}
}
