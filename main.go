package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
)

const (
	filename    = "out.png"
	size        = 1024
	workerCount = 4
)

// func main() {
// 	img := mandelbrotWorkers(size, workerCount)

// 	filename := "out.png"
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		fmt.Printf("Could not create file: %v \n", filename)
// 	}
// 	defer f.Close()

// 	png.Encode(f, img)
// }

func main() {
	http.Handle("/", http.HandlerFunc(mandelbrotServer))

	fmt.Println("Listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mandelbrotServer(h http.ResponseWriter, r *http.Request) {
	img := mandelbrotWorkers(size, workerCount)

	h.Header().Add("Content-Type", "image/png")
	png.Encode(h, img)
}
