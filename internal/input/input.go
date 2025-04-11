package input

import (
	"ascii-art-reverse/internal/banner"
	"fmt"
	"os"
)

// InputConfig represents the configuration for ASCII art generation
type InputConfig struct {
	Text       string
	BannerType banner.BannerType
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
		// Two arguments: text and banner type
		config.Text = args[0]

		// Validate banner type
		bannerType, err := banner.ValidateBannerType(args[1])
		if err != nil {
			return nil, err
		}
		config.BannerType = bannerType
	default:
		// Incorrect number of arguments
		return nil, fmt.Errorf("usage: go run . [STRING] [BANNER]\n" +
			"EX: go run . something standard")
	}

	// Validate input text
	if config.Text == "" {
		return nil, fmt.Errorf("text cannot be empty")
	}

	return config, nil
}
