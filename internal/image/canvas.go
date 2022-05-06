package image

import (
	"errors"
	"math"
)

type Canvas_s struct {
	data   []uint8
	width  int
	height int

	bpp int

	// color components that used for PutPixel()
	// "pen" color, default - white defined in BuildCanvas()
	color_r uint8
	color_g uint8
	color_b uint8
}

func BuildCanvas(xs, ys, bpp int) (Canvas_s, error) {
	// check only depth parameter,
	// maybe check canvas size bounds?
	if (bpp < 3) || (bpp > 4) {
		return Canvas_s{data: nil,
			width:  0,
			height: 0,
			bpp:    0}, errors.New("BuildCanvas(): Error! cnvs_bpp can't be less than zero")
	}

	byteArray := make([]uint8, xs*ys*bpp)

	return Canvas_s{data: byteArray,
		width:   xs,
		height:  ys,
		bpp:     bpp,
		color_r: 255,
		color_g: 255,
		color_b: 255}, nil
}

func (cnvs Canvas_s) GetDataPtr() []uint8 {
	return cnvs.data
}

func (cnvs Canvas_s) GetWidth() int {
	return cnvs.width
}

func (cnvs Canvas_s) GetHeight() int {
	return cnvs.height
}

func (cnvs Canvas_s) GetBpp() int {
	return cnvs.bpp
}

// Set color at to pixel at cnvs.data[x, y].
// Color takes from corrent cnvs.color* fields.
func (cnvs Canvas_s) PutPixel(x, y int) {
	if (x >= cnvs.width) || (y >= cnvs.width) {
		return
	}

	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+0] = cnvs.color_r
	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+1] = cnvs.color_g
	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+2] = cnvs.color_b
}

// Draw a line at canvas with Brasenham algoritm.
// Coordinates starts at upper left corner of canvas
func (cnvs *Canvas_s) BrasenhamLine(xs, ys int, xe, ye int) {
	var (
		signX, signY, err, err2 int
		// now point coords
		np_x int = xs
		np_y int = ys
	)

	dX := int(math.Abs((float64)(xe - xs)))
	dY := int(math.Abs((float64)(ye - ys)))

	err = dX - dY

	if xs < xe {
		signX = 1
	} else {
		signX = -1
	}

	if ys < ye {
		signY = 1
	} else {
		signY = -1
	}

	cnvs.PutPixel(xe, ye)

	for (np_x != xe) || (np_y != ye) {
		cnvs.PutPixel(np_x, np_y)

		err2 = err * 2

		if err2 > -dY {
			err -= dY
			np_x += signX
		}

		if err2 < dX {
			err += dX
			np_y += signY
		}
	}
}
