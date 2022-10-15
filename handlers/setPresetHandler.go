package handlers

import (
	models "camera-emulator/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleSetPreset(w http.ResponseWriter, r *http.Request, tp *models.TargetPosition) {
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
