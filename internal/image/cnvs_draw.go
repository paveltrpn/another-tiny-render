package image

import (
	alg "another-tiny-render/pkg/algebra"
)

// Set color to pixel at pen_color[x, y].
// Color takes from corrent cnvs.color* fields.
func (cnvs SCanvas) PutPixel(x, y int) {
	if (x >= cnvs.width) || (y >= cnvs.width) || (x <= 0) || (y <= 0) {
		return
	}

	cnvs.data[((x*4)*cnvs.height+y*4)+0] = cnvs.pen_color[0]
	cnvs.data[((x*4)*cnvs.height+y*4)+1] = cnvs.pen_color[1]
	cnvs.data[((x*4)*cnvs.height+y*4)+2] = cnvs.pen_color[2]
}

// Fill square sector of canvas with given color.
// Obviously, equivalent of square brush with side
// length equal to size.
func (cnvs SCanvas) PutSquareBrush(x, y, size int) {
	half_size := size / 2

	for i := -half_size; i < half_size; i++ {
		for j := -half_size; j < half_size; j++ {
			cnvs.PutPixel(x+i, y+j)
		}
	}
}

// Fill circle sector of canvas with given color.
// Equivalent of round brush with radius = "rad".
// Naive implementation - just draw a concentric
// circles from radius zero to "rad" witch center
// in x, y to emulate a filled circle brush.
func (cnvs SCanvas) PutRoundBrush(x, y, rad int) {
	for i := 1; i < rad; i++ {
		cnvs.BrasenhamCircle(x, y, i)
	}
}

// Draw a line at canvas with Brasenham algoritm.
// Coordinates starts at upper left corner of canvas
func (cnvs *SCanvas) BrasenhamLine(xs, ys int, xe, ye int) {
	var (
		signX, signY, err, err2 int
		// now point coords
		np_x int = xs
		np_y int = ys
	)

	dX := int(alg.Fabs(float32(xe - xs)))
	dY := int(alg.Fabs(float32(ye - ys)))

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

func (cnvs *SCanvas) BrasenhamCircle(cx, cy int, rad int) {
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

func (cnvs *SCanvas) DDALine(xs, ys int, xe, ye int) {
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
