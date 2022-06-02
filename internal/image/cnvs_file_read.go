package image

import (
	"image/jpeg"
	"os"
)

func (cnvs *SCanvas) LoadFromJpegFile(fname string) {
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

	cnvs.width = sz.Max.X - sz.Min.X
	cnvs.height = sz.Max.Y - sz.Min.Y
	cnvs.depth = 3

	cnvs.data = make([]uint8, cnvs.width*cnvs.height*3)

	idx := 0
	for y := sz.Min.Y; y < sz.Max.Y; y++ {
		for x := sz.Min.X; x < sz.Max.X; x++ {
			r, g, b, _ := imageData.At(x, y).RGBA()
			cnvs.data[idx], cnvs.data[idx+1], cnvs.data[idx+2] = uint8(r), uint8(g), uint8(b)
			idx += 3
		}
	}
}
