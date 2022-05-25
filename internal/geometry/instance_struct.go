package geometry

const (
	BUILTIN = iota // Load basic shape from built in data
	OBJ            // Wavefront obj
	STL            // Stereolitography
	GLTF           // Khronos GLTF
)

type SGmtryInstance struct {
	Name    string
	IsExist bool
	// Coordinates of points in space, three floats per vertex
	Vertices []float32
	// Only triangle faces
	FaceIds []uint32
	// Three normal per face
	VertNormals []float32
	// Three texture coords per face
	FaceTexCoords []float32

	// Three vertex color per face
	VtColors []float32

	// From which file type data is obtained
	FileType int
}
