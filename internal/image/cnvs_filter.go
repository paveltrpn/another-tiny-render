package image

import (
	alg "another-tiny-render/pkg/algebra_go"
	fgblur "another-tiny-render/pkg/fast_gauss_blur"
	"math"
)

// Mult float and uint8 and clamp result to byte range [0...255]
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

			cnvs.data[((row*4)*cnvs.height+col*3)+4] = multWithClamp(factor, pixel[0])
			cnvs.data[((row*4)*cnvs.height+col*3)+4] = multWithClamp(factor, pixel[1])
			cnvs.data[((row*4)*cnvs.height+col*3)+4] = multWithClamp(factor, pixel[2])
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

			cnvs.data[((row*4)*cnvs.height+col*4)+0] = multWithClamp(fR, pixel[0])
			cnvs.data[((row*4)*cnvs.height+col*4)+1] = multWithClamp(fG, pixel[1])
			cnvs.data[((row*4)*cnvs.height+col*4)+2] = multWithClamp(fB, pixel[2])
		}
	}
}

// Calculate new pixel color as average sum of surrounding
// pixels (covered by square kernel) multiplyed at corresponding
// kernel numbers.
func (cnvs *SCanvas) MultByKernel(krnl []float32) {

}

func (cnvs SCanvas) ExtractSurround5by5(pixId int) *[25]int {
	const (
		krnlSize     = 5
		halfKrnlSize = krnlSize / 2
	)

	var (
		r, c, i  int
		row, cmn int
	)

	rows := cnvs.height

	rt := new([krnlSize * krnlSize]int)

	rowById := pixId / rows
	cmnById := pixId - rowById*rows

	for r = -halfKrnlSize; r < halfKrnlSize+1; r++ {
		for c = -halfKrnlSize; c < halfKrnlSize+1; c++ {
			row = rowById + r
			cmn = cmnById + c
			if (row < 0) || (cmn < 0) {
				rt[i] = -1
			} else {
				rt[i] = row*rows + cmn
			}
			i++
		}
	}

	return rt
}

func ExtractSurround(pixId int32) *[25]int32 {
	const (
		rows = 7 // 0...6
		clms = 7 // 0...6

		krnlSize     = 5 // Only odd numbers
		halfKrnlSize = krnlSize / 2
	)

	var (
		r, c     int32
		row, cmn int32
	)

	rt := new([krnlSize * krnlSize]int32)

	rowById := pixId / rows
	cmnById := pixId - rowById*rows

	i := 0
	for r = -halfKrnlSize; r < halfKrnlSize+1; r++ {
		for c = -halfKrnlSize; c < halfKrnlSize+1; c++ {
			row = rowById + r
			cmn = cmnById + c
			if (row < 0) || (cmn < 0) {
				rt[i] = -1
			} else {
				rt[i] = row*rows + cmn
			}
			i++
		}
	}

	return rt
}

func genGausFilterKernel() [5][5]float32 {
	var (
		// initialising standard deviation to 1.0
		sigma float32 = 1.0
		r     float32
		s     float32 = 2.0 * sigma * sigma
		// sum is for normalization
		sum    float32 = 0.0
		kernel [5][5]float32
	)

	// generating 5x5 kernel
	for x := -2; x <= 2; x++ {
		for y := -2; y <= 2; y++ {
			r = alg.Sqrtf(float32(x*x + y*y))
			kernel[x+2][y+2] = (float32(alg.Exp((-(r * r) / s))) / (math.Pi * s))
			sum += kernel[x+2][y+2]
		}
	}

	// normalising the Kernel
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			kernel[i][j] /= sum
		}
	}

	return kernel
}

func FastGaussianBlurRGB(in []uint8, out []uint8, w int, h int, c int, sigma float32) {
	fgblur.Fast_gaussian_blur_rgb(in, out, w, h, c, sigma)
}
