package service

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	oglContextLegacy = iota
	oglContextModern
)

type SAppState struct {
	WndWidth   int
	WndHeight  int
	Aspect     float32
	AppName    string
	GlfwWndPtr *glfw.Window

	oglContextType int

	// contains global global ImGUI related fields
	// related to current OpenGL context
	SImGUIState

	// contains info strings about GLFW and OpenGL contexts
	// [0] - GLFW version string
	// [1] - OpenGL render string
	// [2] - OpenGL version string
	// [3] - OpenGL GLSL version string
	contextInfo [4]string
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
	state.contextInfo[0] = fmt.Sprintf("%d.%d.%d", major, minor, rev)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindowModern(): Error! Can't init OpenGl!")
		panic(err)
	}

	state.contextInfo[1] = gl.GoStr(gl.GetString(gl.RENDERER))
	state.contextInfo[2] = gl.GoStr(gl.GetString(gl.VERSION))
	state.contextInfo[3] = gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	state.oglContextType = oglContextModern
}

func (state *SAppState) InitGlfwWindowLegacy() {
	if err := glfw.Init(); err != nil {
		fmt.Println("InitGlfwWindowLegacy(): Error! Can't init glfw!")
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

	major, minor, rev := glfw.GetVersion()
	state.contextInfo[0] = fmt.Sprintf("%d.%d.%d", major, minor, rev)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindowLegacy(): Error! Can't init OpenGl!")
		panic(err)
	}

	state.contextInfo[1] = gl.GoStr(gl.GetString(gl.RENDERER))
	state.contextInfo[2] = gl.GoStr(gl.GetString(gl.VERSION))
	state.contextInfo[3] = gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	state.oglContextType = oglContextLegacy
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
