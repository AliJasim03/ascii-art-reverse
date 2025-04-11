package main

import (
	"fmt"
	"log"

	"ascii-art-reverse/internal/banner"
	"ascii-art-reverse/internal/input"
	"ascii-art-reverse/internal/output"
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
	renderedOutput, err := artRenderer.RenderText(config.Text)
	if err != nil {
		log.Fatalf("Rendering error: %v", err)
	}

	// Determine output method
	if config.OutputFile != "" {
		// Write to file
		err := output.WriteOutputToFile(config.OutputFile, renderedOutput)
		if err != nil {
			log.Fatalf("File writing error: %v", err)
		}
		fmt.Printf("Output written to %s\n", config.OutputFile)
	} else {
		// Print to console
		for _, line := range renderedOutput {
			fmt.Println(line)
		}
	}
}
