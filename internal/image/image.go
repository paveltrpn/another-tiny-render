package image

import (
	"image"
	"image/jpeg"
	"os"
)

func loadImage(fname string) image.Image {
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

	return imageData
}

func imageToRGB(img image.Image) []uint8 {
	sz := img.Bounds()
	raw := make([]uint8, (sz.Max.X-sz.Min.X)*(sz.Max.Y-sz.Min.Y)*3)
	idx := 0
	for y := sz.Min.Y; y < sz.Max.Y; y++ {
		for x := sz.Min.X; x < sz.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			raw[idx], raw[idx+1], raw[idx+2] = uint8(r), uint8(g), uint8(b)
			idx += 3
		}
	}
	return raw
}
