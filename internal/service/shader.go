package service

import (
	alg "another-tiny-render/pkg/algebra"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type SOglShader struct {
	Type   uint32
	Handle uint32
}

type SOglProgram struct {
	Name   string
	Handle uint32

	ShaderLst []SOglShader
}

func (prog *SOglProgram) AppendShader(types []uint32, files []string) {
	if len(types) != len(files) {
		fmt.Println("SOglProgram.AppendShader(): ERROR! Wrong argument count!")
		panic(nil)
	} else {
		var shader SOglShader

		for i, fname := range files {
			file, _ := os.Open(fname)
			defer file.Close()

			src, _ := ioutil.ReadAll(file)

			shader.Type = types[i]
			shader.Handle = gl.CreateShader(types[i])

			glSrc, freeFn := gl.Strs(string(src) + "\x00")
			defer freeFn()

			gl.ShaderSource(shader.Handle, 1, glSrc, nil)
			gl.CompileShader(shader.Handle)

			var success int32
			gl.GetShaderiv(shader.Handle, gl.COMPILE_STATUS, &success)

			if success == gl.FALSE {
				var logLength int32
				gl.GetShaderiv(shader.Handle, gl.INFO_LOG_LENGTH, &logLength)

				log := gl.Str(strings.Repeat("\x00", int(logLength)))
				gl.GetShaderInfoLog(shader.Handle, logLength, nil, log)

				fmt.Printf("%s: %s", "AppendShader(): Error!", gl.GoStr(log))
			}

			prog.ShaderLst = append(prog.ShaderLst, shader)
		}
	}
}

func (prog *SOglProgram) LinkProgram() {
	prog.Handle = gl.CreateProgram()

	for _, shList := range prog.ShaderLst {
		gl.AttachShader(prog.Handle, shList.Handle)
	}

	gl.LinkProgram(prog.Handle)

	var success int32
	gl.GetProgramiv(prog.Handle, gl.LINK_STATUS, &success)

	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(prog.Handle, gl.INFO_LOG_LENGTH, &logLength)

		log := gl.Str(strings.Repeat("\x00", int(logLength)))
		gl.GetProgramInfoLog(prog.Handle, logLength, nil, log)

		fmt.Printf("%s: %s", "LinkProgram(): Error!", gl.GoStr(log))
	}
}

func (prog *SOglProgram) Use() {
	gl.UseProgram(prog.Handle)
}

func (prog *SOglProgram) PassMtrx2(name string, A alg.Mtrx2) {
	glName := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(prog.Handle, glName)
	gl.UniformMatrix2fv(location, 1, false, &A[0])
}

func (prog *SOglProgram) PassMtrx4(name string, A alg.Mtrx4) {
	glName := gl.Str(name + "\x00")
	location := gl.GetUniformLocation(prog.Handle, glName)
	gl.UniformMatrix4fv(location, 1, false, &A[0])
}

func (prog *SOglProgram) PassVertexAtribArray(count int32, name string) {
	location := uint32(gl.GetAttribLocation(prog.Handle, gl.Str(name+"\x00")))
	gl.EnableVertexAttribArray(uint32(location))
	gl.VertexAttribPointer(uint32(location), count, gl.FLOAT, false, 0, nil)
}
