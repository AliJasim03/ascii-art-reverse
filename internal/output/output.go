package output

import (
	"os"
	"strings"
)

// WriteOutputToFile writes the ASCII art to a specified file
// with each line suffixed by a '$' character
func WriteOutputToFile(filename string, content []string) error {
	// Add $ at the end of each line for the specified output format
	var formattedContent []string
	for _, line := range content {
		formattedContent = append(formattedContent, line+"$")
	}

	// Join the lines with newline character
	fileContent := strings.Join(formattedContent, "\n")

	// Write to file with read/write permissions for the owner
	return os.WriteFile(filename, []byte(fileContent), 0644)
}
