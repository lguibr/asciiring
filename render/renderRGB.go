package render

import (
	"fmt"
	"math"
	"strings"

	"github.com/lguibr/asciiring/types"
)

// ASCII characters for grayscale, from lighter to darker
const asciiChars = "@80GCLft1i;:,:.  "

// Dividing factor to convert RGB color space to grayscale
const grayFactor = 255.0 / float64(len(asciiChars)-1)

// Grayscale conversion factors for RGB components
const (
	RFactor = 1
	GFactor = 1
	BFactor = 1
)

// rgbToGray converts an RGB pixel to grayscale using the luminosity method
func rgbToGray(pixel types.RGBPixel) uint8 {
	r := RFactor * float64(pixel.R)
	g := GFactor * float64(pixel.G)
	b := BFactor * float64(pixel.B)
	return uint8(r + g + b)
}

// grayToAscii maps a grayscale value to an ASCII character
func grayToAscii(gray uint8) string {
	index := int(float64(gray) / grayFactor)
	return string(asciiChars[index])
}

// ansiColorCode returns the ANSI escape code for a given RGB color
func ansiColorCode(r, g, b uint8) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// RenderToASCII converts a 2D slice of types.RGBPixels to an ASCII string
func RenderToASCII(pixels [][]types.RGBPixel, resolution int, color *types.RGBPixel) string {
	height := len(pixels)
	if height == 0 {
		return ""
	}
	width := len(pixels[0])
	stepX, stepY := float64(width)/float64(resolution), float64(height)/float64(resolution)
	var ascii strings.Builder
	// Default to white if no color specified
	r, g, b := uint8(255), uint8(255), uint8(255)
	if color != nil {
		r, g, b = color.R, color.G, color.B
	}
	colorCode := ansiColorCode(r, g, b)
	for y := 0.0; y < float64(height-1); y += stepY {
		for x := 0.0; x < float64(width-1); x += stepX {
			i, j := int(math.Round(x)), int(math.Round(y))
			pixel := pixels[j][i]
			gray := rgbToGray(pixel)
			// Convert pixel to colored ASCII character
			ascii.WriteString(colorCode + grayToAscii(gray) + "\033[0m") // Reset color after each character
		}
		ascii.WriteString("\n")
	}
	return ascii.String()
}
