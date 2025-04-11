package banner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// BannerType represents different ASCII art banner styles
type BannerType string

const (
	Standard   BannerType = "standard"
	Shadow     BannerType = "shadow"
	Thinkertoy BannerType = "thinkertoy"
)

// Banner represents the ASCII art for a specific banner type
type Banner struct {
	characters map[rune][]string
}

// LoadBanner loads the ASCII art for a specific banner type
func LoadBanner(bannerType BannerType) (*Banner, error) {
	// Construct the path to the banner file
	filename := filepath.Join("assets", "banners", string(bannerType)+".txt")

	// Open the banner file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open banner file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	// Create banner instance
	banner := &Banner{
		characters: make(map[rune][]string),
	}

	// Scan the file
	scanner := bufio.NewScanner(file)

	// Track current character being processed
	var currentChar rune = 32 // Start from space character
	var currentCharLines []string

	for scanner.Scan() {
		line := scanner.Text()

		// Empty line indicates end of a character's representation
		if line == "" {
			if len(currentCharLines) > 0 {
				banner.characters[currentChar] = currentCharLines
				currentCharLines = []string{}
				currentChar++
			}
			continue
		}

		currentCharLines = append(currentCharLines, line)
	}

	// Add the last character if not already added
	if len(currentCharLines) > 0 {
		banner.characters[currentChar] = currentCharLines
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %v", err)
	}

	return banner, nil
}

// GetCharacter returns the ASCII art representation of a specific character
func (b *Banner) GetCharacter(char rune) ([]string, bool) {
	// Ensure the character is within printable ASCII range
	if char < 32 || char > 126 {
		char = 32 // Default to space for out-of-range characters
	}

	lines, exists := b.characters[char]
	return lines, exists
}

// ValidateBannerType checks if the given banner type is valid
func ValidateBannerType(bannerType string) (BannerType, error) {
	switch BannerType(bannerType) {
	case Standard, Shadow, Thinkertoy:
		return BannerType(bannerType), nil
	default:
		return "", fmt.Errorf("invalid banner type. Choose from: standard, shadow, thinkertoy")
	}
}
