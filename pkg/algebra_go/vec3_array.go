package algebra_go

type Vec3Array struct {
	Data []float32
}

func (array *Vec3Array) ApplyVec3(vec *Vec3) {
	for i := 0; i < len(array.Data); i += 3 {
		array.Data[i+0] = array.Data[i+0] + vec[0]
		array.Data[i+1] = array.Data[i+1] + vec[1]
		array.Data[i+2] = array.Data[i+2] + vec[2]
	}
}

func (array *Vec3Array) ApplyMtrx3(mat *Mtrx3) {
	for i := 0; i < len(array.Data); i += 3 {
		x := array.Data[i+0]
		y := array.Data[i+1]
		z := array.Data[i+2]

		array.Data[i+0] = x*mat[0] + y*mat[3] + z*mat[6]
		array.Data[i+1] = x*mat[1] + y*mat[4] + z*mat[7]
		array.Data[i+2] = x*mat[2] + y*mat[5] + z*mat[8]
	}
}

func (array *Vec3Array) ApplyMtrx4(mat *Mtrx4) {
	for i := 0; i < len(array.Data); i += 3 {
		x := array.Data[i+0]
		y := array.Data[i+1]
		z := array.Data[i+2]

		w := mat[3]*x + mat[7]*y + mat[11]*z + mat[15]
		if w < SmallestNonzeroFloat32 {
			w = 1.0
		}

		array.Data[i+0] = (mat[0]*x + mat[4]*y + mat[8]*z + mat[12]) / w
		array.Data[i+1] = (mat[1]*x + mat[5]*y + mat[9]*z + mat[13]) / w
		array.Data[i+2] = (mat[2]*x + mat[6]*y + mat[10]*z + mat[14]) / w
	}
}
