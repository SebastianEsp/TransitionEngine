package shaders

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func Compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	content, err := ioutil.ReadFile(source)
	if err != nil {
		return 0, fmt.Errorf("Error: %s", err)
	}
	csource, free := gl.Strs(string(content))
	gl.ShaderSource(shader, 1, csource, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}
	return shader, nil
}

func CreateShaderProgram(shaders []uint32) (uint32, error) {
	shaderProgram := gl.CreateProgram()
	for _, s := range shaders {
		gl.AttachShader(shaderProgram, s)
	}
	gl.LinkProgram(shaderProgram)

	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Shader program linking failed: %v", log)
	}

	for i := range shaders {
		gl.DeleteShader(shaders[i])
	}

	return shaderProgram, nil
}

func SetBool(shaderProgram uint32, name string, value bool) {
	gl.Uniform1i(gl.GetUniformLocation(shaderProgram, gl.Str(name)), btoi(value))
}

func SetInt(shaderProgram uint32, name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(shaderProgram, gl.Str(name)), value)
}

func SetFloat(shaderProgram uint32, name string, value float32) {
	gl.Uniform1f(gl.GetUniformLocation(shaderProgram, gl.Str(name)), value)
}

func btoi(b bool) int32 {
	if b {
		return 1
	}
	return 0
}
