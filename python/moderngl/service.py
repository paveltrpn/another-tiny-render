import sys
import OpenGL
OpenGL.ERROR_CHECKING = False
from OpenGL.GL import * #type: ignore
from glfw.GLFW import *
from dataclasses import dataclass

@dataclass
class CAppState(object):
    def __init__(self):
        self.wndWidth: int      = 1152
        self.wndHeight: int     = 864
        self.appName: str       = "003"
        self.glfwWndPtr         = None

        self.glfwVersionStr: str = ""
        self.glRenderStr: str       = ""
        self.glVersionStr: str      = ""
        self.glslVersionStr: str    = ""
        self.glExtList: list[str] = []

    def printAppStats(self):
        print("GLFW info:")
        print(self.glfwVersionStr)

        print("OpenGL info:")
        print(self.glRenderStr)
        print(self.glVersionStr)
        print(self.glslVersionStr)

    def printExtensionsList(self):
        if self.glExtList:
            for ext in self.glExtList:
                print(ext)


def initGlfwWindow(appState: CAppState) -> None:
    if not glfwInit():
        print("initGlfwWindow(): ERROR! Can't initialize glfw!")
        sys.exit()

    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 4)
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 1)
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE)
    glfwWindowHint(GLFW_DOUBLEBUFFER, GLFW_TRUE)
    glfwWindowHint(GLFW_RESIZABLE,    GLFW_FALSE)

    appState.glfwWndPtr = glfwCreateWindow(appState.wndWidth, appState.wndHeight, appState.appName, None, None)
    if not appState.glfwWndPtr:
        print("initGlfwWindow(): Failed to create GLFW window!\n")
        glfwTerminate()
        sys.exit()

    def errorCallback(int, err_str: str):
        print("initGlfwWindow(): GLFW Error: {}".format(err_str))
    glfwSetErrorCallback(errorCallback)

    glfwMakeContextCurrent(appState.glfwWndPtr)

    [major, minor, rev] = glfwGetVersion()
    appState.glfwVersionStr = "{}.{}.{}".format(major, minor, rev)

    #  ВКЛ-ВЫКЛ вертикальную синхронизацию (VSYNC)
	#  Лок на 60 фпс
    glfwSwapInterval(GLFW_TRUE)

    appState.glRenderStr = str(glGetString(GL_RENDERER), "utf-8") #type: ignore
    appState.glVersionStr = str(glGetString(GL_VERSION), "utf-8") #type: ignore
    appState.glslVersionStr = str(glGetString(GL_SHADING_LANGUAGE_VERSION), "utf-8") #type: ignore

    extNum = glGetIntegerv(GL_NUM_EXTENSIONS)
    for curExt in range(extNum):
        appState.glExtList.append(str(glGetStringi(GL_EXTENSIONS, curExt), "utf-8")) #type: ignore

def registerGlfwCallbacks(appState: CAppState) -> None:
    def keyCallback(window, key, scancode, action, mods):
        if (key == GLFW_KEY_ESCAPE):
            glfwSetWindowShouldClose(window, GLFW_TRUE)

        if (key == GLFW_KEY_Q):
            print("Q is pressed!")

    glfwSetKeyCallback(appState.glfwWndPtr, keyCallback)
