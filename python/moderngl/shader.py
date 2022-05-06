
import sys
import os
from typing import Any
from OpenGL.GL import * #type: ignore
from collections import namedtuple

import algebra_py as alg

oglShader = namedtuple("oglShader",[
                        "name", 
                        "type",
                        "shaderObj"
                        ])

oglProgram = namedtuple("oglProgram",[
                        "name",
                        "progObj"
                        ])


class COglProgramMngr:
    shaderList: list[oglShader]
    programList: list[oglProgram]

    def __init__(self):
        self.shaderList = []
        self.programList = []

    def appendShader(self, type: Any, fname: str):
        with open(fname, "r") as f:
            content = f.readlines()
            src = " ".join(content)

        shaderObj = glCreateShader(type)
                
        glShaderSource(shaderObj, src)
        glCompileShader(shaderObj)

        infoLog = glGetShaderInfoLog(shaderObj)
        if (len(infoLog) > 0):
            print("COglProgramMngr.appendShader(): Error! Calling sys.terminate()")
            lst = str(infoLog).replace("b'","").replace("'","").split("\\n")
            print(*lst, sep='\n')

        file = os.path.basename(fname)   # "c:/foo/bar/name.type" -> "name.type"
        name = os.path.splitext(file)[0] # "name.type" -> "name"
        self.shaderList.append(oglShader(name, type, shaderObj))

    def appendProgram(self, name: str, names: list[str]):
        prog = glCreateProgram()

        for shName in names:
            flag = False
            for shList in self.shaderList:
                if shList.name == shName:
                    glAttachShader(prog, shList.shaderObj)
                    flag = True
            
            if flag == False:
                print("COglProgramMngr.appendProgram(): Error! Can't find a shader - {}".format(shName))
                sys.exit()
            
            
        glLinkProgram(prog)

        infoLog = glGetProgramInfoLog(prog)
        if (len(infoLog) > 0 ):
            print("COglProgramMngr.appendProgram(): Error! Calling sys.terminate()")
            lst = str(infoLog).replace("b'","").replace("'","").split("\\n")
            print(*lst, sep='\n')
            sys.exit()

        self.programList.append(oglProgram(name, prog))

    def getProgObj(self, name: str) -> Any:
        for prog in self.programList:
            if prog.name == name:
                return prog.progObj
        print("COglProgramMngr.getProgObj(): Error! Can't find program - {}".format(name))

    def __getitem__(self, key: str) -> Any:
        return self.getProgObj(key)

    def passUniformMatrix4f(self, progName: str, uniformName: str, A: alg.mtrx4) -> None:
        uniformLocation = glGetUniformLocation(self.getProgObj(progName), uniformName)
        glUniformMatrix4fv(uniformLocation, 1, GL_FALSE, A.tobytes())
    
    def passUniform1f(self, progName: str, uniformName: str, param: float) -> None:
        uniformLocation = glGetUniformLocation(self.getProgObj(progName), uniformName)
        glUniform1fv(uniformLocation, 1, param)
    

class COglProgram(object):
    programName: str
    programObject: Any

    shaderList: list[oglShader]

    def __init__(self, pname: str):
        self.shaderList = []
        self.programName = pname

    def __del__(self):
        for sh in self.shaderList:
            glDeleteShader(sh.shaderObj)
        glDeleteProgram(self.programObject)

    def appendShader(self, types: list[Any], files: list[str]) -> None:
        """ Add shaders to list for future program linking.

            Parameters:
                types - list of shader types, i.e - GL_VERTEX_SHADER, GL_FRAGMENT_SHADER etc.

                files - list of source file paths.

            Length of both types[] and files[] must be equal.
            Each element of types[] corresponds to element of files[].
        """
        if len(types) != len(files):
            print("COglProgram.appendShader(): ERROR! Wrong argument count!")
            sys.exit()
        else:
            for i, fname in enumerate(files):
                with open(fname, "r") as f:
                    content = f.readlines()
                    src = " ".join(content)

                shaderObj = glCreateShader(types[i])

                glShaderSource(shaderObj, src)
                glCompileShader(shaderObj)
                success = glGetShaderiv(shaderObj, GL_COMPILE_STATUS)
                if (not success):
                    infoLog = glGetShaderInfoLog(shaderObj)
                    print("COglProgram.appendShader(): can't compile shader - {}\n".format(fname) + infoLog.decode())

                file = os.path.basename(fname)   # "c:/foo/bar/name.type" -> "name.type"
                name = os.path.splitext(file)[0] # "name.type" -> "name"
                self.shaderList.append(oglShader(name, types[i], shaderObj))

    def linkProgram(self) -> None:
        self.programObject = glCreateProgram()

        for shList in self.shaderList:
            glAttachShader(self.programObject, shList.shaderObj)
            
        glLinkProgram(self.programObject)
        success = glGetProgramiv(self.programObject, GL_LINK_STATUS)
        if (not success):
            infoLog = glGetProgramInfoLog(self.programObject)
            print("COglProgram.linkProgram(): can't link program - {}\n".format(self.programName) + infoLog.decode())
            sys.exit()

    def use(self) -> None:
        glUseProgram(self.programObject)

    def passMat2(self, name: str, A: alg.mtrx2) -> None:
        glUniformMatrix2fv(glGetUniformLocation(self.programObject, name), 1, GL_FALSE, A.tobytes())

    def passMat3(self, name: str, A: alg.mtrx3) -> None:
        glUniformMatrix3fv(glGetUniformLocation(self.programObject, name), 1, GL_FALSE, A.tobytes())

    def passMtrx4(self, uniformName: str, A: alg.mtrx4) -> None:
        uniformLocation = glGetUniformLocation(self.programObject, uniformName)
        glUniformMatrix4fv(uniformLocation, 1, GL_FALSE, A.tobytes())
    
    def passBool(self, name: str, value: bool) -> None:
        glUniform1i(glGetUniformLocation(self.programObject, name), int(value))

    def passInt(self, name: str, value: int) -> None:
        glUniform1i(glGetUniformLocation(self.programObject, name), value)

    def passFloat(self, uniformName: str, value: float) -> None:
        uniformLocation = glGetUniformLocation(self.programObject, uniformName)
        glUniform1fv(uniformLocation, 1, value)

    def passVec2(self, name: str, *args) -> None:
        if (len(args) == 1 and type(args[0]) == alg.vec2):
            glUniform2fv(glGetUniformLocation(self.programObject, name), 1, args[0].tobytes())
        elif (len(args) == 2 and all(map(lambda x: type(x) == float, args))):
            glUniform2f(glGetUniformLocation(self.programObject, name), *args)

    def passVec3(self, name: str, *args) -> None:
        if (len(args) == 1 and type(args[0]) == alg.vec3):
            glUniform3fv(glGetUniformLocation(self.programObject, name), 1, args[0].tobytes())
        elif (len(args) == 3 and all(map(lambda x: type(x) == float, args))):
            glUniform3f(glGetUniformLocation(self.programObject, name), *args)

    def passVec4(self, name: str, *args) -> None:
        if (len(args) == 1 and type(args[0]) == alg.vec4):
            glUniform4fv(glGetUniformLocation(self.programObject, name), 1, args[0].tobytes())
        elif (len(args) == 3 and all(map(lambda x: type(x) == float, args))):
            glUniform4f(glGetUniformLocation(self.programObject, name), *args)

    def passVertexAtribArray(self, count: int, attrName: str):
        attrLocation = glGetAttribLocation(self.programObject, attrName) 
        glVertexAttribPointer(attrLocation, count, GL_FLOAT, GL_FALSE, 0, ctypes.c_void_p(0)) 
        glEnableVertexAttribArray(attrLocation)


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
        print(error) #type: ignore 
        retCode = 1
        err = glGetError()

    return retCode
