
import glfw

# import OpenGL
# OpenGL.ERROR_CHECKING = False

from OpenGL.GL import *
import OpenGL.GL.shaders

import imgui as imgui
import imgui.integrations.glfw as imguiIntgr
import glm as glm

import mesh
import ogl_utils

def initGlfwWindow():
    global g_window
    global g_glfwVersion
    global g_wndWidth
    global g_wndHeight
    global g_appName

    g_wndWidth = 1152
    g_wndHeight = 864
    g_appName = "019"

    if not glfw.init():
        return
    
    # Если в хинтах выставить версию то скрипт падает на ошибке OpenGL (если проверка включена, см. вверху)
    # на вызовых glDrawArrays() glDrawElements() и т.д. c GL_INVALID_OPERATION
    # 
    # glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 4)
    # glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 6)
    # glfw.window_hint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)

    glfw.window_hint(glfw.DOUBLEBUFFER, glfw.TRUE)
    glfw.window_hint(glfw.RESIZABLE,    glfw.FALSE)
    
    g_window = glfw.create_window(g_wndWidth, g_wndHeight, g_appName, None, None)
    if not g_window:
        glfw.terminate()
        return
    
    glfw.make_context_current(g_window)

    [ver_min, ver_max, ver] = glfw.get_version()
    g_glfwVersion = str(ver_min) + "." + str(ver_max) + "." + str(ver)

    #  ВКЛ-ВЫКЛ вертикальную синхронизацию (VSYNC)
	#  Лок на 60 фпс
    glfw.swap_interval(glfw.TRUE)

def registerGlfwCallbacks():
    def keyCallback(window, key, scancode, action, mods):
        if (key == glfw.KEY_ESCAPE): 
            glfw.set_window_should_close(window, glfw.TRUE)

        if (key == glfw.KEY_Q): 
            print("Q is pressed!")

    glfw.set_key_callback(g_window, keyCallback)

def setOglDefaults():
    global g_glRenderStr
    global g_glVersionStr
    global g_glslVersionStr

    g_glRenderStr = str(glGetString(GL_RENDERER), "utf-8")
    g_glVersionStr = str(glGetString(GL_VERSION), "utf-8")
    g_glslVersionStr = str(glGetString(GL_SHADING_LANGUAGE_VERSION), "utf-8")

    glViewport(0, 0, g_wndWidth, g_wndHeight)
    glClearColor(0.0, 0.0, 0.0, 0.0)

    g_progMngr = ogl_utils.COglProgramMngr()
    g_progMngr.appendShader("VS", GL_VERTEX_SHADER, ogl_utils.g_VS)
    g_progMngr.appendShader("FS", GL_FRAGMENT_SHADER, ogl_utils.g_FS)
    g_progMngr.appendProgram("main_prog", ["VS", "FS"])

    triangle = glm.array(glm.vec3(-0.5, -0.5, 0.0), glm.vec3(1.0, 0.0, 0.0),
                         glm.vec3( 0.5, -0.5, 0.0), glm.vec3(0.0, 1.0, 0.0),
                         glm.vec3( 0.0,  0.5, 0.0), glm.vec3(0.0, 0.0, 1.0))

    # triangle = np.array(triangle, dtype = np.float32)

    VBO = glGenBuffers(1)
    glBindBuffer(GL_ARRAY_BUFFER, VBO)
    glBufferData(GL_ARRAY_BUFFER, 72, triangle.ptr, GL_STATIC_DRAW)

    position = glGetAttribLocation(g_progMngr.programList[0].progObj, "position")
    glVertexAttribPointer(position, 3, GL_FLOAT, GL_FALSE, 24, ctypes.c_void_p(0))
    glEnableVertexAttribArray(position)

    color = glGetAttribLocation(g_progMngr.programList[0].progObj, "color")
    glVertexAttribPointer(color, 3, GL_FLOAT, GL_FALSE, 24, ctypes.c_void_p(12))
    glEnableVertexAttribArray(color)

    glUseProgram(g_progMngr.programList[0].progObj)

def initImgui():
    global g_imguiImpl
    global g_imguiVersion
    global g_imguiRobotoFont

    imgui.create_context()

    g_imguiVersion = str(imgui.get_version())

    g_imguiImpl = imguiIntgr.GlfwRenderer(g_window)
    imgui.style_colors_dark()
    
    io = imgui.get_io()
    
    g_imguiRobotoFont = io.fonts.add_font_from_file_ttf("assets/RobotoMono-Medium.ttf", 16)
    imgui.font(g_imguiRobotoFont)
    g_imguiImpl.refresh_font_texture()

def loop():
    while not glfw.window_should_close(g_window):
        glfw.poll_events()

        glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
        glDrawArrays(GL_TRIANGLES, 0, 3)

        #============GUI draw section================
        g_imguiImpl.process_inputs()
        imgui.new_frame()

        # with imgui.font(roboto):

        imgui.push_font(g_imguiRobotoFont)
        imgui.begin("019_python_OpenGL")
        imgui.text("Frame time - {f_time:3.2f} ms/frame {fps:3.2f}".format(f_time=(1000.0 / imgui.get_io().framerate), fps=imgui.get_io().framerate)); 
        imgui.text("OpenGL GL_RENDERER - " + g_glRenderStr)
        imgui.text("OpenGL GL_VERSION - " + g_glVersionStr)
        imgui.text("OpenGL GLSL_VERSION - " + g_glslVersionStr)
        imgui.text("Glfw version - " + g_glfwVersion)
        imgui.text("Imgui version - " + g_imguiVersion)
        imgui.end()
        imgui.pop_font()
        
        imgui.end_frame()
        imgui.render()
        g_imguiImpl.render(imgui.get_draw_data())
        #===========================================

        glfw.swap_buffers(g_window)

def shutdown():
    g_imguiImpl.shutdown()
    glfw.terminate()

def main():
    initGlfwWindow()

    setOglDefaults()

    initImgui()

    registerGlfwCallbacks()

    loop()

    shutdown()

if __name__ == "__main__":
    main()