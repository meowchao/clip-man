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
	jsonData1, err := os.ReadFile("copied.json") //this is previously copied text
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	if string(jsonData) == string(jsonData1) {
		log.Println("JSON data is equal")
	} else {
		log.Println("JSON data is not equal")
		err = os.WriteFile("copied.json", jsonData, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("File written successfully.")
	}

}
