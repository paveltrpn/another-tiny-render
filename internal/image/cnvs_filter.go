package image

func multWithClamp[T float32 | float64](f T, val uint8) uint8 {
	tmp := T(val) * f
	if tmp > 255 {
		tmp = 255
	} else if tmp < 0.0 {
		tmp = 0
	}
	return uint8(tmp)
}

func (cnvs *SCanvas) MultScalar(factor float32) {
	var (
		pixel [3]uint8
	)

	for row := 0; row < cnvs.width; row++ {
		for col := 0; col < cnvs.height; col++ {
			pixel[0] = cnvs.getPixelRVal(row, col)
			pixel[1] = cnvs.getPixelGVal(row, col)
			pixel[2] = cnvs.getPixelBVal(row, col)

			cnvs.data[((row*3)*cnvs.height+col*3)+0] = multWithClamp(factor, pixel[0])
			cnvs.data[((row*3)*cnvs.height+col*3)+1] = multWithClamp(factor, pixel[1])
			cnvs.data[((row*3)*cnvs.height+col*3)+2] = multWithClamp(factor, pixel[2])
		}
	}
}

func (cnvs *SCanvas) MultPerComponent(fR, fG, fB float32) {
	var (
		pixel [3]uint8
	)

	for row := 0; row < cnvs.width; row++ {
		for col := 0; col < cnvs.height; col++ {
			pixel[0] = cnvs.getPixelRVal(row, col)
			pixel[1] = cnvs.getPixelGVal(row, col)
			pixel[2] = cnvs.getPixelBVal(row, col)

			cnvs.data[((row*3)*cnvs.height+col*3)+0] = multWithClamp(fR, pixel[0])
			cnvs.data[((row*3)*cnvs.height+col*3)+1] = multWithClamp(fG, pixel[1])
			cnvs.data[((row*3)*cnvs.height+col*3)+2] = multWithClamp(fB, pixel[2])
		}
	}
}
