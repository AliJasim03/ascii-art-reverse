package renderer

import (
	"fmt"
	"strings"

	"ascii-art-reverse/internal/banner"
)

// Renderer handles the rendering of ASCII art
type Renderer struct {
	banner *banner.Banner
}

// NewRenderer creates a new Renderer with a specific banner
func NewRenderer(b *banner.Banner) *Renderer {
	return &Renderer{
		banner: b,
	}
}

// RenderText converts a string to ASCII art
func (r *Renderer) RenderText(text string) ([]string, error) {
	// Split the input text into lines
	lines := strings.Split(text, "\\n")

	var output []string

	for _, line := range lines {
		// If line is empty, add an empty line to output
		if line == "" {
			output = append(output, "")
			continue
		}

		// Render the current line
		lineOutput, err := r.renderLine(line)
		if err != nil {
			return nil, err
		}
		output = append(output, lineOutput...)
	}

	return output, nil
}

// renderLine renders a single line of text
func (r *Renderer) renderLine(line string) ([]string, error) {
	// Prepare character lines for the entire line
	charLines := make([][]string, len(line))

	// Maximum height of characters in this line
	maxHeight := 8

	// Get ASCII art for each character
	for i, char := range line {
		charArt, exists := r.banner.GetCharacter(char)
		if !exists {
			return nil, fmt.Errorf("character not found in banner: %c", char)
		}
		charLines[i] = charArt
	}

	// Construct the output line
	outputLines := make([]string, maxHeight)
	for i := 0; i < maxHeight; i++ {
		var lineBuilder strings.Builder
		for _, charLine := range charLines {
			lineBuilder.WriteString(charLine[i])
		}
		outputLines[i] = lineBuilder.String()
	}

	return outputLines, nil
}
