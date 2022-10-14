package main

import (
	httpHandler "camera-emulator/lib/http"
	streamProvider "camera-emulator/lib/stream"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

func main() {
	stream := httpHandler.Handler{
		Next: func() (image.Image, error) {
			return streamProvider.FetchFrame()
		},
		Options: &jpeg.Options{Quality: 80},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
