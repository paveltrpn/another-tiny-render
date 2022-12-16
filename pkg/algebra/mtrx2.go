package algebra

import (
	"fmt"
	"math"
)

type Mtrx2 [4]float32

func Mtrx2IdttNaive() (rt Mtrx2) {
	const (
		mrange int32 = 2
	)

	var (
		i, j int32
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

func Mtrx2Idtt() (rt Mtrx2) {
	for elem := range rt {
		if (elem % 3) == 0 {
			rt[elem] = 1.0
		} else {
			rt[elem] = 0.0
		}
	}

	return rt
}

func Mtrx2FromArray(m [4]float32) (rt Mtrx2) {
	for i := range rt {
		rt[i] = m[i]
	}

	return rt
}

func Mtrx2FromFloat(a00, a01, a10, a11 float32) (rt Mtrx2) {
	rt[0] = a00
	rt[1] = a01
	rt[2] = a10
	rt[3] = a11

	return rt
}

func Mtrx2FromRtn(phi float32) (rt Mtrx2) {
	var (
		cosphi, sinphi float32
	)

	sinphi = Sinf(phi)
	cosphi = Cosf(phi)

	rt[0] = cosphi
	rt[1] = -sinphi
	rt[2] = -sinphi
	rt[3] = cosphi

	return rt
}

func Mtrx2Det(m Mtrx2) float32 {
	return m[0]*m[3] - m[1]*m[2]
}

func Mtrx2DetLU(m Mtrx2) (rt float32) {
	const (
		mrange int32 = 2
	)

	var (
		i            int32
		l, u         Mtrx2
		l_det, u_det float32
	)

	l, u = Mtrx2LU(m)

	l_det = l[0]
	u_det = u[0]

	for i = 1; i < mrange; i++ {
		l_det *= l[IdRw(i, i, mrange)]
		u_det *= u[IdRw(i, i, mrange)]
	}

	return l_det * u_det
}

func Mtrx2Mult(a, b Mtrx2) (rt Mtrx2) {
	const (
		mrange int32 = 2
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

func Mtrx2MultVec2(m Mtrx2, v Vec2) (rt Vec2) {
	const (
		mrange int32 = 2
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
	Нижнетреугольная (L, lm) матрица имеет единицы по диагонали
*/

func Mtrx2LU(m Mtrx2) (lm, um Mtrx2) {
	const (
		mrange int32 = 2
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

func Mtrx2LDLT(m Mtrx2) (lm Mtrx2, dv Vec2) {
	const (
		mrange int32 = 2
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
						fmt.Println("Mtrx2LDLT():A is not positive deﬁnite")
						return Mtrx2Idtt(), Vec2Set(0.0, 0.0)
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

func Mtrx2Transpose(m Mtrx2) (rt Mtrx2) {
	const (
		mrange int32 = 2
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

func Mtrx2Invert(m Mtrx2) Mtrx2 {
	det := Mtrx2Det(m)

	if Fabs(det) < SmallestNonzeroFloat32 {
		fmt.Println("Mtrx2Invert(): determinant is a zero!")
		return Mtrx2Idtt()
	}

	return Mtrx2FromFloat(m[3], -m[1]/det, -m[2]/det, m[0]/det)
}

func Mtrx2SolveGauss(m Mtrx2, v Vec2) (rt Vec2) {
	const (
		mrange int32 = 2
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

func Mtrx2InsertCmn(m Mtrx2, v Vec2, cmn int32) (rt Mtrx2) {
	const (
		mrange int32 = 2
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

func Mtrx2SolveKramer(m Mtrx2, v Vec2) (rt Vec2) {
	const (
		mrange int32 = 2
	)

	var (
		i       int32
		det     float32
		kr_mtrx Mtrx2
	)

	det = Mtrx2Det(m)

	if Fabs(det) < SmallestNonzeroFloat32 {
		fmt.Println("Mtrx2SolveKramer(): system has no solve")
		return Vec2Set(0.0, 0.0)
	}

	for i = 0; i < mrange; i++ {
		kr_mtrx = Mtrx2InsertCmn(m, v, i)
		rt[i] = Mtrx2Det(kr_mtrx) / det
	}

	return rt
}
