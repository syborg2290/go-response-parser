package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Check for command-line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: response-analyzer <response>")
		return
	}

	// Retrieve the response argument
	response := os.Args[1]

	// Parse the response status line to get the status code
	statusCode := parseStatusCode([]byte(response))
	fmt.Println("Response Status Code:", statusCode)

	// Parse and print the response headers
	headers := parseHeaders([]byte(response))
	fmt.Println("Response Headers:")
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Extract and print specific headers
	contentType := headers["Content-Type"]
	serverInfo := headers["Server"]

	fmt.Println("Content-Type:", contentType)
	fmt.Println("Server Info:", serverInfo)

	// Example: Check if the "Server" header contains specific information
	if containsString([]byte(serverInfo), "Apache") {
		fmt.Println("Server uses Apache.")
	} else {
		fmt.Println("Server is not using Apache.")
	}

	// Example: Check if the response body contains a specific string
	if containsString([]byte(response), "error") {
		fmt.Println("Response contains the word 'error'.")
	} else {
		fmt.Println("Response does not contain the word 'error'.")
	}

	// Additional analysis logic can be added here
	// ...
}

// Helper function to check if a byte slice contains a specific string
func containsString(data []byte, target string) bool {
	return bytes.Index(data, []byte(target)) != -1
}

// Helper function to parse the response status code from the headers
func parseStatusCode(response []byte) string {
	headers := bytes.Split(response, []byte("\r\n\r\n"))[0]
	headerLines := bytes.Split(headers, []byte("\r\n"))
	for _, line := range headerLines {
		if bytes.HasPrefix(line, []byte("HTTP/")) {
			fields := strings.Fields(string(line))
			if len(fields) >= 2 {
				return fields[1]
			}
		}
	}
	return "Unknown"
}

// Helper function to parse headers from key: value format
// Helper function to parse headers from key: value format
func parseHeaders(response []byte) map[string]string {
	headers := make(map[string]string)
	headerLines := bytes.Split(response, []byte("\r\n\r\n"))[0]
	for _, line := range headerLines {
		lineStr := string(line)
		if strings.Contains(lineStr, ":") {
			parts := strings.SplitN(lineStr, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				headers[key] = value
			}
		}
	}
	return headers
}
