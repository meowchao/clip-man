package main

import (
	"encoding/json"
	"log"
	"os"
)

type CopiedText struct {
	Text string `json:"text"`
}

func saveToJSON(text string) {
	copiedText := CopiedText{Text: text}
	jsonData, err := json.Marshal(copiedText)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	err = os.WriteFile("copied.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File written successfully.")
}
