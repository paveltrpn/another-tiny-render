package image

import (
	"image/jpeg"
	"os"
)

type SImage struct {
	data []uint8

	width  int
	height int
	depth  int
}

func (img SImage) GetWidth() int {
	return img.width
}

func (img SImage) GetHeight() int {
	return img.height
}

func (img SImage) GetDepth() int {
	return img.depth
}

func (img SImage) GetDataPtr() *uint8 {
	return &img.data[0]
}

func (img *SImage) LoadFromJpegFile(fname string) {
	// Load image file as image.Image interface
	imagefile, err := os.Open(fname)
	if err != nil {
		println(err)
		panic(0)
	}
	defer imagefile.Close()

	imageData, err := jpeg.Decode(imagefile)
	if err != nil {
		println(err)
		panic(0)
	}

	// Convert image.Image interface to []uint8 rgb slice
	sz := imageData.Bounds()

	img.width = sz.Max.X - sz.Min.X
	img.height = sz.Max.Y - sz.Min.Y
	img.depth = 3

	img.data = make([]uint8, img.width*img.height*3)

	idx := 0
	for y := sz.Min.Y; y < sz.Max.Y; y++ {
		for x := sz.Min.X; x < sz.Max.X; x++ {
			r, g, b, _ := imageData.At(x, y).RGBA()
			img.data[idx], img.data[idx+1], img.data[idx+2] = uint8(r), uint8(g), uint8(b)
			idx += 3
		}
	}
}
