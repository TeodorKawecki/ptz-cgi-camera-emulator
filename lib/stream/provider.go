// Package mjpeg implements mjpeg streaming handlers with a simple API.
package streamProvider

import (
	"image"
	"os"

	"github.com/oliamb/cutter"
)

const fileLocation = "/home/teodorek/Github/camera-emulator/lena.jpg"

type CamParams struct {
	xAxis int
	yAxis int

	zoom int
}

var (
	camParams CamParams
)

func init() {
	camParams = CamParams{
		xAxis: 10,
		yAxis: 10,
		zoom:  0,
	}
}

func FetchFrame() (image.Image, error) {
	img, err := ReadImage()

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  img.Bounds().Dx() - camParams.zoom,
		Height: img.Bounds().Dy() - camParams.zoom,
		Anchor: image.Point{camParams.xAxis, camParams.yAxis},
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

func SetXAxis(xAxis int) {
	camParams.xAxis = xAxis
}

func SetYAxis(yAxis int) {
	camParams.xAxis = yAxis
}

func SetZoom(zoom int) {
	camParams.zoom = zoom
}
