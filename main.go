package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

func main() {
	// Try to get current user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		return
	}

	// Build the config path
	configPath := filepath.Join(homeDir, ".ssh", "config")

	// Try to read the config file
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("Error opening the config file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initiate tabwriter with stdout as the output device and some formatting
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)

	// Print the headers
	fmt.Fprintln(writer, "Host\tUser")
	fmt.Fprintln(writer, "-----\t-----")

	var host, user string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Split the line into words, we will search by
		// the keyword "Host" and "User"
		parts := strings.Fields(line)
		if len(parts) > 1 {
			key, value := parts[0], parts[1]

			if key == "Host" {
				host = value
			} else if key == "User" {
				user = value

				// Show only when both host and user are available
				fmt.Fprintf(writer, "%s\t%s\n", host, user)
				host, user = "", ""
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the config file: %v\n", err)
	}

	if err := writer.Flush(); err != nil {
		fmt.Printf("Error showing the output: %v\n", err)
	}
}
