package main

import (
	"fmt"
	frames "go-terminal-display/processors"
	"os"
	"time"

	"golang.org/x/term"
)

var imageLocation string = "Images/cat2.jpg"

func main() {
	// allow overriding the default image via first CLI argument
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Printf("Usage: %s [image-file]\n", os.Args[0])
			os.Exit(0)
		}
		imageLocation = os.Args[1]
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	println("Width:", width, "Height:", height)

	// Start a performance timer
	start := time.Now()

	// project 1: take an image and "project" it to the terminal
	imageData := frames.ReadImage(imageLocation)

	fmt.Println("ImageData Size: ", len(imageData), "x", len(imageData[0]))

	style := frames.ProcessImageToTerminal(imageData, width, height)

	for _, s := range style {
		fmt.Print(s)
	}
	fmt.Println()

	// Stop the timer
	fmt.Println("Processing Time:", time.Since(start))

}
