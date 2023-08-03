package render

import (
	"image"
	"image/color"
	"image/draw"
	"strings"

	"github.com/lguibr/asciiring/types"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func CreateTextImage(text string, lineHeight, charWidth int) *image.RGBA {
	lines := strings.Split(text, "\n")

	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// adding padding to the width and height
	padding := 10
	width := maxWidth*charWidth + padding*2
	height := len(lines)*lineHeight + padding*2
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	for i, line := range lines {
		x := ((maxWidth-len(line))/2)*charWidth + charWidth/2 + padding // adjust x position considering padding
		y := ((i + 1) * lineHeight) + padding                           // adjust y position considering padding

		d := &font.Drawer{
			Dst:  img,
			Src:  &image.Uniform{color.Black},
			Face: basicfont.Face7x13,
			Dot:  fixed.P(x, y),
		}

		d.DrawString(line)
	}

	return img
}

func ImageToRGBArray(img image.Image) [][]types.RGBPixel {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	pixels := make([][]types.RGBPixel, height)
	for y := 0; y < height; y++ {
		pixels[y] = make([]types.RGBPixel, width)
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Converting from range [0, 65535] to [0, 255]
			pixels[y][x] = types.RGBPixel{
				R: uint8(r / 257),
				G: uint8(g / 257),
				B: uint8(b / 257),
			}
		}
	}
	return pixels
}

func TextToRGB(text string, lineHeight, charWidth int) [][]types.RGBPixel {
	image := CreateTextImage(text, lineHeight, charWidth)
	return ImageToRGBArray(image)
}
