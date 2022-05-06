package algebra_go

type Vec4 [4]float32

func Vec4Copy(v Vec4) (rt Vec4) {
	rt[0] = v[0]
	rt[1] = v[1]
	rt[2] = v[2]
	rt[3] = v[3]

	return rt
}

func Vec4Set(x, y, z, w float32) (rt Vec4) {
	rt[0] = x
	rt[1] = y
	rt[2] = z
	rt[3] = w

	return rt
}

func Vec4Lenght(v Vec4) float32 {
	return Sqrtf(v[0]*v[0] +
		v[1]*v[1] +
		v[2]*v[2] +
		v[3]*v[3])

}

func Vec4Normalize(v Vec4) (rt Vec4) {
	len := Vec4Lenght(v)

	if len != 0.0 {
		rt[0] = v[0] / len
		rt[1] = v[1] / len
		rt[2] = v[2] / len
		rt[3] = v[3] / len
	}

	return rt
}

func Vec4Scale(v Vec4, scale float32) (rt Vec4) {
	v[0] *= scale
	v[1] *= scale
	v[2] *= scale
	v[3] *= scale

	return rt
}

func Vec4Invert(v Vec4) (rt Vec4) {
	rt[0] = -v[0]
	rt[1] = -v[1]
	rt[2] = -v[2]
	rt[3] = -v[3]

	return rt
}

func Vec4Dot(a Vec4, b Vec4) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func Vec4Sum(a, b Vec4) (rt Vec4) {
	rt[0] = a[0] + b[0]
	rt[1] = a[1] + b[1]
	rt[2] = a[2] + b[2]
	rt[3] = a[3] + b[3]

	return rt
}

func Vec4Sub(a, b Vec4) (rt Vec4) {
	rt[0] = a[0] - b[0]
	rt[1] = a[1] - b[1]
	rt[2] = a[2] - b[2]
	rt[3] = a[3] - b[3]

	return rt
}
