package main

import (
	mjpeg "camera-emulator/lib"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	stream := mjpeg.Handler{
		Next: func() (image.Image, error) {
			img := image.NewGray(image.Rect(0, 0, 100, 100))
			for i := 0; i < 100; i++ {
				for j := 0; j < 100; j++ {
					n := rand.Intn(256)
					gray := color.Gray{uint8(n)}
					img.SetGray(i, j, gray)
				}
			}
			return img, nil
		},
		Options: &jpeg.Options{Quality: 80},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
