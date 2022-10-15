// Package mjpeg implements mjpeg streaming handlers with a simple API.
package streamProvider

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"os"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/oliamb/cutter"
)

const fileLocation = "/home/teodorek/Github/camera-emulator/lena.jpg"

type CurrentPosition struct {
	Pan  float64
	Tilt float64

	Zoom int
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
		Zoom: 200,
	}
}

func FetchFrame() (image.Image, error) {
	img, err := ReadImage()

	croppedImg, err := cutter.Crop(img, cutter.Config{
		Width:  img.Bounds().Dx() / 2,
		Height: img.Bounds().Dy() / 2,
		Anchor: image.Point{int(currentPosition.Pan), int(currentPosition.Tilt)},
	})

	imageRgba := imageToRGBA(croppedImg)

	addLabel(imageRgba)

	return imageRgba, err
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

func imageToRGBA(src image.Image) *image.RGBA {

	// No conversion needed if image is an *image.RGBA.
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	// Use the image/draw package to convert to *image.RGBA.
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Src)
	return dst
}

func addLabel(img *image.RGBA) {
	col := color.RGBA{200, 100, 0, 255}

	y := img.Rect.Size().Y - 50
	x := 10

	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	panLabel := fmt.Sprintf("Pan : %d", int(currentPosition.Pan))
	tiltLabel := fmt.Sprintf("Tilt: %d", int(currentPosition.Tilt))
	zoomLabel := fmt.Sprintf("Zoom: %d", int(currentPosition.Zoom))

	d.DrawString(panLabel)
	d.Dot = fixed.Point26_6{fixed.I(x), fixed.I(y + 10)}
	d.DrawString(tiltLabel)
	d.Dot = fixed.Point26_6{fixed.I(x), fixed.I(y + 20)}
	d.DrawString(zoomLabel)
}
