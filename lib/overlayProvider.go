package lib

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"
)

func addOverlay(img *image.RGBA) {
	col := color.RGBA{255, 0, 0, 255}

	y := img.Rect.Size().Y - 50
	x := 10

	point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Bold8x16,
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
