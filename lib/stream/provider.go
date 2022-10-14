// Package mjpeg implements mjpeg streaming handlers with a simple API.
package streamProvider

import (
	"image"
	"os"

	"github.com/oliamb/cutter"
)

type CamParams struct {
	xAxis int
	yAxis int

	zoom int
}

func FetchFrame() (image.Image, error) {
	params := CamParams{
		xAxis: 10,
		yAxis: 10,
		zoom:  0,
	}

	f, err := os.Open("/home/teodorek/Github/camera-emulator/lena.jpg")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  img.Bounds().Dx() - params.zoom,
		Height: img.Bounds().Dy() - params.zoom,
		Anchor: image.Point{params.xAxis, params.yAxis},
	})

	return croppedImg, err
}
