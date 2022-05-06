
from OpenGL.GL import *
from collections import namedtuple

g_VS = """
#version 330 core
in vec3 position;
in vec3 color;

out vec3 newColor;
void main()
{
    gl_Position = vec4(position, 1.0f);
    newColor = color;
}
"""

g_FS = """
#version 330
in vec3 newColor;

out vec4 outColor;
void main()
{
    outColor = vec4(newColor, 1.0f);
}
"""

oglShader = namedtuple("oglShader", ["name", "type", "shaderObj"])
oglProgram = namedtuple("oglProgram", ["name", "type", "progObj"])

class COglProgramMngr:
    def __init__(self):
        self.shaderList = []
        self.programList = []

    def appendShader(self, name: str, type, src: str):
        shaderObj = glCreateShader(type)
                
        glShaderSource(shaderObj,  src)
        glCompileShader(shaderObj)
        printShaderInfoLog(shaderObj)

        self.shaderList.append(oglShader(name, type, shaderObj))

        # names - лист строк, имён шейдеров из shaderList
    def appendProgram(self, name: str, names: list[str]):
        prog = glCreateProgram()

        for sh_name in names:
            for sh in self.shaderList:
                if sh.name == sh_name:
                    glAttachShader(prog, sh.shaderObj)
            
        glLinkProgram(prog)
        printProgramInfoLog(prog)

        self.programList.append(oglProgram(name, 0, prog))

    def getProgObj(self, name: str):
        pass

def printOglError() -> int:
    err = glGetError()

    retCode = 0

    while (err != GL_NO_ERROR):
        if err==GL_INVALID_ENUM:
                error = "INVALID_ENUM" 
        if err==GL_INVALID_VALUE:
                error = "INVALID_VALUE"
        if err==GL_INVALID_OPERATION:
                error = "INVALID_OPERATION"
        if err==GL_STACK_OVERFLOW:
                error = "STACK_OVERFLOW"
        if err==GL_STACK_UNDERFLOW:
                error = "STACK_UNDERFLOW"
        if err==GL_OUT_OF_MEMORY:
                error = "OUT_OF_MEMORY"
        if err==GL_INVALID_FRAMEBUFFER_OPERATION:
                error = "INVALID_FRAMEBUFFER_OPERATION"
        print(error)
        retCode = 1
        err = glGetError()

    return retCode

def printShaderInfoLog(obj: GLuint) -> None:
    infologLength = glGetShaderiv(obj, GL_INFO_LOG_LENGTH)
 
    if (infologLength > 0):
        infoLog = glGetShaderInfoLog(obj, infologLength)
        print(infoLog)

def printProgramInfoLog(obj: GLuint) -> None:
    infologLength = 0
 
    infologLength = glGetProgramiv(obj, GL_INFO_LOG_LENGTH)
 
    if (infologLength > 0):
        infoLog = glGetProgramInfoLog(obj, infologLength)
        print(infoLog)
