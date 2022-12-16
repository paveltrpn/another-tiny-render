package algebra

type Vec3 [3]float32

func Vec3Copy(v Vec3) (rt Vec3) {
	rt[0] = v[0]
	rt[1] = v[1]
	rt[2] = v[2]

	return rt
}

func Vec3Set(x float32, y float32, z float32) (rt Vec3) {
	rt[0] = x
	rt[1] = y
	rt[2] = z

	return rt
}

func Vec3Lenght(v Vec3) float32 {
	return Sqrtf(v[0]*v[0] +
		v[1]*v[1] +
		v[2]*v[2])

}

func Vec3Normalize(v Vec3) (rt Vec3) {
	len := Vec3Lenght(v)

	if len > SmallestNonzeroFloat32 {
		rt[0] = v[0] / len
		rt[1] = v[1] / len
		rt[2] = v[2] / len
	}

	return rt
}

func Vec3Scale(v Vec3, scale float32) (rt Vec3) {
	rt[0] = v[0] * scale
	rt[1] = v[1] * scale
	rt[2] = v[2] * scale

	return rt
}

func Vec3Invert(v Vec3) (rt Vec3) {
	rt[0] = -v[0]
	rt[1] = -v[1]
	rt[2] = -v[2]

	return rt
}

func Vec3Dot(a Vec3, b Vec3) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Vec3Sum(a, b Vec3) (rt Vec3) {
	rt[0] = a[0] + b[0]
	rt[1] = a[1] + b[1]
	rt[2] = a[2] + b[2]

	return rt
}

func Vec3Sub(a, b Vec3) (rt Vec3) {
	rt[0] = a[0] - b[0]
	rt[1] = a[1] - b[1]
	rt[2] = a[2] - b[2]

	return rt
}

func Vec3Cross(a, b Vec3) (rt Vec3) {
	rt[0] = a[1]*b[2] - a[2]*b[1]
	rt[1] = a[2]*b[0] - a[0]*b[2]
	rt[2] = a[0]*b[1] - a[1]*b[0]

	return rt
}
