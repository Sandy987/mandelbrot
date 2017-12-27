package main

import (
	"fmt"
	"image/png"
	"os"
)

const (
	filename   = "out.png"
	complexity = 2056 * 2
	size       = 4096 * 2
)

func main() {
	img := mandelbrotPerRow(size)

	filename := "out.png"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Could not create file: %v \n", filename)
	}
	defer f.Close()

	png.Encode(f, img)
}
