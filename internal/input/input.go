package input

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-reverse/internal/banner"
)

// InputConfig represents the configuration for ASCII art generation
type InputConfig struct {
	Text       string
	BannerType banner.BannerType
	OutputFile string
}

// ParseArgs parses command-line arguments
func ParseArgs() (*InputConfig, error) {
	args := os.Args[1:]

	// Default configuration
	config := &InputConfig{
		BannerType: banner.Standard,
	}

	// Handle different argument scenarios
	switch len(args) {
	case 1:
		// Single argument: text
		config.Text = args[0]
	case 2:
		// Two possibilities:
		// 1. Text and banner type
		// 2. Output flag and text
		if strings.HasPrefix(args[0], "--output=") {
			// Output flag and text
			config.OutputFile = strings.TrimPrefix(args[0], "--output=")
			config.Text = args[1]
		} else {
			// Text and banner type
			config.Text = args[0]
			bannerType, err := banner.ValidateBannerType(args[1])
			if err != nil {
				return nil, err
			}
			config.BannerType = bannerType
		}
	case 3:
		// Three arguments: output flag, text, and banner type
		if !strings.HasPrefix(args[0], "--output=") {
			return nil, fmt.Errorf("usage: go run . [OPTION] [STRING] [BANNER]\n" +
				"EX: go run . --output=<fileName.txt> something standard")
		}
		config.OutputFile = strings.TrimPrefix(args[0], "--output=")
		config.Text = args[1]

		bannerType, err := banner.ValidateBannerType(args[2])
		if err != nil {
			return nil, err
		}
		config.BannerType = bannerType
	default:
		// Incorrect number of arguments
		return nil, fmt.Errorf("usage: go run . [OPTION] [STRING] [BANNER]\n" +
			"EX: go run . --output=<fileName.txt> something standard")
	}

	// Validate input text
	if config.Text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	return config, nil
}
