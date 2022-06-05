package image

import (
	"errors"
)

const (
	CNVS_RGB  = 3
	CNVS_RGBA = 4
)

type SCanvas struct {
	data   []uint8
	width  int
	height int

	depth int

	// color components that used for PutPixel()
	// "pen" color, default - white defined in BuildCanvas()
	pen_color [3]uint8
}

func BuildCanvas(xs, ys, bpp int) (SCanvas, error) {
	// check only depth parameter,
	// maybe check canvas size bounds?
	if (bpp < 3) || (bpp > 4) {
		return SCanvas{data: nil,
			width:  0,
			height: 0,
			depth:  0}, errors.New("BuildCanvas(): Error! Wrong canvas pixel depth value")
	}

	byteArray := make([]uint8, xs*ys*bpp)

	return SCanvas{data: byteArray,
		width:     xs,
		height:    ys,
		depth:     bpp,
		pen_color: [...]uint8{255, 255, 255}}, nil
}

func (cnvs SCanvas) GetDataPtr() []uint8 {
	return cnvs.data
}

func (cnvs SCanvas) GetWidth() int {
	return cnvs.width
}

func (cnvs SCanvas) GetHeight() int {
	return cnvs.height
}

func (cnvs SCanvas) GetDepth() int {
	return cnvs.depth
}

func (cnvs *SCanvas) SetPenColor(r, g, b uint8) {
	cnvs.pen_color[0] = r
	cnvs.pen_color[1] = b
	cnvs.pen_color[2] = g
}

func (cnvs *SCanvas) getPixelIndex(row, col int) int {
	return (row*cnvs.depth)*cnvs.height + col*cnvs.depth
}

func (cnvs *SCanvas) getPixelRVal(row, col int) uint8 {
	return cnvs.data[cnvs.getPixelIndex(row, col)+0]
}

func (cnvs *SCanvas) getPixelGVal(row, col int) uint8 {
	return cnvs.data[cnvs.getPixelIndex(row, col)+1]
}

func (cnvs *SCanvas) getPixelBVal(row, col int) uint8 {
	return cnvs.data[cnvs.getPixelIndex(row, col)+2]
}

func (cnvs *SCanvas) getPixelAVal(row, col int) uint8 {
	return cnvs.data[cnvs.getPixelIndex(row, col)+3]
}
