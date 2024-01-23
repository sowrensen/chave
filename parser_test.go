package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"
)

func TestReadSSHConfig(t *testing.T) {
	// Mock config content
	const sshConfigContent = `
	Host github.com
		User git

	Host server_1
		User ubuntu
	`

	// Initiate a reader with the mock content
	reader := strings.NewReader(sshConfigContent)

	entries, err := ReadSSHConfig(reader)

	if err != nil {
		t.Fatalf("ReadSSHConfig() error: %v", err)
	}

	// Expected output
	expectations := []SSHConnection{
		{Host: "github.com", User: "git"},
		{Host: "server_1", User: "ubuntu"},
	}

	for i, entry := range entries {
		if expectations[i] != entry {
			t.Errorf("Expected entry %#v, got %#v", expectations[i], entry)
		}
	}
}

func TestWriteSSHConnections(t *testing.T) {
	// Connections to write
	connections := []SSHConnection{
		{Host: "github.com", User: "git"},
		{Host: "server_1", User: "ubuntu"},
	}

	// Define the expected output using a tabwriter,
	// we will use a buffer for the output
	var expectedBuf bytes.Buffer
	expectedTw := tabwriter.NewWriter(&expectedBuf, 0, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(expectedTw, "HOST\tUSER")
	fmt.Fprintln(expectedTw, "-----\t-----")
	fmt.Fprintln(expectedTw, "github.com\tgit")
	fmt.Fprintln(expectedTw, "server_1\tubuntu")
	expectedTw.Flush()
	expectedOutput := expectedBuf.String()

	// Capture the actual output in another buffer
	var actualBuf bytes.Buffer
	err := WriteSSHConnections(&actualBuf, connections)
	if err != nil {
		t.Fatalf("WriteSSHConnections() error: %v", err)
	}

	actualOutput := actualBuf.String()
	if actualOutput != expectedOutput {
		t.Errorf("WriteSSHConnections() output %q, expected %q", actualOutput, expectedOutput)
	}
}
