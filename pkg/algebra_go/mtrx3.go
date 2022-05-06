package algebra_go

import (
	"fmt"
	"math"
)

type Mtrx3 [9]float32

func Mtrx3IdttNaive() (rt Mtrx3) {
	var (
		i, j   int32
		mrange int32 = 3
	)

	for i = 0; i < mrange; i++ {
		for j = 0; j < mrange; j++ {
			if i == j {
				rt[IdRw(i, j, mrange)] = 1.0
			} else {
				rt[IdRw(i, j, mrange)] = 0.0
			}
		}
	}

	return rt
}

func Mtrx3Idtt() (rt Mtrx3) {
	for elem := range rt {
		if (elem % 4) == 0 {
			rt[elem] = 1.0
		} else {
			rt[elem] = 0.0
		}
	}

	return rt
}

func Mtrx3FromArray(m [9]float32) (rt Mtrx3) {
	for i := range rt {
		rt[i] = m[i]
	}

	return rt
}

func Mtrx3FromFloat(a00, a01, a02,
	a10, a11, a12,
	a20, a21, a22 float32) (rt Mtrx3) {

	rt[0] = a00
	rt[1] = a01
	rt[2] = a02

	rt[3] = a10
	rt[4] = a11
	rt[5] = a12

	rt[6] = a20
	rt[7] = a21
	rt[8] = a22

	return rt
}

func Mtrx3FromEuler(yaw, pitch, roll float32) (rt Mtrx3) {
	siny, cosy := Sincosf(yaw)
	sinp, cosp := Sincosf(pitch)
	sinr, cosr := Sincosf(roll)

	rt[0] = cosy*cosr - siny*cosp*sinr
	rt[1] = -cosy*sinr - siny*cosp*cosr
	rt[2] = siny * sinp

	rt[3] = siny*cosr + cosy*cosp*sinr
	rt[4] = -siny*sinr + cosy*cosp*cosr
	rt[5] = -cosy * sinp

	rt[6] = sinp * sinr
	rt[7] = sinp * cosr
	rt[8] = cosp

	return rt
}

func Mtrx3FromAxisangl(ax Vec3, phi float32) (rt Mtrx3) {
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

	rt[3] = t*vxvy - sinphi*vz
	rt[4] = cosphi + t*vy*vy
	rt[5] = t*vyvz + sinphi*vx

	rt[6] = t*vxvz + sinphi*vy
	rt[7] = t*vyvz - sinphi*vx
	rt[8] = cosphi + t*vz*vz

	return rt
}

func Mtrx3det(m Mtrx3) float32 {
	return m[0]*m[4]*m[8] +
		m[6]*m[1]*m[5] +
		m[2]*m[3]*m[7] -
		m[0]*m[7]*m[5] -
		m[8]*m[3]*m[1]
}

func Mtrx3DetLU(m Mtrx3) (rt float32) {
	const (
		mrange int32 = 3
	)

	var (
		i            int32
		l, u         Mtrx3
		l_det, u_det float32
	)

	l, u = Mtrx3LU(m)

	l_det = l[0]
	u_det = u[0]

	for i = 1; i < mrange; i++ {
		l_det *= l[IdRw(i, i, mrange)]
		u_det *= u[IdRw(i, i, mrange)]
	}

	return l_det * u_det
}

func Mtrx3Mult(a, b Mtrx3) (rt Mtrx3) {
	const (
		mrange int32 = 3
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

func Mtrx3MultVec3(m Mtrx3, v Vec3) (rt Vec3) {
	const (
		mrange int32 = 3
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

/*
	Где-то здесь ошибка, долго искал
	ничего не вышло и взял код из сети
*/
/*
func Mtrx3lu(m Mtrx3) (l, u Mtrx3) {
	var (
		i, j, k int32
		lm, um  Mtrx3
		sum     float32
	)

	for j = 0; j < 3; j++ {
		um[IdRw(0, j, 3)] = m[IdRw(0, j, 3)]
	}

	for j = 0; j < 3; j++ {
		lm[IdRw(j, 0, 3)] = m[IdRw(j, 0, 3)] / um[IdRw(0, 0, 3)]
	}

	for i = 1; i < 3; i++ {
		for j = i; j < 3; j++ {
			sum = 0.0
			for k = 0; k < i; k++ {
				sum = sum + lm[IdRw(i, k, 3)]*um[IdRw(k, j, 3)]
			}
			um[IdRw(i, j, 3)] = m[IdRw(i, j, 3)] - sum
		}
	}

	for i = 1; i < 3; i++ {
		for j = i; j < 3; j++ {
			if i > j {
				lm[IdRw(j, i, 3)] = 0.0
			} else {
				sum = 0.0
				for k = 0; k < i; k++ {
					sum = sum + lm[IdRw(j, k, 3)]*um[IdRw(k, i, 3)]
				}
				lm[IdRw(j, i, 3)] = (1.0 / um[IdRw(i, i, 3)]) * (m[IdRw(j, i, 3)] - sum)
			}
		}
	}

	return lm, um
}
*/

/*
	Нижнетреугольная (L, lm) матрица имеет единицы по диагонали
*/
func Mtrx3LU(m Mtrx3) (lm, um Mtrx3) {
	const (
		mrange int32 = 3
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

func Mtrx3LDLT(m Mtrx3) (lm Mtrx3, dv Vec3) {
	const (
		mrange int32 = 3
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
						fmt.Println("Mtrx3LDLT(): mtrx is not positive deﬁnite")
						return Mtrx3Idtt(), Vec3Set(0.0, 0.0, 0.0)
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

func Mtrx3Tanspose(m Mtrx3) (rt Mtrx3) {
	const (
		mrange int32 = 3
	)

	var (
		i, j int32
		tmp  float32
	)

	rt = m

	for i = 0; i < mrange; i++ {
		for j = 0; j < i; j++ {
			tmp = rt[IdRw(i, i, mrange)]
			rt[IdRw(i, j, mrange)] = rt[IdRw(j, i, mrange)]
			rt[IdRw(j, i, mrange)] = tmp
		}
	}

	return rt
}

func Mtrx3Invert(m Mtrx3) (rt Mtrx3) {
	var (
		inverse     Mtrx3
		det, invDet float32
	)

	inverse[0] = m[4]*m[8] - m[5]*m[7]
	inverse[3] = m[5]*m[6] - m[3]*m[8]
	inverse[6] = m[3]*m[7] - m[4]*m[6]

	det = m[0]*inverse[0] + m[1]*inverse[3] +
		m[2]*inverse[6]

	if Fabs(det) < SmallestNonzeroFloat32 {
		fmt.Println("Mtrx3Invert(): determinant is a zero!")
		return Mtrx3Idtt()
	}

	invDet = 1.0 / det

	inverse[1] = m[2]*m[7] - m[1]*m[8]
	inverse[2] = m[1]*m[5] - m[2]*m[4]
	inverse[4] = m[0]*m[8] - m[2]*m[6]
	inverse[5] = m[2]*m[3] - m[0]*m[5]
	inverse[7] = m[1]*m[6] - m[0]*m[7]
	inverse[8] = m[0]*m[4] - m[1]*m[3]

	rt[0] = inverse[0] * invDet
	rt[1] = inverse[1] * invDet
	rt[2] = inverse[2] * invDet

	rt[3] = inverse[3] * invDet
	rt[4] = inverse[4] * invDet
	rt[5] = inverse[5] * invDet

	rt[6] = inverse[6] * invDet
	rt[7] = inverse[7] * invDet
	rt[8] = inverse[8] * invDet

	return rt
}

func Mtrx3SolveGauss(m Mtrx3, v Vec3) (rt Vec3) {
	const (
		mrange int32 = 3
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

func Mtrx3InsertCmn(m Mtrx3, v Vec3, cmn int32) (rt Mtrx3) {
	const (
		mrange int32 = 3
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

func Mtrx3SolveKramer(m Mtrx3, v Vec3) (rt Vec3) {
	const (
		mrange int32 = 3
	)

	var (
		i       int32
		det     float32
		kr_mtrx Mtrx3
	)

	det = Mtrx3det(m)

	if Fabs(det) < SmallestNonzeroFloat32 {
		fmt.Println("Mtrx3SolveKramer(): system has no solve")
		return Vec3Set(0.0, 0.0, 0.0)
	}

	for i = 0; i < mrange; i++ {
		kr_mtrx = Mtrx3InsertCmn(m, v, i)
		rt[i] = Mtrx3det(kr_mtrx) / det
	}

	return rt
}
