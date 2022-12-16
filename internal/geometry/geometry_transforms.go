package geometry

import (
	alg "another-tiny-render/pkg/algebra"
)

func (gmtry *SGmtryInstance) ScaleVerticesFloat(factor float32) {
	for i := 0; i < len(gmtry.Vertices); i = i + 3 {
		gmtry.Vertices[i+0] *= factor
		gmtry.Vertices[i+1] *= factor
		gmtry.Vertices[i+2] *= factor
	}
}

func (gmtry *SGmtryInstance) ScaleVerticesVec3(vec alg.Vec3) {
	for i := 0; i < len(gmtry.Vertices); i = i + 3 {
		gmtry.Vertices[i+0] *= vec[0]
		gmtry.Vertices[i+1] *= vec[1]
		gmtry.Vertices[i+2] *= vec[2]
	}
}

func (gmtry *SGmtryInstance) TransformVerticesVec3(vec alg.Vec3) {
	for i := 0; i < len(gmtry.Vertices); i = i + 3 {
		gmtry.Vertices[i+0] += vec[0]
		gmtry.Vertices[i+1] += vec[1]
		gmtry.Vertices[i+2] += vec[2]
	}
}

func (gmtry *SGmtryInstance) TransformVerticesMtrx4(A alg.Mtrx4) {
	for i := 0; i < len(gmtry.Vertices); i = i + 3 {
		x := gmtry.Vertices[i+0]
		y := gmtry.Vertices[i+1]
		z := gmtry.Vertices[i+2]

		w := A[3]*x + A[7]*y + A[11]*z + A[15]
		if w < 0.0000001 {
			w = 1.0
		}

		gmtry.Vertices[i+0] = (A[0]*x + A[4]*y + A[8]*z + A[12]) / w
		gmtry.Vertices[i+1] = (A[1]*x + A[5]*y + A[9]*z + A[13]) / w
		gmtry.Vertices[i+2] = (A[2]*x + A[6]*y + A[10]*z + A[14]) / w
	}
}

func (gmtry *SGmtryInstance) TransformNormalsMtrx4(A alg.Mtrx4) {
	for i := 0; i < len(gmtry.Vertices); i = i + 3 {
		x := gmtry.VertNormals[i+0]
		y := gmtry.VertNormals[i+1]
		z := gmtry.VertNormals[i+2]

		w := A[3]*x + A[7]*y + A[11]*z + A[15]
		if w < 0.0000001 {
			w = 1.0
		}

		gmtry.VertNormals[i+0] = (A[0]*x + A[4]*y + A[8]*z + A[12]) / w
		gmtry.VertNormals[i+1] = (A[1]*x + A[5]*y + A[9]*z + A[13]) / w
		gmtry.VertNormals[i+2] = (A[2]*x + A[6]*y + A[10]*z + A[14]) / w
	}
}
