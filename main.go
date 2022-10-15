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
	go streamProvider.UpdatePosition()

	stream := httpHandler.Handler{
		Next: func() (image.Image, error) {
			return streamProvider.FetchFrame()
		},
		Options: &jpeg.Options{Quality: 80},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)

	mux.HandleFunc("/setPreset", func(w http.ResponseWriter, r *http.Request) {
		targetPosition := &streamProvider.TargetPosition{}
		httpHandler.HandleSetPreset(w, r, targetPosition)
		streamProvider.SetTargetPosition(targetPosition)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
