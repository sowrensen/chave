package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Try to get current user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v\n", err)
	}

	// Build the config path and pass it to the parser
	configPath := filepath.Join(homeDir, ".ssh", "config")
	if err := ParseConfig(configPath, os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
