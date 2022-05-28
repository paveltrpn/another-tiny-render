package geometry

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	GEOMETRIC_VERTICES = iota
	VERTEX_NORMALS
	TEXTURE_VERTICES
	PARAMETER_SPACE_VERTICES
	POINT_ELEMENT
	LINE_ELEMENT
	FACE_ELEMENT
	OBJECT_NAME
	SMOOTH_GROUP
	COMMENT
)

var objTokens = map[string]int{
	"#":  COMMENT,
	"v":  GEOMETRIC_VERTICES,
	"vn": VERTEX_NORMALS,
	"vt": TEXTURE_VERTICES,
	"vp": PARAMETER_SPACE_VERTICES,
	"p":  POINT_ELEMENT,
	"l":  LINE_ELEMENT,
	"f":  FACE_ELEMENT,
	"o":  OBJECT_NAME,
	"s":  SMOOTH_GROUP,
}

// Wavefron Obj file format with triangle faces records must look like this:
// "f v/vt/vn v/vt/vn v/vt/vn", i.e. must contains three groups
// of indexes and an "f" start token. If number of indexes groups greater than
// 3 then file contains a quad or poly faces.
func checkObjFacesIsTrangulated(line string) bool {
	tmp := strings.Split(line, " ")
	return len(tmp) <= 4
}

// Load only triangle faces geometry data (othervise throw an error)
// from file, just for only one object in file (othervise throw an error)
// and ignoring geometry groups, smooth groups and materials.
func (gmtry *SGmtryInstance) LoadFromWavefrontOBJ(fname string) error {
	var (
		objName  string
		floatBuf [4]float32
		intBuf   [9]uint32
	)

	if gmtry.IsExist {
		return fmt.Errorf("gmtry.LoadFromWavefrontOBJ(): object %v is allready allocated", gmtry.Name)
	}

	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	gmtry.Vertices = make([]float32, 0)
	gmtry.FaceIds = make([]uint32, 0)
	gmtry.VertNormals = make([]float32, 0)
	gmtry.FaceTexCoords = make([]float32, 0)

	const objNameUndefined = "__UNDEFINED__"
	gmtry.Name = objNameUndefined

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if line == "" {
			continue
		}
		token := strings.Split(line, " ")[0]

		switch token {
		case "#": // Skip comment lines
		case "o":
			if gmtry.Name != objNameUndefined {
				return fmt.Errorf("gmtry.LoadFromWavefrontOBJ(): Error!!! Attemt to load many object from single file - %s", fname)
			}
			fmt.Sscanf(line, "o %s", &objName)
			gmtry.Name = objName

		case "v":
			fmt.Sscanf(line, "v %f %f %f", &floatBuf[0], &floatBuf[1], &floatBuf[2])
			gmtry.Vertices = append(gmtry.Vertices, floatBuf[0], floatBuf[1], floatBuf[2])

		case "vn":
			fmt.Sscanf(line, "vn %f %f %f", &floatBuf[0], &floatBuf[1], &floatBuf[2])
			gmtry.VertNormals = append(gmtry.VertNormals, floatBuf[0], floatBuf[1], floatBuf[2])

		case "vt":
			fmt.Sscanf(line, "vt %f %f", &floatBuf[0], &floatBuf[1])
			gmtry.FaceTexCoords = append(gmtry.FaceTexCoords, floatBuf[0], floatBuf[1])

		case "f":
			if !checkObjFacesIsTrangulated(line) {
				return fmt.Errorf("gmtry.LoadFromWavefrontOBJ(): Error!!! Attemt to load non triangle face in file - %s", fname)
			}
			fmt.Sscanf(line, "f %d/%d/%d %d/%d/%d %d/%d/%d",
				&intBuf[0], &intBuf[1], &intBuf[2],
				&intBuf[3], &intBuf[4], &intBuf[5],
				&intBuf[6], &intBuf[7], &intBuf[8])
			// Decrease each element by one because of wavefront obj
			// format starts enumerate indexes with 1 instead of 0 (Pascal style).
			gmtry.FaceIds = append(gmtry.FaceIds, intBuf[0]-1, intBuf[1]-1, intBuf[2]-1,
				intBuf[3]-1, intBuf[4]-1, intBuf[5]-1,
				intBuf[6]-1, intBuf[7]-1, intBuf[8]-1)

		default:
			fmt.Printf("gmtry.LoadFromWavefrontOBJ(): Warn!!! Unknown token - %s in file - %s!\n", line, fname)
		}
	}

	gmtry.IsExist = true
	return nil
}
