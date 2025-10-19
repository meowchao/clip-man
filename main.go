package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	filename     = "copied.json"
	maxHistory   = 10
	pollInterval = 5 * time.Second
)

type ClipboardData struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	clipboard []ClipboardData
	mu        sync.Mutex
)

func loadClipboard() error {
	mu.Lock()
	defer mu.Unlock()

	jsonData, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		clipboard = []ClipboardData{}
		return saveClipboard()
	}
	if err != nil {
		return err
	}

	if len(jsonData) > 0 {
		err = json.Unmarshal(jsonData, &clipboard)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveClipboard() error {
	newJson, err := json.MarshalIndent(clipboard, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, newJson, 0644)
}

func updateClipboard(lastCopied string) error {
	mu.Lock()
	defer mu.Unlock()

	lastCopied = string([]byte(lastCopied)) // trim trailing newline if needed
	if len(lastCopied) == 0 {
		return nil // ignore empty clipboard
	}

	// Check if this is different from the last entry
	if len(clipboard) > 0 && clipboard[len(clipboard)-1].Text == lastCopied {
		return nil // duplicate, skip
	}

	// Enforce maximum history size
	if len(clipboard) >= maxHistory {
		clipboard = clipboard[1:]
	}

	clipboard = append(clipboard, ClipboardData{
		Text:      lastCopied,
		Timestamp: time.Now(),
	})

	return saveClipboard()
}

func getClipboardContent() (string, error) {
	cmd := exec.Command("wl-paste")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func main() {
	// Load existing clipboard history
	if err := loadClipboard(); err != nil {
		log.Fatalf("Failed to load clipboard: %v", err)
	}
	log.Println("Clipboard monitor started")

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-sigChan:
			log.Println("Shutting down gracefully")
			return
		case <-ticker.C:
			content, err := getClipboardContent()
			if err != nil {
				log.Printf("Failed to read clipboard: %v", err)
				continue
			}

			if err := updateClipboard(content); err != nil {
				log.Printf("Failed to update clipboard: %v", err)
			}
		}
	}
}
