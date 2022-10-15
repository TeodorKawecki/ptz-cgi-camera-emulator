package lib

import (
	"image"
	"image/draw"
	"os"
)

const fileLocation = "./media/lena.jpg"

func readFrame() (image.Image, error) {
	f, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)

	return img, err
}

func imageToRGBA(src image.Image) *image.RGBA {

	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)

	return dst
}
