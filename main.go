package main

import (
	"fmt"
	frames "go-terminal-display/processors"
	"time"

	"golang.org/x/term"
)

var imageLocation string = "Images/cat.jpg"

func main() {
	width, height, err := term.GetSize(0)
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
