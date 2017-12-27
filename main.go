package main

import (
	"fmt"
	"image/png"
	"os"
)

const (
	filename    = "out.png"
	complexity  = 1024
	size        = 2048
	workerCount = 4
)

func main() {
	img := mandelbrotWorkers(size, workerCount)

	filename := "out.png"
	f, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Could not create file: %v \n", filename)
	}
	defer f.Close()

	png.Encode(f, img)
}
