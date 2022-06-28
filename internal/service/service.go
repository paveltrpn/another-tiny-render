package service

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	OGlContextLegacy = iota
	OGlContextModern
)

type SAppState struct {
	WndWidth   int
	WndHeight  int
	Aspect     float32
	AppName    string
	GlfwWndPtr *glfw.Window

	oglContextType int

	SImGUIState

	GlfwVersionStr string
	GlRenderStr    string
	GlVersionStr   string
	GlslVersionStr string
}

func (state *SAppState) Print() {
	fmt.Println(state.GlRenderStr)
	fmt.Println(state.GlVersionStr)
	fmt.Println(state.GlslVersionStr)
	fmt.Println(state.GlfwVersionStr)
}

func (state *SAppState) RegisterGlfwCallbacks() {
	keyCallback := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}

	state.GlfwWndPtr.SetKeyCallback(keyCallback)
}

func (state *SAppState) InitGlfwWindowModern() {
	if err := glfw.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init glfw!")
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	wnd, err := glfw.CreateWindow(state.WndWidth, state.WndHeight, state.AppName, nil, nil)
	if err != nil {
		panic(err)
	}
	state.GlfwWndPtr = wnd

	state.GlfwWndPtr.MakeContextCurrent()
	glfw.SwapInterval(1)

	major, minor, rev := glfw.GetVersion()
	state.GlfwVersionStr = fmt.Sprintf("%d.%d.%d", major, minor, rev)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init OpenGl!")
		panic(err)
	}

	state.GlRenderStr = gl.GoStr(gl.GetString(gl.RENDERER))
	state.GlVersionStr = gl.GoStr(gl.GetString(gl.VERSION))
	state.GlslVersionStr = gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	state.oglContextType = OGlContextModern
}

func (state *SAppState) InitGlfwWindowLegacy() {
	if err := glfw.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init glfw!")
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)
	glfw.WindowHint(glfw.Samples, 4)

	wnd, err := glfw.CreateWindow(state.WndWidth, state.WndHeight, state.AppName, nil, nil)
	if err != nil {
		panic(err)
	}
	state.GlfwWndPtr = wnd

	state.GlfwWndPtr.MakeContextCurrent()
	glfw.SwapInterval(1)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init OpenGl!")
		panic(err)
	}

	state.oglContextType = OGlContextLegacy
}

func (state *SAppState) SetOglDefaults() {
	gl.Viewport(0, 0, int32(state.WndWidth), int32(state.WndHeight))
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
}

func (state SAppState) GetDisplaySize() [2]float32 {
	w, h := state.GlfwWndPtr.GetSize()
	return [2]float32{float32(w), float32(h)}
}

func (state SAppState) GetFramebufferSize() [2]float32 {
	w, h := state.GlfwWndPtr.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}
