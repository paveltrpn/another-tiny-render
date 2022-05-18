package image

import (
	"errors"
	"math"

	alg "tiny-render-go/pkg/algebra_go"
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

func (cnvs *Canvas_s) SetPenColor(r, g, b uint8) {
	cnvs.color_r = r
	cnvs.color_b = b
	cnvs.color_g = g
}

// Set color to pixel at cnvs.data[x, y].
// Color takes from corrent cnvs.color* fields.
func (cnvs Canvas_s) PutPixel(x, y int) {
	if (x >= cnvs.width) || (y >= cnvs.width) {
		return
	}

	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+0] = cnvs.color_r
	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+1] = cnvs.color_g
	cnvs.data[((x*cnvs.bpp)*cnvs.height+y*cnvs.bpp)+2] = cnvs.color_b
}

// Fill square sector of canvas with given color.
// Obviously, equivalent of square brush with side
// length equal to size.
func (cnvs Canvas_s) PutSquareBrush(x, y, size int) {
	half_size := size / 2

	for i := -half_size; i < half_size; i++ {
		for j := -half_size; j < half_size; j++ {
			cnvs.PutPixel(x+i, y+j)
		}
	}
}

// Fill circle sector of canvas with given color.
// Equivalent of round brush with radius = rad.
// NOT WORK PROPER BY NOW!!!!
func (cnvs Canvas_s) PutRoundBrush(x, y, rad int) {
	half_rad := rad / 2

	for i := -half_rad; i < half_rad; i++ {
		for j := i + half_rad; j < half_rad; j++ {
			cnvs.PutPixel(x+i, y+j)
		}
	}
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

	cnvs.PutRoundBrush(xe, ye, 20)

	for (np_x != xe) || (np_y != ye) {
		cnvs.PutRoundBrush(np_x, np_y, 20)

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

func (cnvs *Canvas_s) BrasenhamCircle(cx, cy int, rad int) {
	var (
		x     int = 0
		y     int = rad
		delta int = 1 - 2*rad
		error int = 0
	)

	for y >= 0 {
		cnvs.PutPixel(cx+x, cy+y)
		cnvs.PutPixel(cx+x, cy-y)
		cnvs.PutPixel(cx-x, cy+y)
		cnvs.PutPixel(cx-x, cy-y)

		error = 2*(delta+y) - 1

		if (delta < 0) && (error <= 0) {
			x = x + 1
			delta += 2*x + 1
			continue
		}

		if (delta > 0) && (error > 0) {
			y = y - 1
			delta -= 2*y + 1
			continue
		}

		x += 1
		y -= 1
		delta += 2 * (x - y)
	}
}

func (cnvs *Canvas_s) DDALine(xs, ys int, xe, ye int) {
	var (
		dx                      float32 = float32(xe - xs)
		dy                      float32 = float32(ye - ys)
		steps, Xinc, Yinc, X, Y float32
	)

	if alg.Fabs(dx) > alg.Fabs(dy) {
		steps = alg.Fabs(dx)
	} else {
		steps = alg.Fabs(dy)
	}

	Xinc = dx / steps
	Yinc = dy / steps

	X = float32(xs)
	Y = float32(ys)

	for i := 0; i <= int(steps); i++ {
		cnvs.PutPixel(int(X), int(Y))
		X += Xinc
		Y += Yinc
	}
}
