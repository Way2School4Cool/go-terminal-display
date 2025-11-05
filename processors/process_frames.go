package processors

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg" // Register JPEG format
	"os"
	"strings"
	"sync"

	gloss "github.com/charmbracelet/lipgloss"
)

func ReadImage(location string) [][]color.Color {
	if location == "" {
		fmt.Print("No Image Location Provided")
		return nil
	}

	if strings.Split(location, ".")[1] != "jpg" {
		fmt.Print("Unsupported Image Type:", location)
		return nil
	}

	// Open the image file
	imageFile, err := os.Open(location)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return nil
	}
	defer imageFile.Close()

	// Decode the image
	img, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return nil
	}

	// Get image bounds
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y

	// Create 2D array to store colors
	pixels := make([][]color.Color, height)
	for y := range pixels {
		pixels[y] = make([]color.Color, width)
		for x := 0; x < width; x++ {
			pixels[y][x] = img.At(x+bounds.Min.X, y+bounds.Min.Y)
		}
	}

	return pixels
}

func ProcessImageToTerminal(pixels [][]color.Color, width int, height int) []gloss.Style {
	// Calculate scaling factors
	imgHeight := len(pixels)
	imgWidth := len(pixels[0])
	scaleX := float64(imgWidth) / float64(width)
	scaleY := float64(imgHeight) / float64(height)

	// Pre-allocate the slice with the exact size needed
	styles := make([]gloss.Style, width*height)
	var wg sync.WaitGroup

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			wg.Add(1)
			go func(x, y int) {
				defer wg.Done()
				// Determine the corresponding pixel in the image
				imgX := int(float64(x) * scaleX)
				imgY := int(float64(y) * scaleY)

				// Get the color of the the specific scaled pixel
				// NOTE: this does not average colors, it just picks one pixel
				pixelColor := pixels[imgY][imgX]
				r, g, b, _ := pixelColor.RGBA()
				r8 := uint8(r >> 8)
				g8 := uint8(g >> 8)
				b8 := uint8(b >> 8)

				// Write directly to the pre-allocated slice
				index := y*width + x
				styles[index] = gloss.NewStyle().Background(gloss.Color(fmt.Sprintf("#%02x%02x%02x", r8, g8, b8))).SetString(" ")
			}(x, y)
		}
	}
	wg.Wait()

	println(styles)

	return styles
}
