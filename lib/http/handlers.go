package httpHandler

import (
	streamProvider "camera-emulator/lib/stream"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

type Handler struct {
	Next    func() (image.Image, error)
	Options *jpeg.Options
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "multipart/x-mixed-replace; boundary=frame")
	boundary := "\r\n--frame\r\nContent-Type: image/jpeg\r\n\r\n"
	for {
		img, err := h.Next()
		if err != nil {
			return
		}

		n, err := io.WriteString(w, boundary)
		if err != nil || n != len(boundary) {
			return
		}

		err = jpeg.Encode(w, img, h.Options)
		if err != nil {
			return
		}

		n, err = io.WriteString(w, "\r\n")
		if err != nil || n != 2 {
			return
		}
	}
}

func HandleSetPreset(w http.ResponseWriter, r *http.Request, tp *streamProvider.TargetPosition) {
	if r.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}

	err := json.NewDecoder(r.Body).Decode(tp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Print(tp)
}
