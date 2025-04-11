package main

import (
	"fmt"
	"log"

	"ascii-art-reverse/internal/banner"
	"ascii-art-reverse/internal/input"
	"ascii-art-reverse/internal/renderer"
)

func main() {
	// Parse input arguments
	config, err := input.ParseArgs()
	if err != nil {
		log.Fatalf("Input error: %v", err)
	}

	// Load the specified banner
	bannerArt, err := banner.LoadBanner(config.BannerType)
	if err != nil {
		log.Fatalf("Banner loading error: %v", err)
	}

	// Create renderer
	artRenderer := renderer.NewRenderer(bannerArt)

	// Render the text
	output, err := artRenderer.RenderText(config.Text)
	if err != nil {
		log.Fatalf("Rendering error: %v", err)
	}

	// Print the output
	for _, line := range output {
		fmt.Println(line)
	}
}
