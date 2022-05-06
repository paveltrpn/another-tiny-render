package main

import (
	"fmt"
	"math"
	"runtime"
	srv "tiny-render-go/internal/service"
	alg "tiny-render-go/pkg/algebra_go"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

var boxTris alg.Vec3Array = alg.Vec3Array{Data: []float32{
	1.0, 1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0, -1.0,
	1.0, -1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0, 1.0, -1.0, -1.0, 1.0,
	1.0, 1.0, 1.0, 1.0, -1.0, 1.0, -1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 1.0, -1.0, -1.0, -1.0, -1.0,
	1.0, 1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0,
	-1.0, -1.0, 1.0, -1.0, -1.0, -1.0, -1.0, 1.0, -1.0,
	1.0, 1.0, -1.0, 1.0, 1.0, 1.0, 1.0, -1.0, 1.0,
	1.0, -1.0, 1.0, 1.0, -1.0, -1.0, 1.0, 1.0, -1.0,
}}

var boxNormals alg.Vec3Array = alg.Vec3Array{Data: []float32{
	0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
	0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0,
	0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
	0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0,
	-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	-1.0, 0.0, 0.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	1.0, 0.0, 0.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
}}

var boxColors alg.Vec3Array = alg.Vec3Array{Data: []float32{
	0.5, 0.7, 0.5, 0.5, 0.7, 0.5, 0.5, 0.7, 0.5,
	0.5, 0.7, 0.5, 0.5, 0.7, 0.5, 0.5, 0.7, 0.5,
	0.5, 0.5, 0.9, 0.5, 0.5, 0.9, 0.5, 0.5, 0.9,
	0.5, 0.5, 0.9, 0.5, 0.5, 0.9, 0.5, 0.5, 0.9,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5,
	0.5, 0.5, 0.1, 0.5, 0.5, 0.1, 0.5, 0.5, 0.1,
	0.5, 0.5, 0.1, 0.5, 0.5, 0.1, 0.5, 0.5, 0.1,
	0.2, 0.5, 0.5, 0.2, 0.5, 0.5, 0.2, 0.5, 0.5,
	0.2, 0.5, 0.5, 0.2, 0.5, 0.5, 0.2, 0.5, 0.5,
}}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	0: glfw.MouseButton1,
	1: glfw.MouseButton2,
	2: glfw.MouseButton3,
}

var mouseJustPressed [3]bool
var platformTime float64 = 0.0

func imguiGetDisplaySize(appState srv.SAppState) [2]float32 {
	w, h := appState.GlfwWndPtr.GetSize()
	return [2]float32{float32(w), float32(h)}
}

func imguiGetFramebufferSize(appState srv.SAppState) [2]float32 {
	w, h := appState.GlfwWndPtr.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

func imguiNewFrame(appState srv.SAppState, imguiIO imgui.IO) {
	// Setup display size (every frame to accommodate for window resizing)
	displaySize := imguiGetFramebufferSize(appState)
	imguiIO.SetDisplaySize(imgui.Vec2{X: displaySize[0], Y: displaySize[1]})

	// Setup time step
	currentTime := glfw.GetTime()
	if platformTime > 0 {
		imguiIO.SetDeltaTime(float32(currentTime - platformTime))
	}
	platformTime = currentTime

	if appState.GlfwWndPtr.GetAttrib(glfw.Focused) != 0 {
		x, y := appState.GlfwWndPtr.GetCursorPos()
		imguiIO.SetMousePosition(imgui.Vec2{X: float32(x), Y: float32(y)})
	} else {
		imguiIO.SetMousePosition(imgui.Vec2{X: -math.MaxFloat32, Y: -math.MaxFloat32})
	}

	for i := 0; i < len(mouseJustPressed); i++ {
		down := mouseJustPressed[i] || (appState.GlfwWndPtr.GetMouseButton(glfwButtonIDByIndex[i]) == glfw.Press)
		imguiIO.SetMouseButtonDown(i, down)
		mouseJustPressed[i] = false
	}
}

func main() {
	var (
		cubeVAO  uint32
		boxBO    uint32
		normalBO uint32
		colorBO  uint32
	)
	appState := srv.SAppState{WndWidth: 1152, WndHeight: 768, AppName: "002_moderngl", Aspect: 1152.0 / 768.0}

	runtime.LockOSThread()

	srv.InitGlfwWindow(&appState)
	srv.RegisterGlfwCallbacks(&appState)
	srv.SetOglDefaults(&appState)

	appState.Print()

	flatLight := srv.SOglProgram{Name: "flat shade"}
	flatLight.AppendShader([]uint32{gl.VERTEX_SHADER, gl.FRAGMENT_SHADER}, []string{"vs.glsl", "fs.glsl"})
	flatLight.LinkProgram()

	flatLight.Use()
	prsp := alg.Mtrx4FromPerspective(alg.DegToRad(45.0), appState.Aspect, 0.01, 100.0)
	mdl := alg.Mtrx4FromLookAt(alg.Vec3{5.0, 3.0, 5.0}, alg.Vec3{0.0, 0.0, 0.0}, alg.Vec3{0.0, 1.0, 0.0})
	flatLight.PassMtrx4("projMtrx", prsp)
	flatLight.PassMtrx4("viewMtrx", mdl)

	gl.GenVertexArrays(1, &cubeVAO)
	gl.BindVertexArray(cubeVAO)

	gl.GenBuffers(1, &boxBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, boxBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(boxTris.Data)*3*4, gl.Ptr(boxTris.Data), gl.DYNAMIC_DRAW)

	flatLight.PassVertexAtribArray(3, "position")

	gl.GenBuffers(1, &normalBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, normalBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(boxNormals.Data)*3*4, gl.Ptr(boxNormals.Data), gl.DYNAMIC_DRAW)
	flatLight.PassVertexAtribArray(3, "normal")

	gl.GenBuffers(1, &colorBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(boxColors.Data)*3*4, gl.Ptr(boxColors.Data), gl.DYNAMIC_DRAW)
	flatLight.PassVertexAtribArray(3, "color")

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	rtn := alg.Mtrx4FromAxisAngl(alg.Vec3{1.0, 0.0, 1.0}, alg.DegToRad(0.7))
	// rtn := alg.Mtrx4FromRtnPitch(alg.DegToRad(4.0))
	// rtn := alg.Mtrx4FromEuler(0.0, 0.0, alg.DegToRad(4.0))
	rtnNrm := alg.Mtrx4GetTranspose(rtn)

	// === ImGUI === //
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	renderer, _ := srv.NewOpenGL3(io)
	// ============= //

	var eyePos [3]float32 = [3]float32{5.0, 3.0, 5.0}
	var targetPos [3]float32 = [3]float32{0.0, 0.0, 0.0}

	for !appState.GlfwWndPtr.ShouldClose() {
		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.Clear(gl.DEPTH_BUFFER_BIT)

		flatLight.Use()
		prsp = alg.Mtrx4FromPerspective(alg.DegToRad(45.0), appState.Aspect, 0.01, 100.0)
		mdl = alg.Mtrx4FromLookAt(alg.Vec3{eyePos[0], eyePos[1], eyePos[2]},
			alg.Vec3{targetPos[0], targetPos[1], targetPos[2]},
			alg.Vec3{0.0, 1.0, 0.0})

		flatLight.PassMtrx4("projMtrx", prsp)
		flatLight.PassMtrx4("viewMtrx", mdl)

		boxTris.ApplyMtrx4(&rtn)
		gl.BindBuffer(gl.ARRAY_BUFFER, boxBO)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(boxTris.Data)*3*4, gl.Ptr(boxTris.Data))

		boxNormals.ApplyMtrx4(&rtnNrm)
		gl.BindBuffer(gl.ARRAY_BUFFER, normalBO)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(boxNormals.Data)*3*4, gl.Ptr(boxNormals.Data))

		// gl.BindVertexArray(cubeVAO)

		gl.DrawArrays(gl.TRIANGLES, 0, 36)

		// -----------------------------------------------------------
		// Отрисовка меню
		// -----------------------------------------------------------

		imguiNewFrame(appState, io)
		imgui.NewFrame()
		{
			imgui.Begin("002_moderngl")
			fpsString := fmt.Sprintf("Frame time - %.3f ms/frame (%.1f FPS)", 1000.0/io.Framerate(), io.Framerate())
			imgui.Text(fpsString)

			imgui.SliderFloat("eyePosX", &eyePos[0], -5.0, 5)
			imgui.SliderFloat("eyePosY", &eyePos[1], -5.0, 5)
			imgui.SliderFloat("eyePosZ", &eyePos[2], -5.0, 5)

			imgui.SliderFloat("targetPosX", &targetPos[0], -5.0, 5.0)
			imgui.SliderFloat("targetPosY", &targetPos[1], -5.0, 5.0)
			imgui.SliderFloat("targetPosZ", &targetPos[2], -5.0, 5.0)

			imgui.End()
		}
		imgui.Render()
		renderer.Render(imguiGetDisplaySize(appState), imguiGetFramebufferSize(appState), imgui.RenderedDrawData())

		// -----------------------------------------------------------

		appState.GlfwWndPtr.SwapBuffers()
	}
}
