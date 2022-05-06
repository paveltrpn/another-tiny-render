package algebra_go

type Vec2 [2]float32

func Vec2Copy(v Vec2) (rt Vec2) {
	rt[0] = v[0]
	rt[1] = v[1]

	return rt
}

func Vec2Set(x float32, y float32) (rt Vec2) {
	rt[0] = x
	rt[1] = y

	return rt
}

func Vec2Lenght(v Vec2) float32 {
	return Sqrtf(v[0]*v[0] +
		v[1]*v[1])

}

func Vec2Normalize(v Vec2) (rt Vec2) {
	len := Vec2Lenght(v)

	if len != 0.0 {
		rt[0] = v[0] / len
		rt[1] = v[1] / len
	}

	return rt
}

func Vec2Scale(v Vec2, scale float32) (rt Vec2) {
	v[0] *= scale
	v[1] *= scale

	return rt
}

func Vec2Invert(v Vec2) (rt Vec2) {
	rt[0] = -v[0]
	rt[1] = -v[1]

	return rt
}

func Vec2Dot(a Vec2, b Vec2) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

func Vec2Sum(a, b Vec2) (rt Vec2) {
	rt[0] = a[0] + b[0]
	rt[1] = a[1] + b[1]

	return rt
}

func Vec2Sub(a, b Vec2) (rt Vec2) {
	rt[0] = a[0] - b[0]
	rt[1] = a[1] - b[1]

	return rt
}
