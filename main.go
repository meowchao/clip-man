package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	copiedText := []interface{}{}

	// Define the command and its arguments
	cmd := exec.Command("wl-paste")

	// Execute the command and capture its standard output
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	copiedText = append(copiedText, string(out))
	fmt.Println(copiedText)
	os.WriteFile("copied.txt", (out), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File written successfully.")

}
