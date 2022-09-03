package main

import (
	srv "another-tiny-render/internal/service"
	alg "another-tiny-render/pkg/algebra_go"
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/inkyblackness/imgui-go/v4"
)

/*
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
*/

func main() {
	var (
	//cubeVAO  uint32
	//boxBO    uint32
	//normalBO uint32
	//colorBO  uint32
	)
	appState := srv.SAppState{WndWidth: 1152, WndHeight: 768, AppName: "002_moderngl", Aspect: 1152.0 / 768.0}

	runtime.LockOSThread()

	appState.InitGlfwWindowModern()
	appState.RegisterGlfwCallbacks()
	appState.SetOglDefaults()

	appState.InitImGUI()
	defer appState.DestroyImGUI()

	flatLight := srv.SOglProgram{Name: "flat shade"}
	flatLight.AppendShader([]uint32{gl.VERTEX_SHADER, gl.FRAGMENT_SHADER}, []string{"assets/vs.glsl", "assets/fs.glsl"})
	flatLight.LinkProgram()

	flatLight.Use()
	prsp := alg.Mtrx4FromPerspective(alg.DegToRad(45.0), appState.Aspect, 0.01, 100.0)
	mdl := alg.Mtrx4FromLookAt(alg.Vec3{5.0, 3.0, 5.0}, alg.Vec3{0.0, 0.0, 0.0}, alg.Vec3{0.0, 1.0, 0.0})
	flatLight.PassMtrx4("projMtrx", prsp)
	flatLight.PassMtrx4("viewMtrx", mdl)

	/*
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
	*/
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	//rtn := alg.Mtrx4FromAxisAngl(alg.Vec3{1.0, 0.0, 1.0}, alg.DegToRad(0.7))
	// rtn := alg.Mtrx4FromRtnPitch(alg.DegToRad(4.0))
	// rtn := alg.Mtrx4FromEuler(0.0, 0.0, alg.DegToRad(4.0))
	//rtnNrm := alg.Mtrx4GetTranspose(rtn)

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

		//flatLight.PassMtrx4("projMtrx", prsp)
		//flatLight.PassMtrx4("viewMtrx", mdl)

		//boxTris.ApplyMtrx4(&rtn)
		//gl.BindBuffer(gl.ARRAY_BUFFER, boxBO)
		//gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(boxTris.Data)*3*4, gl.Ptr(boxTris.Data))

		//boxNormals.ApplyMtrx4(&rtnNrm)
		//gl.BindBuffer(gl.ARRAY_BUFFER, normalBO)
		//gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(boxNormals.Data)*3*4, gl.Ptr(boxNormals.Data))

		// gl.BindVertexArray(cubeVAO)

		//gl.DrawArrays(gl.TRIANGLES, 0, 36)

		// -----------------------------------------------------------
		// Отрисовка меню
		// -----------------------------------------------------------

		appState.ImguiNewFrame()
		imgui.NewFrame()
		{
			imgui.Begin(appState.AppName)
			fpsString := fmt.Sprintf("Frame time")
			imgui.Text(fpsString)

			imgui.End()
		}
		appState.RenderImGUI()

		// appState.ImguiNewFrame()
		// imgui.NewFrame()
		// {
		// 	imgui.Begin("002_moderngl")
		// 	fpsString := fmt.Sprintf("Frame time - %.3f ms/frame (%.1f FPS)")
		// 	imgui.Text(fpsString)

		// 	imgui.SliderFloat("eyePosX", &eyePos[0], -5.0, 5)
		// 	imgui.SliderFloat("eyePosY", &eyePos[1], -5.0, 5)
		// 	imgui.SliderFloat("eyePosZ", &eyePos[2], -5.0, 5)

		// 	imgui.SliderFloat("targetPosX", &targetPos[0], -5.0, 5.0)
		// 	imgui.SliderFloat("targetPosY", &targetPos[1], -5.0, 5.0)
		// 	imgui.SliderFloat("targetPosZ", &targetPos[2], -5.0, 5.0)

		// 	imgui.End()
		// }
		// imgui.Render()
		// appState.RenderImGUI()

		// -----------------------------------------------------------

		appState.GlfwWndPtr.SwapBuffers()
	}
}
