package image

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func BuildCanvasFromFile(fname string) (SCanvas, error) {
	imagefile, err := os.Open(fname)
	if err != nil {
		return buildNilCanvas(), err
	}
	defer imagefile.Close()

	var (
		model     color.Model
		imageData image.Image
		cnvs      SCanvas
		idx       int
	)

	suffix := filepath.Ext(fname)

	switch suffix {
	case ".jpg":
		imageData, err = jpeg.Decode(imagefile)
		if err != nil {
			return buildNilCanvas(), err
		}
		model = color.NRGBAModel
	case ".png":
		imageData, err = png.Decode(imagefile)
		if err != nil {
			return buildNilCanvas(), err
		}
		model = color.RGBAModel
	default:
		return buildNilCanvas(), fmt.Errorf("BuildCanvasFromFile(): Error! Wrong file type with file - %v", fname)
	}

	// Convert image.Image interface to []uint8 rgba slice
	sz := imageData.Bounds()

	cnvs.width = sz.Max.X - sz.Min.X
	cnvs.height = sz.Max.Y - sz.Min.Y

	cnvs.data = make([]uint8, cnvs.width*cnvs.height*4)

	for y := sz.Min.Y; y < sz.Max.Y; y++ {
		for x := sz.Min.X; x < sz.Max.X; x++ {
			rgba := imageData.At(x, y)
			// Convert an alpha-premultiplied colors to non-alpha-premultiplied 32-bit color
			nrgba := model.Convert(rgba)
			r, g, b, a := nrgba.RGBA()

			cnvs.data[idx], cnvs.data[idx+1], cnvs.data[idx+2], cnvs.data[idx+3] = uint8(r), uint8(g), uint8(b), uint8(a)
			idx += 4
		}
	}

	return cnvs, nil
}
