package fastgaussblur

import (
	alg "another-tiny-render/pkg/algebra"
)

// Golang interpretation of fast Gaussian blur algorithm by Ivan Kutskir.
// Original code was writen in C++. I keep author function naming style, all
// comments and functions descriptions. I only add one exported function - Fast_gaussian_blur_rgb()
// for interface purpose to another part of module.
//!
//! \file blur.cpp
//! \author Basile Fraboni
//! \date 2017
//!
//! \brief The software is a C++ implementation of a fast
//! Gaussian blur algorithm by Ivan Kutskir. For further details
//! please refer to :
//! http://blog.ivank.net/fastest-gaussian-blur.html
//!
//! Unsigned char version
//!

// !
// ! \fn void std_to_box(float boxes[], float sigma, int n)
// !
// ! \brief this function converts the standard deviation of
// ! Gaussian blur into dimensions of boxes for box blur. For
// ! further details please refer to :
// ! https://www.peterkovesi.com/matlabfns/#integral
// ! https://www.peterkovesi.com/papers/FastGaussianSmoothing.pdf
// !
// ! \param[out] boxes   boxes dimensions
// ! \param[in] sigma    Gaussian standard deviation
// ! \param[in] n        number of boxes
// !
func std_to_box(boxes []int, sigma float32, n int) {
	var wi float32 = alg.Sqrtf((12 * sigma * sigma / float32(n)) + 1)
	var wl int = int(alg.Floor(wi))
	if wl%2 == 0 {
		wl--
	}
	var wu int = wl + 2

	var mi float32 = (12*sigma*sigma - float32(n*wl*wl) - float32(4*n*wl) - float32(3*n)) / (-4*float32(wl) - 4)
	var m int = int(alg.Round(mi))

	for i := 0; i < n; i++ {
		if i < m {
			boxes[i] = (wu - 1) / 2
		} else {
			boxes[i] = (wl - 1) / 2
		}
	}
}

// !
// ! \fn void horizontal_blur_rgb(uchar * in, uchar * out, int w, int h, int c, int r)
// !
// ! \brief this function performs the horizontal blur pass for box blur.
// !
// ! \param[in,out] in       source channel
// ! \param[in,out] out      target channel
// ! \param[in] w            image width
// ! \param[in] h            image height
// ! \param[in] c            image channels
// ! \param[in] r            box dimension
// !
func horizontal_blur_rgb(in []uint8, out []uint8, w, h, c, r int) {
	var iarr float32 = 1.0 / (float32(r + r + 1))
	for i := 0; i < h; i++ {
		var ti int = i * w
		var li int = ti
		var ri int = ti + r

		var fv = [3]int{int(in[ti*c+0]), int(in[ti*c+1]), int(in[ti*c+2])}
		var lv = [3]int{int(in[(ti+w-1)*c+0]), int(in[(ti+w-1)*c+1]), int(in[(ti+w-1)*c+2])}
		var val = [3]int{(r + 1) * fv[0], (r + 1) * fv[1], (r + 1) * fv[2]}

		for j := 0; j < r; j++ {
			val[0] += int(in[(ti+j)*c+0])
			val[1] += int(in[(ti+j)*c+1])
			val[2] += int(in[(ti+j)*c+2])
		}

		for j := 0; j <= r; j, ri, ti = j+1, ri+1, ti+1 {
			val[0] += int(in[ri*c+0]) - fv[0]
			val[1] += int(in[ri*c+1]) - fv[1]
			val[2] += int(in[ri*c+2]) - fv[2]
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}

		for j := r + 1; j < w-r; j, ri, ti, li = j+1, ri+1, ti+1, li+1 {
			val[0] += int(in[ri*c+0]) - int(in[li*c+0])
			val[1] += int(in[ri*c+1]) - int(in[li*c+1])
			val[2] += int(in[ri*c+2]) - int(in[li*c+2])
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}

		for j := w - r; j < w; j, ti, li = j+1, ti+1, li+1 {
			val[0] += lv[0] - int(in[li*c+0])
			val[1] += lv[1] - int(in[li*c+1])
			val[2] += lv[2] - int(in[li*c+2])
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}
	}
}

// !
// ! \fn void total_blur_rgb(uchar * in, uchar * out, int w, int h, int c, int r)
// !
// ! \brief this function performs the total blur pass for box blur.
// !
// ! \param[in,out] in       source channel
// ! \param[in,out] out      target channel
// ! \param[in] w            image width
// ! \param[in] h            image height
// ! \param[in] c            image channels
// ! \param[in] r            box dimension
// !
func total_blur_rgb(in []uint8, out []uint8, w, h, c, r int) {
	// radius range on either side of a pixel + the pixel itself
	var iarr float32 = 1.0 / (float32(r + r + 1))
	for i := 0; i < w; i++ {
		var ti int = i
		var li int = ti
		var ri int = ti + r*w

		var fv = [3]int{int(in[ti*c+0]), int(in[ti*c+1]), int(in[ti*c+2])}
		var lv = [3]int{int(in[(ti+w*(h-1))*c+0]), int(in[(ti+w*(h-1))*c+1]), int(in[(ti+w*(h-1))*c+2])}
		var val = [3]int{(r + 1) * fv[0], (r + 1) * fv[1], (r + 1) * fv[2]}

		for j := 0; j < r; j++ {
			val[0] += int(in[(ti+j*w)*c+0])
			val[1] += int(in[(ti+j*w)*c+1])
			val[2] += int(in[(ti+j*w)*c+2])
		}

		for j := 0; j <= r; j, ri, ti = j+1, ri+w, ti+w {
			val[0] += int(in[ri*c+0]) - fv[0]
			val[1] += int(in[ri*c+1]) - fv[1]
			val[2] += int(in[ri*c+2]) - fv[2]
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}

		for j := r + 1; j < h-r; j, ri, ti, li = j+1, ri+w, ti+w, li+w {
			val[0] += int(in[ri*c+0]) - int(in[li*c+0])
			val[1] += int(in[ri*c+1]) - int(in[li*c+1])
			val[2] += int(in[ri*c+2]) - int(in[li*c+2])
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}

		for j := h - r; j < h; j, ti, li = j+1, ti+w, li+w {
			val[0] += lv[0] - int(in[li*c+0])
			val[1] += lv[1] - int(in[li*c+1])
			val[2] += lv[2] - int(in[li*c+2])
			out[ti*c+0] = uint8(alg.Round(float32(val[0]) * iarr))
			out[ti*c+1] = uint8(alg.Round(float32(val[1]) * iarr))
			out[ti*c+2] = uint8(alg.Round(float32(val[2]) * iarr))
		}
	}
}

// !
// ! \fn void box_blur_rgb(uchar * in, uchar * out, int w, int h, int c, int r)
// !
// ! \brief this function performs a box blur pass.
// !
// ! \param[in,out] in       source channel
// ! \param[in,out] out      target channel
// ! \param[in] w            image width
// ! \param[in] h            image height
// ! \param[in] c            image channels
// ! \param[in] r            box dimension
// !
func box_blur_rgb(in *[]uint8, out *[]uint8, w, h, c, r int) {
	in, out = out, in
	horizontal_blur_rgb((*out), (*in), w, h, c, r)
	total_blur_rgb((*in), (*out), w, h, c, r)
	// Note to myself :
	// here we could go anisotropic with different radiis rx,ry in HBlur and TBlur
}

// !
// ! \fn void fast_gaussian_blur_rgb(uchar * in, uchar * out, int w, int h, int c, float sigma)
// !
// ! \brief this function performs a fast Gaussian blur. Applying several
// ! times box blur tends towards a true Gaussian blur. Three passes are sufficient
// ! for good results. For further details please refer to :
// ! http://blog.ivank.net/fastest-gaussian-blur.html
// !
// ! \param[in,out] in       source channel
// ! \param[in,out] out      target channel
// ! \param[in] w            image width
// ! \param[in] h            image height
// ! \param[in] c            image channels
// ! \param[in] sigma        gaussian std dev
// !
func fast_gaussian_blur_rgb(in []uint8, out []uint8, w int, h int, c int, sigma float32) {
	// sigma conversion to box dimensions
	var boxes = []int{0, 0, 0}
	std_to_box(boxes, sigma, 3)
	box_blur_rgb(&in, &out, w, h, c, boxes[0])
	box_blur_rgb(&out, &in, w, h, c, boxes[1])
	box_blur_rgb(&in, &out, w, h, c, boxes[2])
}

func Fast_gaussian_blur_rgb(in []uint8, out []uint8, w int, h int, c int, sigma float32) {
	fast_gaussian_blur_rgb(in, out, w, h, c, sigma)
}
