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
		Width:  1280,
		Height: 720,
		Anchor: image.Point{int(currentPosition.Pan), int(currentPosition.Tilt)},
	})

	return croppedImg, err
}

func zoomImage(img *image.Image) (image.Image, error) {
	//dst := image.NewRGBA(image.Rect(0, 0, (*img).Bounds().Max.X/2, (*img).Bounds().Max.Y/2))

	//draw.NearestNeighbor.Scale(dst, dst.Rect, *img, (*img).Bounds(), draw.Over, nil)

	m := resize.Resize(1000, 0, *img, resize.Lanczos3)

	return m, nil
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
