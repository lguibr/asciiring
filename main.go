package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/lguibr/asciiring/helpers"
	"github.com/lguibr/asciiring/render"
	"github.com/lguibr/asciiring/types"
)

type Frame struct {
	Text       string `json:"text"`
	FontSize   int    `json:"fontSize"`
	LineHeight int    `json:"lineHeight"`
	CharWidth  int    `json:"charWidth"`
	Duration   int    `json:"duration"` // in seconds
}

func main() {
	// Load and parse the frames from a JSON file
	data, err := ioutil.ReadFile("frames.json")
	if err != nil {
		log.Fatal(err)
	}

	var frames []Frame
	if err := json.Unmarshal(data, &frames); err != nil {
		log.Fatal(err)
	}

	// Loop through the frames
	for _, frame := range frames {
		// Create the ASCII art and display it
		img := render.CreateTextImage(frame.Text, frame.LineHeight, frame.CharWidth)
		helpers.SaveImageToFile(img, "text.png")
		pixels := render.ImageToRGBArray(img)
		color := types.RGBPixel{R: 255, G: 255, B: 255} // Red color

		ascii := render.RenderToASCII(pixels, 100, &color)
		fmt.Println(ascii)

		// Pause for the duration of the frame
		time.Sleep(time.Duration(frame.Duration) * time.Second)
		helpers.ClearScreen()

	}
}
