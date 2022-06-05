package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	gmtry "another-tiny-render/internal/geometry"
	img "another-tiny-render/internal/image"
	alg "another-tiny-render/pkg/algebra_go"
)

var (
	GlfwWndPtr *glfw.Window
	WndWidth   int = 1280
	WndHeight  int = 720

	Aspect = float32(WndWidth) / float32(WndHeight)
)

func InitGlfwWindow() {
	if err := glfw.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init glfw!")
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.Samples, 4)

	wnd, err := glfw.CreateWindow(WndWidth, WndHeight, "legacy-gl", nil, nil)
	if err != nil {
		panic(err)
	}
	GlfwWndPtr = wnd

	GlfwWndPtr.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init OpenGl!")
		panic(err)
	}
}

func SetOglDefaults() {
	gl.Viewport(0, 0, int32(WndWidth), int32(WndHeight))
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)

	gl.Enable(gl.MULTISAMPLE)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func RegisterGlfwCallbacks() {
	keyCallback := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}

	GlfwWndPtr.SetKeyCallback(keyCallback)
}

func main() {
	runtime.LockOSThread()

	// Test kernel multiplication
	{
		foo := img.ExtractSurround(2)
		for _, i := range foo {
			println(i)
		}
	}

	InitGlfwWindow()
	SetOglDefaults()
	RegisterGlfwCallbacks()

	prspMtrx := alg.Mtrx4FromPerspective(alg.DegToRad(45.0), Aspect, 0.01, 100.0)
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
	texFile, _ := img.BuildCanvas(512, 512, 3)
	// texFile.LoadFromJpegFile("diffuse.jpg")
	texFile.DrawChecker(16)

	{
		var (
			r, g, b uint8
		)

		getRandColor := func() (uint8, uint8, uint8) {
			return uint8(rand.Intn(255)),
				uint8(rand.Intn(255)),
				uint8(rand.Intn(255))
		}

		texFile.MultPerComponent(0.8, 1.8, 0.1)

		r, g, b = getRandColor()
		texFile.SetPenColor(r, g, b)
		texFile.BrasenhamLine(10, 10, 500, 402)

		r, g, b = getRandColor()
		texFile.SetPenColor(r, g, b)
		texFile.BrasenhamLine(400, 20, 40, 350)

		for i := 1; i < 32; i++ {
			r, g, b = getRandColor()
			texFile.SetPenColor(r, g, b)
			texFile.BrasenhamCircle(256, 256, i*15)
		}

		r, g, b = getRandColor()
		texFile.SetPenColor(r, g, b)
		texFile.DDALine(440, 110, 40, 426)
	}

	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &tex)
	gl.BindTexture(gl.TEXTURE_2D, tex)
	gl.TexImage2D(gl.TEXTURE_2D, 0, int32(texFile.GetDepth()),
		int32(texFile.GetWidth()),
		int32(texFile.GetHeight()),
		0,
		gl.RGB,
		gl.UNSIGNED_BYTE,
		gl.Ptr(texFile.GetDataPtr()))

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

	for !GlfwWndPtr.ShouldClose() {
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

		GlfwWndPtr.SwapBuffers()
	}
}
