package service

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type SAppState struct {
	WndWidth   int
	WndHeight  int
	Aspect     float32
	AppName    string
	GlfwWndPtr *glfw.Window

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

func RegisterGlfwCallbacks(appState *SAppState) {
	keyCallback := func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape {
			w.SetShouldClose(true)
		}
	}

	appState.GlfwWndPtr.SetKeyCallback(keyCallback)
}

func InitGlfwWindow(appState *SAppState) {
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

	wnd, err := glfw.CreateWindow(appState.WndWidth, appState.WndHeight, appState.AppName, nil, nil)
	if err != nil {
		panic(err)
	}
	appState.GlfwWndPtr = wnd

	appState.GlfwWndPtr.MakeContextCurrent()
	glfw.SwapInterval(1)

	major, minor, rev := glfw.GetVersion()
	appState.GlfwVersionStr = fmt.Sprintf("%d.%d.%d", major, minor, rev)

	if err := gl.Init(); err != nil {
		fmt.Println("InitGlfwWindow(): Error! Can't init OpenGl!")
		panic(err)
	}

	appState.GlRenderStr = gl.GoStr(gl.GetString(gl.RENDERER))
	appState.GlVersionStr = gl.GoStr(gl.GetString(gl.VERSION))
	appState.GlslVersionStr = gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
}

func SetOglDefaults(appState *SAppState) {
	gl.Viewport(0, 0, int32(appState.WndWidth), int32(appState.WndHeight))
	gl.ClearColor(0.1, 0.1, 0.1, 1.0)
}
