package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

// SSHConnection represents a single entry in the SSH config file.
type SSHConnection struct {
	Host string
	User string
}

// ReadSSHConfig reads and parses the SSH connections from the config file.
func ReadSSHConfig(r io.Reader) ([]SSHConnection, error) {
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(r)

	var connections []SSHConnection
	var currentEntry SSHConnection

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Split the line into words, we will search by
		// the keyword "Host" and "User"
		parts := strings.Fields(line)

		if len(parts) <= 1 {
			continue
		}

		key, value := parts[0], parts[1]

		switch key {
		case "Host":
			// If a new host directive is found, save the previous entry, then start a new entry
			if currentEntry.Host != "" {
				connections = append(connections, currentEntry)
			}
			currentEntry = SSHConnection{Host: value}
		case "User":
			currentEntry.User = value
		}
	}

	// This final check ensures that the last entry will be appended
	if currentEntry.Host != "" {
		connections = append(connections, currentEntry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return connections, nil
}

// WriteSSHConnections formats and writes the parsed SSHConnection slices to the writer.
func WriteSSHConnections(w io.Writer, connections []SSHConnection) error {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', tabwriter.TabIndent)

	fmt.Fprintln(tw, "HOST\tUSER")
	fmt.Fprintln(tw, "-----\t-----")

	for _, entry := range connections {
		fmt.Fprintf(tw, "%s\t%s\n", entry.Host, entry.User)
	}

	if err := tw.Flush(); err != nil {
		return err
	}

	return nil
}

// ParseConfig is the orchestrator of the functionality that reads and writes the SSH connections.
func ParseConfig(configPath string, writer io.Writer) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("error opening the config file: %v", err)
	}
	defer file.Close()

	connections, err := ReadSSHConfig(file)
	if err != nil {
		return fmt.Errorf("error reading the config file: %v", err)
	}

	if err := WriteSSHConnections(writer, connections); err != nil {
		return fmt.Errorf("error showing the output: %v", err)
	}

	return nil
}
