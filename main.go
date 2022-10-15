package main

import (
	handlers "camera-emulator/handlers"
	lib "camera-emulator/lib"
	models "camera-emulator/models"
	"image"
	"image/jpeg"
	"log"
	"net/http"
)

func main() {
	go lib.UpdatePosition()

	stream := handlers.Handler{
		Next: func() (image.Image, error) {
			return lib.FetchFrame()
		},
		Options: &jpeg.Options{Quality: 100},
	}

	mux := http.NewServeMux()
	mux.Handle("/stream", stream)

	mux.HandleFunc("/setPreset", func(w http.ResponseWriter, r *http.Request) {
		targetPosition := &models.TargetPosition{}
		handlers.HandleSetPreset(w, r, targetPosition)
		lib.SetTargetPosition(targetPosition)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
