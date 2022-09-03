package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"

	gmtry "another-tiny-render/internal/geometry"
	img "another-tiny-render/internal/image"
	srv "another-tiny-render/internal/service"
	alg "another-tiny-render/pkg/algebra_go"
)

func main() {
	runtime.LockOSThread()

	appState := srv.SAppState{WndWidth: 1152,
		WndHeight: 768,
		AppName:   "legacygl",
		Aspect:    1152.0 / 768.0}

	appState.InitGlfwWindowLegacy()
	appState.SetOglDefaults()
	appState.RegisterGlfwCallbacks()
	appState.InitImGUI()
	defer appState.DestroyImGUI()

	prspMtrx := alg.Mtrx4FromPerspective(alg.DegToRad(45.0), appState.Aspect, 0.01, 100.0)
	mdlMtrx := alg.Mtrx4FromLookAt(alg.Vec3{5.0, 3.0, 45.0}, alg.Vec3{0.0, -7.0, 0.0}, alg.Vec3{0.0, 1.0, 0.0})
	rtn := alg.Mtrx4FromAxisAngl(alg.Vec3{0.0, 1.0, 0.0}, alg.DegToRad(0.5))

	var foo gmtry.SGmtryInstance
	err := foo.LoadFromWavefrontOBJ("../../assets/demon_baby/demon_baby.obj")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	foo.ScaleVerticesFloat(0.9)
	foo.TransformVerticesVec3(alg.Vec3{-5.0, 0, 0})

	gl.ColorMaterial(gl.FRONT_AND_BACK, gl.DIFFUSE)
	gl.Enable(gl.COLOR_MATERIAL)
	gl.PointSize(13.0)

	var tex uint32
	texFile, err := img.BuildCanvasFromFile("../../assets/demon_baby/diffuse.png")
	if err != nil {
		fmt.Println(err)
	}
	// texFile.DrawChecker(16)

	// {
	// var (
	// r, g, b uint8
	// )

	// getRandColor := func() (uint8, uint8, uint8) {
	// return uint8(rand.Intn(255)),
	// uint8(rand.Intn(255)),
	// uint8(rand.Intn(255))
	// }
	//
	// texFile.MultPerComponent(0.8, 1.8, 0.1)
	//
	// r, g, b = getRandColor()
	// texFile.SetPenColor(r, g, b)
	// texFile.BrasenhamLine(10, 10, 500, 402)
	//
	// r, g, b = getRandColor()
	// texFile.SetPenColor(r, g, b)
	// texFile.BrasenhamLine(400, 20, 40, 350)
	//
	// for i := 1; i < 32; i++ {
	// r, g, b = getRandColor()
	// texFile.SetPenColor(r, g, b)
	// texFile.BrasenhamCircle(256, 256, i*15)
	// }
	//
	// r, g, b = getRandColor()
	// texFile.SetPenColor(r, g, b)
	// texFile.DDALine(440, 110, 40, 426)
	// }

	gl.Enable(gl.MULTISAMPLE)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4,
		int32(texFile.GetWidth()),
		int32(texFile.GetHeight()),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		texFile.GetDataUnsafePtr())

	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	var (
		l1Pos     = [4]float32{18.0, 5.0, 0.0, 1.0}
		l1Diffuse = [4]float32{0.8, 0.8, 0.8, 1.0}
		l1Ambient = [4]float32{0.1, 0.1, 0.1, 1.0}
	)
	gl.Enable(gl.LIGHTING)
	gl.ShadeModel(gl.SMOOTH)
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &l1Ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &l1Diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &l1Pos[0])

	for !appState.GlfwWndPtr.ShouldClose() {
		glfw.PollEvents()

		foo.TransformVerticesMtrx4(rtn)
		foo.TransformNormalsMtrx4(rtn)

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)

		gl.MatrixMode(gl.PROJECTION)
		gl.LoadMatrixf(&prspMtrx[0])

		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadMatrixf(&mdlMtrx[0])

		// RENDER

		gl.Begin(gl.POINTS)
		gl.Vertex3fv(&l1Pos[0])
		gl.End()

		gl.Enable(gl.LIGHTING)
		gl.Enable(gl.LIGHT0)

		gl.Enable(gl.TEXTURE_2D)
		gl.BindTexture(gl.TEXTURE_2D, tex)

		for i := 0; i < len(foo.FaceIds); i = i + 9 {
			gl.Begin(gl.TRIANGLES)

			gl.Normal3fv(&foo.VertNormals[foo.FaceIds[i+2]*3])
			gl.TexCoord2fv(&foo.FaceTexCoords[foo.FaceIds[i+1]*2])
			gl.Vertex3fv(&foo.Vertices[foo.FaceIds[i+0]*3])

			gl.Normal3fv(&foo.VertNormals[foo.FaceIds[i+5]*3])
			gl.TexCoord2fv(&foo.FaceTexCoords[foo.FaceIds[i+4]*2])
			gl.Vertex3fv(&foo.Vertices[foo.FaceIds[i+3]*3])

			gl.Normal3fv(&foo.VertNormals[foo.FaceIds[i+8]*3])
			gl.TexCoord2fv(&foo.FaceTexCoords[foo.FaceIds[i+7]*2])
			gl.Vertex3fv(&foo.Vertices[foo.FaceIds[i+6]*3])
			gl.End()
		}

		gl.Disable(gl.LIGHTING)

		// -----------------------------------------------------------
		// Draw GUI
		// -----------------------------------------------------------

		appState.ImguiNewFrame()
		imgui.NewFrame()
		{
			imgui.Begin(appState.AppName)
			fpsString := "Frame time"
			imgui.Text(fpsString)

			imgui.End()
		}
		appState.RenderImGUI()

		// -----------------------------------------------------------
		// -----------------------------------------------------------
		// -----------------------------------------------------------

		appState.GlfwWndPtr.SwapBuffers()
	}
}
