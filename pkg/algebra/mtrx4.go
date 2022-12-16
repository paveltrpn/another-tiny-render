package algebra

import (
	"fmt"
	"math"
)

type Mtrx4 [16]float32

func Mtrx4IdttNaive() (rt Mtrx4) {
	var (
		i, j int32
		n    int32 = 4
	)

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i == j {
				rt[IdRw(i, j, n)] = 1.0
			} else {
				rt[IdRw(i, j, n)] = 0.0
			}
		}
	}

	return rt
}

func Mtrx4Idtt() (rt Mtrx4) {
	for elem := range rt {
		if (elem % 5) == 0 {
			rt[elem] = 1.0
		} else {
			rt[elem] = 0.0
		}
	}

	return rt
}

func Mtrx4FromArray(m [16]float32) (rt Mtrx4) {
	for i := range rt {
		rt[i] = m[i]
	}

	return rt
}

func Mtrx4FromFloat(a00, a01, a02, a03,
	a10, a11, a12, a13,
	a20, a21, a22, a23,
	a30, a31, a32, a33 float32) (rt Mtrx4) {

	rt[0] = a00
	rt[1] = a01
	rt[2] = a02
	rt[3] = a03

	rt[4] = a10
	rt[5] = a11
	rt[6] = a12
	rt[7] = a13

	rt[8] = a20
	rt[9] = a21
	rt[10] = a22
	rt[11] = a23

	rt[12] = a30
	rt[13] = a31
	rt[14] = a32
	rt[15] = a33

	return rt
}

func Mtrx4FromEuler(yaw, pitch, roll float32) (rt Mtrx4) {
	siny, cosy := Sincosf(yaw)
	sinp, cosp := Sincosf(pitch)
	sinr, cosr := Sincosf(roll)

	rt[0] = cosy*cosr - siny*cosp*sinr
	rt[1] = -cosy*sinr - siny*cosp*cosr
	rt[2] = siny * sinp
	rt[3] = 0.0

	rt[4] = siny*cosr + cosy*cosp*sinr
	rt[5] = -siny*sinr + cosy*cosp*cosr
	rt[6] = -cosy * sinp
	rt[7] = 0.0

	rt[8] = sinp * sinr
	rt[9] = sinp * cosr
	rt[10] = cosp
	rt[11] = 0.0

	rt[12] = 0.0
	rt[13] = 0.0
	rt[14] = 0.0
	rt[15] = 1.0

	return rt
}

func Mtrx4FromRtnYaw(angl float32) (rt Mtrx4) {
	rt = Mtrx4Idtt()
	var (
		sa, ca float32
	)

	sa, ca = Sincosf(angl)

	rt[5] = ca
	rt[6] = -sa
	rt[9] = sa
	rt[10] = ca

	return rt
}

func Mtrx4FromRtnPitch(angl float32) (rt Mtrx4) {
	rt = Mtrx4Idtt()
	var (
		sa, ca float32
	)

	sa, ca = Sincosf(angl)

	rt[0] = ca
	rt[2] = sa
	rt[8] = -sa
	rt[10] = ca

	return rt
}

func Mtrx4FromRtnRoll(angl float32) (rt Mtrx4) {
	rt = Mtrx4Idtt()
	var (
		sa, ca float32
	)

	sa, ca = Sincosf(angl)

	rt[0] = ca
	rt[1] = -sa
	rt[4] = sa
	rt[5] = ca

	return rt
}

func Mtrx4FromAxisAngl(ax Vec3, phi float32) (rt Mtrx4) {
	var (
		vxvy, vxvz, vyvz, vx, vy, vz float32
	)

	sinphi, cosphi := Sincosf(phi)

	vx = ax[0]
	vy = ax[1]
	vz = ax[2]

	// Нормализуем
	len := Hypot(vx, vy, vz)

	if len < SmallestNonzeroFloat32 {
		return rt
	}

	len = 1 / len
	vx *= len
	vy *= len
	vz *= len
	// ===========

	vxvy = vx * vy
	vxvz = vx * vz
	vyvz = vy * vz

	t := 1.0 - cosphi

	rt[0] = cosphi + t*vx*vx
	rt[1] = t*vxvy + sinphi*vz
	rt[2] = t*vxvz - sinphi*vy
	rt[3] = 0.0

	rt[4] = t*vxvy - sinphi*vz
	rt[5] = cosphi + t*vy*vy
	rt[6] = t*vyvz + sinphi*vx
	rt[7] = 0.0

	rt[8] = t*vxvz + sinphi*vy
	rt[9] = t*vyvz - sinphi*vx
	rt[10] = cosphi + t*vz*vz
	rt[11] = 0.0

	rt[12] = 0.0
	rt[13] = 0.0
	rt[14] = 0.0
	rt[15] = 1.0

	return rt
}

func Mtrx4FromOffset(src Vec3) (rt Mtrx4) {
	rt = Mtrx4Idtt()

	rt[3] = src[0]
	rt[7] = src[1]
	rt[11] = src[2]

	return rt
}

func Mtrx4FromScale(src Vec3) (rt Mtrx4) {
	rt = Mtrx4Idtt()

	rt[0] = src[0]
	rt[5] = src[1]
	rt[10] = src[2]

	return rt
}

func Mtrx4FromPerspective(fovy float32, aspect float32, near float32, far float32) (rt Mtrx4) {
	rt = Mtrx4Idtt()
	f := 1.0 / Tanf(fovy/2.0)
	var nf float32

	rt[0] = f / aspect
	rt[1] = 0.0
	rt[2] = 0.0
	rt[3] = 0.0
	rt[4] = 0.0
	rt[5] = f
	rt[6] = 0.0
	rt[7] = 0.0
	rt[8] = 0.0
	rt[9] = 0.0
	rt[11] = -1.0
	rt[12] = 0.0
	rt[13] = 0.0
	rt[15] = 0.0

	if far >= SmallestNonzeroFloat32 {
		nf = 1.0 / (near - far)
		rt[10] = (far + near) * nf
		rt[14] = 2.0 * far * near * nf
	} else {
		rt[10] = -1.0
		rt[14] = -2.0 * near
	}

	return rt
}

func Mtrx4FromLookAt(eye Vec3, center Vec3, up Vec3) (out Mtrx4) {
	out = Mtrx4Idtt()

	if Fabs(eye[0]-center[0]) < SmallestNonzeroFloat32 &&
		Fabs(eye[1]-center[1]) < SmallestNonzeroFloat32 &&
		Fabs(eye[2]-center[2]) < SmallestNonzeroFloat32 {
		return out
	}

	z0 := eye[0] - center[0]
	z1 := eye[1] - center[1]
	z2 := eye[2] - center[2]

	len := 1.0 / Hypot(z0, z1, z2) //??? было просто hypot
	z0 *= len
	z1 *= len
	z2 *= len

	x0 := up[1]*z2 - up[2]*z1
	x1 := up[2]*z0 - up[0]*z2
	x2 := up[0]*z1 - up[1]*z0

	len = Hypot(x0, x1, x2)
	if len < SmallestNonzeroFloat32 {
		x0 = 0
		x1 = 0
		x2 = 0
	} else {
		len = 1.0 / len
		x0 *= len
		x1 *= len
		x2 *= len
	}

	y0 := z1*x2 - z2*x1
	y1 := z2*x0 - z0*x2
	y2 := z0*x1 - z1*x0

	len = Hypot(y0, y1, y2)
	if len < SmallestNonzeroFloat32 {
		y0 = 0
		y1 = 0
		y2 = 0
	} else {
		len = 1.0 / len
		y0 *= len
		y1 *= len
		y2 *= len
	}

	out[0] = x0
	out[1] = y0
	out[2] = z0
	out[3] = 0.0
	out[4] = x1
	out[5] = y1
	out[6] = z1
	out[7] = 0.0
	out[8] = x2
	out[9] = y2
	out[10] = z2
	out[11] = 0.0
	out[12] = -(x0*eye[0] + x1*eye[1] + x2*eye[2])
	out[13] = -(y0*eye[0] + y1*eye[1] + y2*eye[2])
	out[14] = -(z0*eye[0] + z1*eye[1] + z2*eye[2])
	out[15] = 1.0

	return out
}

func Mtrx4GetTranspose(m Mtrx4) (rt Mtrx4) {
	var (
		i, j int32
		tmp  float32
	)

	rt = m

	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			tmp = rt[IdRw(i, j, 4)]
			rt[IdRw(i, j, 4)] = rt[IdRw(j, i, 4)]
			rt[IdRw(j, i, 4)] = tmp
		}
	}

	return rt
}

func Mtrx4det(m Mtrx4) float32 {
	return 0.0
}

func Mtrx4det_lu(m Mtrx4) (rt float32) {
	const (
		mrange int32 = 4
	)

	var (
		i            int32
		l, u         Mtrx4
		l_det, u_det float32
	)

	l, u = Mtrx4LU(m)

	l_det = l[0]
	u_det = u[0]

	for i = 1; i < mrange; i++ {
		l_det *= l[IdRw(i, i, mrange)]
		u_det *= u[IdRw(i, i, mrange)]
	}

	return l_det * u_det
}

func Mtrx4Mult(a, b Mtrx4) (rt Mtrx4) {
	const (
		mrange int32 = 4
	)

	var (
		i, j, k int32
		tmp     float32
	)

	for i = 0; i < mrange; i++ {
		for j = 0; j < mrange; j++ {
			tmp = 0.0
			for k = 0; k < mrange; k++ {
				tmp = tmp + a[IdRw(k, j, mrange)]*b[IdRw(i, k, mrange)]
			}
			rt[IdRw(i, j, mrange)] = tmp
		}
	}

	return rt
}

func Mtrx4MultVec(m Mtrx4, v Vec4) (rt Vec4) {
	const (
		mrange int32 = 4
	)

	var (
		i, j int32
		tmp  float32
	)

	for i = 0; i < mrange; i++ {
		tmp = 0
		for j = 0; j < mrange; j++ {
			tmp = tmp + m[IdRw(i, j, mrange)]*v[j]
		}
		rt[i] = tmp
	}

	return rt
}

func Mtrx4LU(m Mtrx4) (lm, um Mtrx4) {
	const (
		mrange int32 = 4
	)

	var (
		i, j, k int32
		sum     float32
	)

	for i = 0; i < mrange; i++ {
		for k = i; k < mrange; k++ {
			sum = 0
			for j = 0; j < i; j++ {
				sum += (lm[IdRw(i, j, mrange)] * um[IdRw(j, k, mrange)])
			}
			um[IdRw(i, k, mrange)] = m[IdRw(i, k, mrange)] - sum
		}

		for k = i; k < mrange; k++ {
			if i == k {
				lm[IdRw(i, i, mrange)] = 1.0
			} else {
				sum = 0
				for j = 0; j < i; j++ {
					sum += lm[IdRw(k, j, mrange)] * um[IdRw(j, i, mrange)]
				}
				lm[IdRw(k, i, mrange)] = (m[IdRw(k, i, mrange)] - sum) / um[IdRw(i, i, mrange)]
			}
		}
	}

	return lm, um
}

func Mtrx4LDLT(m Mtrx4) (lm Mtrx4, dv Vec4) {
	const (
		mrange int32 = 4
	)

	var (
		i, j, k int32
		sum     float32
	)

	for i = 0; i < mrange; i++ {
		for j = i; j < mrange; j++ {
			sum = m[IdRw(j, i, mrange)]
			for k = 0; k < i; k++ {
				sum = sum - lm[IdRw(i, k, mrange)]*dv[k]*lm[IdRw(j, k, mrange)]
				if i == j {
					if sum <= 0 {
						fmt.Println("Mtrx4LDLT(): mtrx is not positive definite")
						return Mtrx4Idtt(), Vec4Set(0.0, 0.0, 0.0, 0.0)
					}
					dv[i] = sum
					lm[IdRw(i, i, mrange)] = 1.0
				} else {
					lm[IdRw(j, i, mrange)] = sum / dv[i]
				}
			}
		}
	}

	return lm, dv
}

func Mtrx4Invert(m Mtrx4) (rt Mtrx4) {
	return Mtrx4Idtt()
}

func Mtrx4SolveGauss(m Mtrx4, v Vec4) (rt Vec4) {
	const (
		mrange int32 = 4
	)

	var (
		i, j, k int32
		t       float32
		a       [mrange][mrange + 1]float32
	)

	for i = 0; i < mrange; i++ { //было ++i
		for j = 0; j < mrange; j++ { //было ++j
			a[i][j] = m[IdRw(i, j, mrange)]
			a[i][mrange] = v[i]
		}
	}

	/* Pivotisation */
	for i = 0; i < mrange; i++ {
		for k = i + 1; k < mrange; k++ {
			if math.Abs(float64(a[i][i])) < math.Abs(float64(a[k][i])) {
				for j = 0; j <= mrange; j++ {
					t = a[i][j]
					a[i][j] = a[k][j]
					a[k][j] = t
				}
			}
		}
	}

	/* прямой ход */
	for k = 1; k < mrange; k++ {
		for j = k; j < mrange; j++ {
			t = a[j][k-1] / a[k-1][k-1]
			for i = 0; i < mrange+1; i++ {
				a[j][i] = a[j][i] - t*a[k-1][i]
			}
		}
	}

	/* обратный ход */
	for i = mrange - 1; i >= 0; i-- {
		rt[i] = a[i][mrange] / a[i][i]
		for j = mrange - 1; j > i; j-- {
			rt[i] = rt[i] - a[i][j]*rt[j]/a[i][i]
		}
	}

	return rt
}

func Mtrx4InsertCmn(m Mtrx4, v Vec4, cmn int32) (rt Mtrx4) {
	const (
		mrange int32 = 4
	)

	var (
		i int32
		j int32 = 0
	)

	rt = m

	for i = cmn; i < mrange*mrange; i += mrange {
		rt[i] = v[j]
		j++
	}
	return rt

}

func Mtrx4SolveKramer(m Mtrx4, v Vec4) (rt Vec4) {
	const (
		mrange int32 = 4
	)

	var (
		i       int32
		det     float32
		kr_mtrx Mtrx4
	)

	det = Mtrx4det_lu(m)

	if Fabs(det) < SmallestNonzeroFloat32 {
		fmt.Println("Mtrx4SolveKramer(): system has no solve")
		return Vec4Set(0.0, 0.0, 0.0, 0.0)
	}

	for i = 0; i < mrange; i++ {
		kr_mtrx = Mtrx4InsertCmn(m, v, i)
		rt[i] = Mtrx4det_lu(kr_mtrx) / det
	}

	return rt
}
