package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

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
	gl.ClearColor(0.1, 0.1, 0.3, 1.0)

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

	InitGlfwWindow()
	SetOglDefaults()
	RegisterGlfwCallbacks()

	prspMtrx := alg.Mtrx4FromPerspective(alg.DegToRad(45.0), Aspect, 0.01, 100.0)
	mdlMtrx := alg.Mtrx4FromLookAt(alg.Vec3{5.0, 3.0, 5.0}, alg.Vec3{0.0, 0.0, 0.0}, alg.Vec3{0.0, 1.0, 0.0})

	for !GlfwWndPtr.ShouldClose() {
		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)

		gl.MatrixMode(gl.PROJECTION)
		gl.LoadMatrixf(&prspMtrx[0])

		gl.MatrixMode(gl.MODELVIEW)
		gl.LoadMatrixf(&mdlMtrx[0])

		// RENDER

		gl.Begin(gl.QUADS)
		gl.Vertex3f(-1.0, 1.0, 0.0)
		gl.Vertex3f(1.0, 1.0, 0.0)
		gl.Vertex3f(1.0, -1.0, 0.0)
		gl.Vertex3f(-1.0, -1.0, 0.0)
		gl.End()

		// gl.LoadIdentity()

		GlfwWndPtr.SwapBuffers()
	}
}
