package lib

import (
	"camera-emulator/models"
	"image"
	"math"
	"time"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

var (
	currentPosition models.CurrentPosition
	targetPosition  models.TargetPosition
	camParams       models.CamParams
)

func init() {
	currentPosition = models.CurrentPosition{
		Pan:  0,
		Tilt: 0,
		Zoom: 0,
	}

	camParams = models.CamParams{
		Width:  800,
		Height: 600,
	}
}

func FetchFrame() (image.Image, error) {
	img, err := readFrame()

	img, err = zoomImage(&img)
	img, err = cropImage(&img)
	imgRgba := imageToRGBA(img)

	addOverlay(imgRgba)

	return imgRgba, err
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

func cropImage(img *image.Image) (image.Image, error) {
	croppedImg, err := cutter.Crop(*img, cutter.Config{
		Width:  camParams.Width,
		Height: camParams.Height,
		Anchor: image.Point{int(currentPosition.Pan), int(currentPosition.Tilt)},
	})

	return croppedImg, err
}

func zoomImage(img *image.Image) (image.Image, error) {
	multiplier := float64(currentPosition.Zoom)/100.0 + 1
	newWidth := uint(float64(camParams.Width) * multiplier)
	newHigh := uint(float64(camParams.Height) * multiplier)

	m := resize.Resize(newWidth, newHigh, *img, resize.Lanczos3)

	return m, nil
}

func calculateCurrentPosition() {
	panDiff := float64(targetPosition.Pan) - currentPosition.Pan
	tiltDiff := float64(targetPosition.Tilt) - currentPosition.Tilt
	zoomDiff := float64(targetPosition.Zoom) - currentPosition.Zoom

	if panDiff < 0 && math.Abs(panDiff) > 0.05 {
		currentPosition.Pan -= 0.05
	} else if math.Abs(panDiff) > 0.05 {
		currentPosition.Pan += 0.05
	}

	if tiltDiff < 0 && math.Abs(tiltDiff) > 0.05 {
		currentPosition.Tilt -= 0.05
	} else if math.Abs(tiltDiff) > 0.05 {
		currentPosition.Tilt += 0.05
	}

	if zoomDiff < 0 && math.Abs(zoomDiff) > 0.05 {
		currentPosition.Zoom -= 0.05
	} else if math.Abs(zoomDiff) > 0.05 {
		currentPosition.Zoom += 0.05
	}
}
