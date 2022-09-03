package service

import (
	"math"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

type IImGUIRender interface {
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

type SImGUIState struct {
	imGUIcontext  *imgui.Context
	imGUIio       imgui.IO
	imGUIrenderer IImGUIRender
}

var (
	glfwButtonIDByIndex = map[int]glfw.MouseButton{
		0: glfw.MouseButton1,
		1: glfw.MouseButton2,
		2: glfw.MouseButton3,
	}

	mouseJustPressed [3]bool
	platformTime     float64 = 0.0
)

func (state SAppState) ImguiNewFrame() {
	// Setup display size (every frame to accommodate for window resizing)
	displaySize := state.GetFramebufferSize()
	state.imGUIio.SetDisplaySize(imgui.Vec2{X: displaySize[0], Y: displaySize[1]})

	// Setup time step
	currentTime := glfw.GetTime()
	if platformTime > 0 {
		state.imGUIio.SetDeltaTime(float32(currentTime - platformTime))
	}
	platformTime = currentTime

	if state.GlfwWndPtr.GetAttrib(glfw.Focused) != 0 {
		x, y := state.GlfwWndPtr.GetCursorPos()
		state.imGUIio.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	} else {
		state.imGUIio.SetMousePosition(imgui.Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(mouseJustPressed); i++ {
		down := mouseJustPressed[i] || (state.GlfwWndPtr.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		state.imGUIio.SetMouseButtonDown(i, down)
		mouseJustPressed[i] = false
	}
}

func (state *SAppState) InitImGUI() {
	state.imGUIcontext = imgui.CreateContext(nil)
	state.imGUIio = imgui.CurrentIO()

	switch state.oglContextType {
	case oglContextLegacy:
		state.imGUIrenderer, _ = NewOpenGL2(state.imGUIio)
	case oglContextModern:
		state.imGUIrenderer, _ = NewOpenGL3(state.imGUIio)
	}
}

func (state *SAppState) DestroyImGUI() {
	state.imGUIcontext.Destroy()
}

func (state *SAppState) RenderImGUI() {
	imgui.Render()
	state.imGUIrenderer.Render(state.GetDisplaySize(), state.GetFramebufferSize(), imgui.RenderedDrawData())
}
