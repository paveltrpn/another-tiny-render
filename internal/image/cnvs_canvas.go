package image

import (
	"unsafe"
)

const (
	CNVS_RGB  = 3
	CNVS_RGBA = 4
)

type SRGBABitmap struct {
	data   []uint8
	width  int
	height int
}

type SCanvas struct {
	SRGBABitmap

	// color components that used for PutPixel()
	// "pen" color, default - white defined in BuildCanvas()
	pen_color [3]uint8
}

type SRGBA struct {
	r, g, b, a uint8
}

func buildNilCanvas() SCanvas {
	return SCanvas{SRGBABitmap: SRGBABitmap{data: nil,
		width:  0,
		height: 0},
		pen_color: [...]uint8{255, 255, 255}}
}

func BuildEmptyCanvas(xs, ys int) (SCanvas, error) {
	return SCanvas{SRGBABitmap: SRGBABitmap{data: make([]uint8, xs*ys*4),
		width:  xs,
		height: ys},
		pen_color: [...]uint8{255, 255, 255}}, nil
}

// Directly cast pointer to first element of data slice
// to unsafe.Pointer to avoid of calling gl.Ptr() function
// which involves some runtime type reflection operations.
// It works because i'm sure that &cnvs.data[0] is
// allways addressable.
// Not very safe, indeed.
func (cnvs SCanvas) GetDataUnsafePtr() unsafe.Pointer {
	return unsafe.Pointer(&cnvs.data[0])
}

// Simple return of pointer to first element of data slice
// to pass they to gl.Ptr() function in further.
// More safe way.
func (cnvs SCanvas) GetDataPtr() *uint8 {
	return &cnvs.data[0]
}

func (cnvs SCanvas) GetData() []uint8 {
	return cnvs.data
}

func (cnvs SCanvas) GetWidth() int {
	return cnvs.width
}

func (cnvs SCanvas) GetHeight() int {
	return cnvs.height
}

func (cnvs *SCanvas) SetPenColor(r, g, b uint8) {
	cnvs.pen_color[0] = r
	cnvs.pen_color[1] = b
	cnvs.pen_color[2] = g
}

func (cnvs *SCanvas) getPixelIndex(row, col int) int {
	return (row*4)*cnvs.height + col*4
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
