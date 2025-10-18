package main

import (
	"log"
	"os/exec"
)

func main() {

	// Define the command and its arguments
	cmd := exec.Command("wl-paste")

	// Execute the command and capture its standard output
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	saveToJSON(string(out))

}
