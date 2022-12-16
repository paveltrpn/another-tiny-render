package algebra

import (
	"math"
	"math/bits"
	"unsafe"
)

/*	multidimensional array mapping, array[i][j]
	row-wise (C, C++):
	(0	1)
	(2	3)

	column-wise (Fortran, Matlab):
	(0	2)
	(1	3)
*/

func IdRw(i int32, j int32, n int32) int32 {
	return (i*n + j)
}

func IdCw(i int32, j int32, n int32) int32 {
	return (j*n + i)
}

func DegToRad(deg float32) float32 {
	return (deg * Pi) / 180.0
}

// ///////////////////////////////////////////////////////////
// Code below taken from https://github.com/chewxy/math32	//
// A float32 version of Go's math package.					//
// ///////////////////////////////////////////////////////////

/* math constants from libc math.h */

const (
	E   = float32(2.71828182845904523536028747135266249775724709369995957496696763) // http://oeis.org/A001113
	Pi  = float32(3.14159265358979323846264338327950288419716939937510582097494459) // http://oeis.org/A000796
	Phi = float32(1.61803398874989484820458683436563811772030917980576286213544862) // http://oeis.org/A001622

	Sqrt2   = float32(1.41421356237309504880168872420969807856967187537694807317667974) // http://oeis.org/A002193
	SqrtE   = float32(1.64872127070012814684865078781416357165377610071014801157507931) // http://oeis.org/A019774
	SqrtPi  = float32(1.77245385090551602729816748334114518279754945612238712821380779) // http://oeis.org/A002161
	SqrtPhi = float32(1.27201964951406896425242246173749149171560804184009624861664038) // http://oeis.org/A139339

	Ln2    = float32(0.693147180559945309417232121458176568075500134360255254120680009) // http://oeis.org/A002162
	Log2E  = float32(1 / Ln2)
	Ln10   = float32(2.30258509299404568401799145468436420760110148862877297603332790) // http://oeis.org/A002392
	Log10E = float32(1 / Ln10)
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32             = 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)
)

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

const (
	uvnan    = 0x7FE00000
	uvinf    = 0x7F800000
	uvone    = 0x3f800000
	uvneginf = 0xFF800000
	mask     = 0xFF
	shift    = 32 - 8 - 1
	bias     = 127
	signMask = 1 << 31
	fracMask = 1<<shift - 1
)

var mPi4 = [...]uint64{
	0x0000000000000001,
	0x45f306dc9c882a53,
	0xf84eafa3ea69bb81,
	0xb6c52b3278872083,
	0xfca2c757bd778ac3,
	0x6e48dc74849ba5c0,
	0x0c925dd413a32439,
	0xfc3bd63962534e7d,
	0xd1046bea5d768909,
	0xd338e04d68befc82,
	0x7323ac7306a673e9,
	0x3908bf177bf25076,
	0x3ff12fffbc0b301f,
	0xde5e2316b414da3e,
	0xda6cfd9e4f96136e,
	0x9e8c7ecd3cbfd45a,
	0xea4f758fd7cbe2f6,
	0x7a0e73ef14a525d4,
	0xd7f6bf623f1aba10,
	0xac06608df8f6d757,
}

// sin coefficients
var _sin = [...]float32{
	1.58962301576546568060e-10, // 0x3de5d8fd1fd19ccd
	-2.50507477628578072866e-8, // 0xbe5ae5e5a9291f5d
	2.75573136213857245213e-6,  // 0x3ec71de3567d48a1
	-1.98412698295895385996e-4, // 0xbf2a01a019bfdf03
	8.33333333332211858878e-3,  // 0x3f8111111110f7d0
	-1.66666666666666307295e-1, // 0xbfc5555555555548
}

// cos coefficients
var _cos = [...]float32{
	-1.13585365213876817300e-11, // 0xbda8fa49a0861a9b
	2.08757008419747316778e-9,   // 0x3e21ee9d7b4e3f05
	-2.75573141792967388112e-7,  // 0xbe927e4f7eac4bc6
	2.48015872888517045348e-5,   // 0x3efa01a019c844f5
	-1.38888888888730564116e-3,  // 0xbf56c16c16c14f91
	4.16666666666665929218e-2,   // 0x3fa555555555554b
}

func inf(sign int) float32 {
	var v uint32
	if sign >= 0 {
		v = uvinf
	} else {
		v = uvneginf
	}
	return float32frombits(v)
}

func naN() float32 { return float32frombits(uvnan) }

func isNaN(f float32) (is bool) {
	return f != f
}

func isInf(f float32, sign int) bool {
	return sign >= 0 && f > MaxFloat32 || sign <= 0 && f < -MaxFloat32
}

func normalize(x float32) (y float32, exp int) {
	const SmallestNormal = 1.1754943508222875079687365e-38 // 2**-(127 - 1)
	if Fabs(x) < SmallestNormal {
		return x * (1 << shift), -shift
	}
	return x, 0
}

func float32bits(f float32) uint32 { return *(*uint32)(unsafe.Pointer(&f)) }

func float32frombits(b uint32) float32 { return *(*float32)(unsafe.Pointer(&b)) }

func Fabs(x float32) float32 {
	return float32frombits(float32bits(x) &^ (1 << 31))
}

func copysign(x, y float32) float32 {
	const sign = 1 << 31
	return math.Float32frombits(math.Float32bits(x)&^sign | math.Float32bits(y)&sign)
}

const reduceThreshold = 1 << 29

func trigReduce(x float32) (j uint64, z float32) {
	const PI4 = Pi / 4
	if x < PI4 {
		return 0, x
	}
	// Extract out the integer and exponent such that,
	// x = ix * 2 ** exp.
	ix := float32bits(x)
	exp := int(ix>>shift&mask) - bias - shift
	ix &^= mask << shift
	ix |= 1 << shift
	// Use the exponent to extract the 3 appropriate uint64 digits from mPi4,
	// B ~ (z0, z1, z2), such that the product leading digit has the exponent -61.
	// Note, exp >= -53 since x >= PI4 and exp < 971 for maximum float64.
	const floatingbits = 32 - 3
	digit, bitshift := uint(exp+floatingbits)/32, uint(exp+floatingbits)%32
	z0 := (mPi4[digit] << bitshift) | (mPi4[digit+1] >> (32 - bitshift))
	z1 := (mPi4[digit+1] << bitshift) | (mPi4[digit+2] >> (32 - bitshift))
	z2 := (mPi4[digit+2] << bitshift) | (mPi4[digit+3] >> (32 - bitshift))
	// Multiply mantissa by the digits and extract the upper two digits (hi, lo).
	z2hi, _ := bits.Mul64(z2, uint64(ix))
	z1hi, z1lo := bits.Mul64(z1, uint64(ix))
	z0lo := z0 * uint64(ix)
	lo, c := bits.Add64(z1lo, z2hi, 0)
	hi, _ := bits.Add64(z0lo, z1hi, c)
	// The top 3 bits are j.
	j = hi >> floatingbits
	// Extract the fraction and find its magnitude.
	hi = hi<<3 | lo>>floatingbits
	lz := uint(bits.LeadingZeros64(hi))
	e := uint64(bias - (lz + 1))
	// Clear implicit mantissa bit and shift into place.
	hi = (hi << (lz + 1)) | (lo >> (32 - (lz + 1)))
	hi >>= 43 - shift
	// Include the exponent and convert to a float.
	hi |= e << shift
	z = float32frombits(uint32(hi))
	// Map zeros to origin.
	if j&1 == 1 {
		j++
		j &= 7
		z--
	}
	// Multiply the fractional part by pi/4.
	return j, z * PI4
}

func Ldexp(frac float32, exp int) float32 {
	// special cases
	switch {
	case frac == 0:
		return frac // correctly return -0
	case isInf(frac, 0) || isNaN(frac):
		return frac
	}
	frac, e := normalize(frac)
	exp += e
	x := float32bits(frac)
	exp += int(x>>shift)&mask - bias
	if exp < -149 {
		return copysign(0, frac) // underflow
	}
	if exp > 127 { // overflow
		if frac < 0 {
			return inf(-1)
		}
		return inf(1)
	}
	var m float32 = 1
	if exp < -(127 - 1) { // denormal
		exp += shift
		m = 1.0 / (1 << 23) // 1/(2**-23)
	}
	x &^= mask << shift
	x |= uint32(exp+bias) << shift
	return m * float32frombits(x)
}

func Expmulti(hi, lo float32, k int) float32 {
	const (
		P1 = float32(1.6666667163e-01)  /* 0x3e2aaaab */
		P2 = float32(-2.7777778450e-03) /* 0xbb360b61 */
		P3 = float32(6.6137559770e-05)  /* 0x388ab355 */
		P4 = float32(-1.6533901999e-06) /* 0xb5ddea0e */
		P5 = float32(4.1381369442e-08)  /* 0x3331bb4c */
	)

	r := hi - lo
	t := r * r
	c := r - t*(P1+t*(P2+t*(P3+t*(P4+t*P5))))
	y := 1 - ((lo - (r*c)/(2-c)) - hi)
	// TODO(rsc): make sure Ldexp can handle boundary k
	return Ldexp(y, k)
}

func Exp(x float32) float32 {
	const (
		Ln2Hi = float32(6.9313812256e-01)
		Ln2Lo = float32(9.0580006145e-06)
		Log2e = float32(1.4426950216e+00)

		Overflow  = 7.09782712893383973096e+02
		Underflow = -7.45133219101941108420e+02
		NearZero  = 1.0 / (1 << 28) // 2**-28

		LogMax = 0x42b2d4fc // The bitmask of log(FLT_MAX), rounded down.  This value is the largest input that can be passed to exp() without producing overflow.
		LogMin = 0x42aeac50 // The bitmask of |log(REAL_FLT_MIN)|, rounding down

	)
	// hx := Float32bits(x) & uint32(0x7fffffff)

	// special cases
	switch {
	case isNaN(x) || isInf(x, 1):
		return x
	case isInf(x, -1):
		return 0
	case x > Overflow:
		return inf(1)
	case x < Underflow:
		return 0
		// case hx > LogMax:
		// 	return Inf(1)
		// case x < 0 && hx > LogMin:
		// return 0
	case -NearZero < x && x < NearZero:
		return 1 + x
	}

	// reduce; computed as r = hi - lo for extra precision.
	var k int
	switch {
	case x < 0:
		k = int(Log2e*x - 0.5)
	case x > 0:
		k = int(Log2e*x + 0.5)
	}
	hi := x - float32(k)*Ln2Hi
	lo := float32(k) * Ln2Lo

	// compute
	return Expmulti(hi, lo, k)
}

func Tanf(x float32) float32 {
	return float32(math.Tan(float64(x)))
}

func Tanhf(x float32) float32 {
	const MAXLOG = 88.02969187150841
	z := Fabs(x)
	switch {
	case z > 0.5*MAXLOG:
		if x < 0 {
			return -1
		}
		return 1
	case z >= 0.625:
		s := Exp(z + z)
		z = 1 - 2/(s+1)
		if x < 0 {
			z = -z
		}
	default:
		if x == 0 {
			return x
		}
		s := x * x
		z = ((((-5.70498872745e-3*s+2.06390887954e-2)*s-5.37397155531e-2)*s+1.33314422036e-1)*s-3.33332819422e-1)*s*x + x
	}
	return z
}

func Sincosf(x float32) (sin, cos float32) {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0:
		return x, 1 // return ±0.0, 1.0
	case isNaN(x) || isInf(x, 0):
		return naN(), naN()
	}

	// make argument positive
	sinSign, cosSign := false, false
	if x < 0 {
		x = -x
		sinSign = true
	}

	var j uint64
	var y, z float32
	if x >= reduceThreshold {
		j, z = trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float32(j)           // integer part of x/(Pi/4), as float

		if j&1 == 1 { // map zeros to origin
			j++
			y++
		}
		j &= 7                               // octant modulo 2Pi radians (360 degrees)
		z = ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	}
	if j > 3 { // reflect in x axis
		j -= 4
		sinSign, cosSign = !sinSign, !cosSign
	}
	if j > 1 {
		cosSign = !cosSign
	}

	zz := z * z
	cos = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	sin = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	if j == 1 || j == 2 {
		sin, cos = cos, sin
	}
	if cosSign {
		cos = -cos
	}
	if sinSign {
		sin = -sin
	}
	return
}

func Sinf(x float32) float32 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case x == 0 || isNaN(x):
		return x // return ±0 || NaN()
	case isInf(x, 0):
		return naN()
	}

	// make argument positive but save the sign
	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	var j uint64
	var y, z float32
	if x >= reduceThreshold {
		j, z = trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float32(j)           // integer part of x/(Pi/4), as float

		// map zeros to origin
		if j&1 == 1 {
			j++
			y++
		}
		j &= 7                               // octant modulo 2Pi radians (360 degrees)
		z = ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	}
	// reflect in x axis
	if j > 3 {
		sign = !sign
		j -= 4
	}
	zz := z * z
	if j == 1 || j == 2 {
		y = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	} else {
		y = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	}
	if sign {
		y = -y
	}
	return y
}

func Cosf(x float32) float32 {
	const (
		PI4A = 7.85398125648498535156e-1  // 0x3fe921fb40000000, Pi/4 split into three parts
		PI4B = 3.77489470793079817668e-8  // 0x3e64442d00000000,
		PI4C = 2.69515142907905952645e-15 // 0x3ce8469898cc5170,
	)
	// special cases
	switch {
	case isNaN(x) || isInf(x, 0):
		return naN()
	}

	// make argument positive
	sign := false
	x = Fabs(x)

	var j uint64
	var y, z float32
	if x >= reduceThreshold {
		j, z = trigReduce(x)
	} else {
		j = uint64(x * (4 / Pi)) // integer part of x/(Pi/4), as integer for tests on the phase angle
		y = float32(j)           // integer part of x/(Pi/4), as float

		// map zeros to origin
		if j&1 == 1 {
			j++
			y++
		}
		j &= 7                               // octant modulo 2Pi radians (360 degrees)
		z = ((x - y*PI4A) - y*PI4B) - y*PI4C // Extended precision modular arithmetic
	}

	if j > 3 {
		j -= 4
		sign = !sign
	}
	if j > 1 {
		sign = !sign
	}

	zz := z * z
	if j == 1 || j == 2 {
		y = z + z*zz*((((((_sin[0]*zz)+_sin[1])*zz+_sin[2])*zz+_sin[3])*zz+_sin[4])*zz+_sin[5])
	} else {
		y = 1.0 - 0.5*zz + zz*zz*((((((_cos[0]*zz)+_cos[1])*zz+_cos[2])*zz+_cos[3])*zz+_cos[4])*zz+_cos[5])
	}
	if sign {
		y = -y
	}
	return y
}

func Sqrtf(x float32) float32 {
	// special cases
	switch {
	case x == 0 || isNaN(x) || isInf(x, 1):
		return x
	case x < 0:
		return naN()
	}
	ix := float32bits(x)

	// normalize x
	exp := int((ix >> shift) & mask)
	if exp == 0 { // subnormal x
		for ix&(1<<shift) == 0 {
			ix <<= 1
			exp--
		}
		exp++
	}
	exp -= bias // unbias exponent
	ix &^= mask << shift
	ix |= 1 << shift
	if exp&1 == 1 { // odd exp, double x to make it even
		ix <<= 1
	}
	exp >>= 1 // exp = exp/2, exponent of square root
	// generate sqrt(x) bit by bit
	ix <<= 1
	var q, s uint32               // q = sqrt(x)
	r := uint32(1 << (shift + 1)) // r = moving bit from MSB to LSB
	for r != 0 {
		t := s + r
		if t <= ix {
			s = t + r
			ix -= t
			q += r
		}
		ix <<= 1
		r >>= 1
	}
	// final rounding
	if ix != 0 { // remainder, result not exact
		q += q & 1 // round according to extra bit
	}
	ix = q>>1 + uint32(exp-1+bias)<<shift // significand + biased exponent
	return float32frombits(ix)

}

func Hypot(a float32, b float32, c float32) float32 {
	return Sqrtf(a*a + b*b + c*c)
}

func Round(x float32) float32 {
	// Round is a faster implementation of:
	//
	// func Round(x float64) float64 {
	//   t := Trunc(x)
	//   if Abs(x-t) >= 0.5 {
	//     return t + Copysign(1, x)
	//   }
	//   return t
	// }
	bits := float32bits(x)
	e := uint(bits>>shift) & mask
	if e < bias {
		// Round abs(x) < 1 including denormals.
		bits &= signMask // +-0
		if e == bias-1 {
			bits |= uvone // +-1
		}
	} else if e < bias+shift {
		// Round any abs(x) >= 1 containing a fractional component [0,1).
		//
		// Numbers with larger exponents are returned unchanged since they
		// must be either an integer, infinity, or NaN.
		const half = 1 << (shift - 1)
		e -= bias
		bits += half >> e
		bits &^= fracMask >> e
	}
	return float32frombits(bits)
}

func Modf(f float32) (int float32, frac float32) {
	if f < 1 {
		switch {
		case f < 0:
			int, frac = Modf(-f)
			return -int, -frac
		case f == 0:
			return f, f // Return -0, -0 when f == -0
		}
		return 0, f
	}

	x := float32bits(f)
	e := uint(x>>shift)&mask - bias

	// Keep the top 9+e bits, the integer part; clear the rest.
	if e < 32-9 {
		x &^= 1<<(32-9-e) - 1
	}
	int = float32frombits(x)
	frac = f - int
	return
}

func Floor(x float32) float32 {
	if x == 0 || isNaN(x) || isInf(x, 0) {
		return x
	}
	if x < 0 {
		d, fract := Modf(-x)
		if fract != 0.0 {
			d = d + 1
		}
		return -d
	}
	d, _ := Modf(x)
	return d
}
