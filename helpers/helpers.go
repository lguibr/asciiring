package helpers

import (
	"image"
	"image/png"
	"os"
	"os/exec"
	"runtime"
)

func SaveImageToFile(img *image.RGBA, path string) error {
	// Create a new file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the image to the file
	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func ClearScreen() {
	var cmd *exec.Cmd
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default: // Unix-like system
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
