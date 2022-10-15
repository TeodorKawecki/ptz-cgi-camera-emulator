// Package mjpeg implements mjpeg streaming handlers with a simple API.
package streamProvider

import (
	"image"
	"math"
	"os"
	"time"

	"github.com/oliamb/cutter"
)

const fileLocation = "/home/teodorek/Github/camera-emulator/lena.jpg"

type CurrentPosition struct {
	Pan  float64
	Tilt float64

	zoom int
}

type TargetPosition struct {
	Pan  int
	Tilt int

	Zoom int
}

var (
	currentPosition CurrentPosition
	targetPosition  TargetPosition
)

func init() {
	currentPosition = CurrentPosition{
		Pan:  600,
		Tilt: 411,
		zoom: 200,
	}
}

func FetchFrame() (image.Image, error) {
	img, err := ReadImage()

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  img.Bounds().Dx() - currentPosition.zoom,
		Height: img.Bounds().Dy() - currentPosition.zoom,
		Anchor: image.Point{int(currentPosition.Pan), int(currentPosition.Tilt)},
	})

	return croppedImg, err
}

func ReadImage() (image.Image, error) {
	f, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	return img, err
}

func SetTargetPosition(tp *TargetPosition) {
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
