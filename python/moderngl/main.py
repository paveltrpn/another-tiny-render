
import OpenGL
OpenGL.ERROR_CHECKING = False
from OpenGL.GL import * #type: ignore
from glfw.GLFW import * #type: ignore

import imgui as imgui
import imgui.integrations.glfw as imguiIntgr

import sys

# currentdir = os.path.dirname(os.path.realpath(__file__))
# parentdir = os.path.dirname(currentdir)
# sys.path.append(os.path.dirname(os.getcwd()))
from pathlib import Path
sys.path.append(str(Path(__file__).parents[1]))

import algebra_py as alg

import service
import shader
import camera

from math import radians

boxTris = alg.CVec3Arrayf([
    #  top
     1.0,  1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,
    -1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0, -1.0,
    #  bottom
     1.0, -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,
    -1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  front
    -1.0,  1.0,  1.0,  1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
     1.0,  1.0,  1.0,  1.0, -1.0,  1.0, -1.0, -1.0,  1.0,
    #  back
    -1.0,  1.0, -1.0,  1.0,  1.0, -1.0, -1.0, -1.0, -1.0,
     1.0,  1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0, -1.0,
    #  left
    -1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,
    -1.0, -1.0,  1.0, -1.0, -1.0, -1.0, -1.0,  1.0, -1.0,
    #  right
     1.0,  1.0, -1.0,  1.0,  1.0,  1.0,  1.0, -1.0,  1.0,
     1.0, -1.0,  1.0,  1.0, -1.0, -1.0,  1.0,  1.0, -1.0
])

boxNormals = alg.CVec3Arrayf([
     #  top
     0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,
     0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,
     #  bottom
     0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,
     0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,
     #  front
     0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,
     0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,
     #  back
     0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,
     0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,
     #  left
    -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0,
    -1.0,  0.0,  0.0, -1.0,  0.0,  0.0, -1.0,  0.0,  0.0,
     #  right
     1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0,
     1.0,  0.0,  0.0,  1.0,  0.0,  0.0,  1.0,  0.0,  0.0
])

boxColors = alg.CVec3Arrayf([
    # top (light green)
    0.5,  0.7,  0.5, 0.5,  0.7,  0.5, 0.5,  0.7,  0.5,
    0.5,  0.7,  0.5, 0.5,  0.7,  0.5, 0.5,  0.7,  0.5,
    #  bottom (dark green)
    0.5,  0.5,  0.9, 0.5,  0.5,  0.9, 0.5,  0.5,  0.9,
    0.5,  0.5,  0.9, 0.5,  0.5,  0.9, 0.5,  0.5,  0.9,
    #  front (light red)
    0.5,  0.5,  0.5, 0.5,  0.5,  0.5, 0.5,  0.5,  0.5,
    0.5,  0.5,  0.5, 0.5,  0.5,  0.5, 0.5,  0.5,  0.5,
    #  back  (dark red)
    0.5,  0.5,  0.5, 0.5,  0.5,  0.5, 0.5,  0.5,  0.5,
    0.5,  0.5,  0.5, 0.5,  0.5,  0.5, 0.5,  0.5,  0.5,
    #  left (light blue)
    0.5,  0.5,  0.1, 0.5,  0.5,  0.1, 0.5,  0.5,  0.1,
    0.5,  0.5,  0.1, 0.5,  0.5,  0.1, 0.5,  0.5,  0.1,
    #  right (dark blue)
    0.2,  0.5,  0.5, 0.2,  0.5,  0.5, 0.2,  0.5,  0.5,
    0.2,  0.5,  0.5, 0.2,  0.5,  0.5, 0.2,  0.5,  0.5,
])

def initImgui(appState: service.CAppState) -> None:
    global g_imguiImpl
    global g_imguiVersion
    global g_imguiRobotoFont

    imgui.core.create_context()

    g_imguiVersion = str(imgui.core.get_version())

    g_imguiImpl = imguiIntgr.GlfwRenderer(appState.glfwWndPtr)
    imgui.core.style_colors_dark()

    io = imgui.core.get_io()

    global g_imguiRobotoFont
    g_imguiRobotoFont = io.fonts.add_font_from_file_ttf("RobotoMono-Medium.ttf", 16)
    imgui.font(g_imguiRobotoFont)
    g_imguiImpl.refresh_font_texture()


def setOglDefaults(appState: service.CAppState) -> None:
    glViewport(0, 0, appState.wndWidth, appState.wndHeight,)
    glClearColor(0.1, 0.1, 0.1, 1.0)

    global g_rtn
    global g_rtnNorm
    g_rtn = alg.mtrx4FromEuler(radians(0.7), radians(0.5), radians(0.2))
    g_rtnNorm = alg.mtrx4GetTranspose(g_rtn)
    scl = alg.mtrx4FromScale(alg.vec3(0.4, 0.3, 0.3))

    boxTris.applyMtrx4(scl)

    global g_pplLight
    g_pplLight = shader.COglProgram("ppl_light")
    g_pplLight.appendShader([GL_VERTEX_SHADER, GL_FRAGMENT_SHADER],["vs_ppl.glsl", "fs_ppl.glsl"])
    g_pplLight.linkProgram()

    global g_flatLight
    g_flatLight = shader.COglProgram("flat_light")
    g_flatLight.appendShader([GL_VERTEX_SHADER, GL_FRAGMENT_SHADER],["vs.glsl", "fs.glsl"])
    g_flatLight.linkProgram()

    global boxBO
    global cubeVAO

    cubeVAO = glGenVertexArrays(1)
    boxBO = glGenBuffers(1)

    glBindBuffer(GL_ARRAY_BUFFER, boxBO)
    glBufferData(GL_ARRAY_BUFFER, boxTris.dataSize, boxTris.tobytes(), GL_DYNAMIC_DRAW)

    glBindVertexArray(cubeVAO)

    g_pplLight.passVertexAtribArray(3, "position")
    g_flatLight.passVertexAtribArray(3, "position")

    global normalBO
    normalBO = glGenBuffers(1)
    glBindBuffer(GL_ARRAY_BUFFER, normalBO)
    glBufferData(GL_ARRAY_BUFFER, boxNormals.dataSize, boxNormals.tobytes(), GL_DYNAMIC_DRAW)
    g_pplLight.passVertexAtribArray(3, "normal")
    g_flatLight.passVertexAtribArray(3, "normal")

    global colorBO
    colorBO = glGenBuffers(1)
    glBindBuffer(GL_ARRAY_BUFFER, colorBO)
    glBufferData(GL_ARRAY_BUFFER, boxColors.dataSize, boxColors.tobytes(), GL_DYNAMIC_DRAW)
    g_pplLight.passVertexAtribArray(3, "color")
    g_flatLight.passVertexAtribArray(3, "color")

    global g_camera
    g_camera = camera.CCamera()
    g_camera.eye = alg.vec3(0.0, 0.0, -2.0)
    g_camera.target = alg.vec3(0.0, 0.0, 0.0)
    g_camera.aspect = appState.wndWidth/appState.wndHeight

    # hdrFBO = glGenFramebuffers(1)
    # glBindFramebuffer(GL_FRAMEBUFFER, hdrFBO)
    # colorBuffer = glGenTextures(1)
    # glBindTexture(GL_TEXTURE_2D, colorBuffer)
    # glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA16F, g_appState.wndWidth, g_appState.wndHeight, 0, GL_RGBA, GL_FLOAT, None)
    # glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR)
    # glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR)
    # glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_S, GL_CLAMP_TO_EDGE)
    # glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_WRAP_T, GL_CLAMP_TO_EDGE)
    # glFramebufferTexture2D(GL_FRAMEBUFFER, GL_COLOR_ATTACHMENT0, GL_TEXTURE_2D, colorBuffer, 0)

    glEnable(GL_DEPTH_TEST)


def loop(appState: service.CAppState) -> None:
    appState.printAppStats()
    # appState.printExtensionsList()

    intense = 4.0
    value = 1.0
    lightType = True
    eyePos = -1.5

    while not glfwWindowShouldClose(appState.glfwWndPtr):
        glfwPollEvents()

        glClear(GL_COLOR_BUFFER_BIT)
        glClear(GL_DEPTH_BUFFER_BIT)

        g_camera.eye = alg.vec3(0.0, 0.0, eyePos)

        # if (lightType):
            # g_pplLight.use()
            # g_pplLight.passMtrx4("projMtrx", g_camera.projMtrx)
            # g_pplLight.passMtrx4("viewMtrx", g_camera.viewMtrx)
            # g_pplLight.passFloat("intense", intense)
            # g_pplLight.passFloat("value", value)
        # else:
            # g_flatLight.use()
            # g_flatLight.passMtrx4("projMtrx", g_camera.projMtrx)
            # g_flatLight.passMtrx4("viewMtrx", g_camera.viewMtrx)

        g_flatLight.use()
        g_flatLight.passMtrx4("projMtrx", g_camera.projMtrx)
        g_flatLight.passMtrx4("viewMtrx", g_camera.viewMtrx)

        boxTris.applyMtrx4(g_rtn)
        glBindBuffer(GL_ARRAY_BUFFER, boxBO)
        glBufferSubData(GL_ARRAY_BUFFER, 0, boxTris.dataSize, boxTris.tobytes())

        boxNormals.applyMtrx4(g_rtnNorm)
        glBindBuffer(GL_ARRAY_BUFFER, normalBO)
        glBufferSubData(GL_ARRAY_BUFFER, 0, boxNormals.dataSize, boxNormals.tobytes())

        # glBindVertexArray(cubeVAO)
        glDrawArrays(GL_TRIANGLES, 0, int(boxTris.capacity/3))

        # glDrawBuffers(1, GL_COLOR_ATTACHMENT0)

        #============GUI draw section================
        g_imguiImpl.process_inputs()
        imgui.core.new_frame()

        with imgui.font(g_imguiRobotoFont):
            imgui.core.begin("main")
            imgui.core.text("Imgui version - " + g_imguiVersion)
            imgui.core.text("Frame time - {f_time:3.2f} ms/frame {fps:3.2f}".format(f_time=(1000.0 / imgui.core.get_io().framerate),
                                                                                    fps=imgui.core.get_io().framerate))
            _, intense = imgui.core.slider_float("intense", intense,
                                                        min_value=0.0, max_value=15.0,
                                                        format="%.3f",
                                                        power=1.0)

            _, value = imgui.core.slider_float("value", value,
                                                        min_value=0.0, max_value=15.0,
                                                        format="%.3f",
                                                        power=1.0)

            _, lightType = imgui.core.checkbox("Flat or per pixel light", lightType)

            _, eyePos = imgui.core.slider_float("eyePos", eyePos,
                                                        min_value=-0.5, max_value=-2.5,
                                                        format="%.3f",
                                                        power=1.0)
            imgui.core.end()

        imgui.core.end_frame()
        imgui.core.render()
        g_imguiImpl.render(imgui.core.get_draw_data())
        #===========================================

        glfwSwapBuffers(appState.glfwWndPtr)


def shutdown() -> None:
    glfwTerminate()


def main() -> None:
    appState = service.CAppState()

    service.initGlfwWindow(appState)

    setOglDefaults(appState)

    initImgui(appState)

    service.registerGlfwCallbacks(appState)

    loop(appState)

    shutdown()


if __name__ == "__main__":
    main()
