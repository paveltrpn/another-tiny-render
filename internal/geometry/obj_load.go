package geometry

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	VERTEX = iota
	NORMAL
	TEX_COORD
	FACE_ID
	OBJ_NAME
	SMOOTH_GROUP
)

var objTokens = map[string]int{
	"v":  VERTEX,
	"vn": NORMAL,
	"vt": TEX_COORD,
	"f":  FACE_ID,
	"o":  OBJ_NAME,
	"s":  SMOOTH_GROUP,
}

func (gmtry *SGmtryInstance) LoadFromWavefrontOBJ(fname string) {
	var (
		objName  string
		floatBuf [4]float32
		intBuf   [4]uint32
	)

	if gmtry.IsExist {
		println("gmtry.LoadFromWavefrontOBJ(): object %v is allready allocated!", gmtry.Name)
		return
	}

	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("LoadFromWavefrontOBJ(): Error!!!\n%v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	gmtry.Vertices = make([]float32, 0)
	gmtry.FaceIds = make([]uint32, 0)
	gmtry.VertNormals = make([]float32, 0)
	gmtry.FaceTexCoords = make([]float32, 0)

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, " ")[0]

		switch token {
		case "o":
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
			fmt.Sscanf(line, "f %d/%d/%d", &intBuf[0], &intBuf[1], intBuf[2])
			gmtry.FaceIds = append(gmtry.FaceIds, intBuf[0], intBuf[1], intBuf[2])

		default:
		}
	}

	gmtry.IsExist = true

	for _, foo := range gmtry.FaceIds {
		fmt.Printf("%v", foo)
	}
}
