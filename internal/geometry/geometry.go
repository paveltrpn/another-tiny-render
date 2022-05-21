package geometry

var boxPoints = []float32{
	// Upper points
	1.0, 1.0, 1.0,
	1.0, 1.0, -1.0,
	-1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	// Lower points
	1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
}

var boxFaceIds = []uint32{
	// Up face
	0, 1, 2, 3, 1, 2,
	// Down face
	4, 5, 6, 7, 5, 6,
	// Right face
	0, 4, 5, 1, 4, 5,
	// Left face
	2, 3, 7, 6, 2, 7,
	// Front face
	0, 4, 7, 3, 4, 7,
	// Back face
	1, 5, 6, 2, 5, 6,
}

var boxVtNormals = []float32{
	0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
	0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
	1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
	0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
}

var boxColors = []float32{
	0.5, 0.7, 0.5, 0.5, 0.7, 0.5, 0.5, 0.7, 0.5,
	0.5, 0.7, 0.5, 0.5, 0.7, 0.5, 0.5, 0.7, 0.5,
	0.5, 0.5, 0.9, 0.5, 0.5, 0.9, 0.5, 0.5, 0.9,
	0.5, 0.5, 0.9, 0.5, 0.5, 0.9, 0.5, 0.5, 0.9,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.1, 0.5, 0.5, 0.1, 0.5, 0.5, 0.1,
	0.5, 0.5, 0.1, 0.5, 0.5, 0.1, 0.5, 0.5, 0.1,
	0.2, 0.5, 0.5, 0.2, 0.5, 0.5, 0.2, 0.5, 0.5,
	0.2, 0.5, 0.5, 0.2, 0.5, 0.5, 0.2, 0.5, 0.5,
}

type SGmtryInstance struct {
	// Points in space coordinates, three floats per vertex
	Vertices []float32
	// Only triangle faces
	FaceIds []uint32
	// Per vertex normals
	VtNormals []float32
	// Per face index texture coords
	FaceTexCoords []float32

	// Per vertex color
	VtColors []float32
}

func (gmtry *SGmtryInstance) LoadBox() {
	gmtry.Vertices = make([]float32, len(boxPoints))
	copy(gmtry.Vertices, boxPoints)

	gmtry.FaceIds = make([]uint32, len(boxFaceIds))
	copy(gmtry.FaceIds, boxFaceIds)

	gmtry.VtNormals = make([]float32, len(boxVtNormals))
	copy(gmtry.VtNormals, boxVtNormals)

	gmtry.VtColors = make([]float32, len(boxColors))
	copy(gmtry.VtColors, boxColors)
}
