package main

import (
	"image"
	"image/color"
)

func mandelbrotSequential(sizeInPx int) image.Image {
	img := image.NewGray(image.Rect(0, 0, sizeInPx, sizeInPx))

	for i := 0; i < sizeInPx; i++ {
		for j := 0; j < sizeInPx; j++ {
			img.Set(i, j, getColour(i, j, sizeInPx, sizeInPx))
		}
	}

	return img
}

func mandelbrotPerRow(sizeInPx int) image.Image {
	img := image.NewGray(image.Rect(0, 0, sizeInPx, sizeInPx))

	for i := 0; i < sizeInPx; i++ {
		go func(x int) {
			for j := 0; j < sizeInPx; j++ {
				img.Set(x, j, getColour(x, j, sizeInPx, sizeInPx))
			}
		}(i)
	}

	return img
}

func normalise(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}

// Get the colour of the Mandelbrot set for the given pre-scaled pixel co-ordinates
func getColour(px, py, width, height int) color.Color {
	x0 := normalise(px, width, -1.0, 2)
	y0 := normalise(py, height, -1.0, 1.0)
	x := 0.0
	y := 0.0
	max := 1000
	for i := 0; x*x+y*y < complexity && i < max; i++ {
		x, y = x*x-y*y+x0, 2*x*y+y0
	}
	return color.Gray{uint8(x)}
}
