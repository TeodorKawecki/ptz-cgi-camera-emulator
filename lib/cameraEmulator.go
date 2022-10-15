package lib

import (
	"camera-emulator/models"
	"image"
	"math"
	"time"

	"github.com/oliamb/cutter"
)

var (
	currentPosition models.CurrentPosition
	targetPosition  models.TargetPosition
)

func init() {
	currentPosition = models.CurrentPosition{
		Pan:  600,
		Tilt: 411,
		Zoom: 200,
	}
}

func FetchFrame() (image.Image, error) {
	img, err := readFrame()

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  img.Bounds().Dx() / 2,
		Height: img.Bounds().Dy() / 2,
		Anchor: image.Point{int(currentPosition.Pan), int(currentPosition.Tilt)},
	})

	imageRgba := imageToRGBA(croppedImg)

	addOverlay(imageRgba)

	return imageRgba, err
}

func SetTargetPosition(tp *models.TargetPosition) {
	targetPosition = *tp
}

func UpdatePosition() {
	for {
		calculateCurrentPosition()
		time.Sleep(100 * time.Microsecond)
	}
}

func calculateCurrentPosition() {
	xDiff := float64(targetPosition.Pan) - currentPosition.Pan
	yDiff := float64(targetPosition.Tilt) - currentPosition.Tilt

	if xDiff < 0 && math.Abs(xDiff) > 0.05 {
		currentPosition.Pan -= 0.05
	} else if math.Abs(xDiff) > 0.05 {
		currentPosition.Pan += 0.05
	}

	if yDiff < 0 && math.Abs(yDiff) > 0.05 {
		currentPosition.Tilt -= 0.05
	} else if math.Abs(yDiff) > 0.05 {
		currentPosition.Tilt += 0.05
	}
}
