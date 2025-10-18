package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"time"
)

type ClipboardData struct {
	Text string `json:"text"`
}

var clipboard []ClipboardData

func updateJSON(lastCopied string) {
	filename := "copied.json"
	jsonData, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		os.WriteFile(filename, []byte(`[{"text":"randomText"}]`), 0644)
		return
	}
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	if len(jsonData) != 0 {
		err = json.Unmarshal(jsonData, &clipboard)
		if err != nil {
			log.Fatalln("failed to unmarshal json:", err)
		}
	}

	if clipboard[len(clipboard)-1].Text != lastCopied {
		clipboard = append(clipboard, ClipboardData{Text: lastCopied})
	}
	newJson, err := json.Marshal(clipboard)
	if err != nil {
		log.Fatalln("failed to unmarshal json:", err)
	}
	err = os.WriteFile(filename, newJson, 0644)
	if err != nil {
		log.Fatalln("failed to write to file:", err)
	}
	log.Println("File written successfully.")

}

func main() {
	for {
		cmd := exec.Command("wl-paste")
		output, err := cmd.Output()
		if err != nil {
			log.Fatalln("failed to execute wl-paste cmd:", err)
		}
		updateJSON(string(output))
		time.Sleep(5 * time.Second)
	}

}
