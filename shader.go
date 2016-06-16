package main

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type Shader uint32

func glStr(str string, fn func(str **uint8)) {
	cstr, free := gl.Strs(str)
	defer free()
	fn(cstr)
}

func CompileShader(
	shaderType uint32,
	source string,
) (s Shader) {
	s = Shader(gl.CreateShader(shaderType))
	s.Source(source)
	s.Compile()
	return
}

func (s Shader) Compile() {
	gl.CompileShader(uint32(s))

	if !s.CompileStatus() {
		panic(errors.New(s.InfoLog()))
	}
}

func (s Shader) Source(source string) {
	glStr(source, func(source **uint8) {
		gl.ShaderSource(uint32(s), 1, source, nil)
	})
}

func (s Shader) IV(pname uint32) (value int32) {
	gl.GetShaderiv(uint32(s), pname, &value)
	return
}

func (s Shader) CompileStatus() bool {
	return s.IV(gl.COMPILE_STATUS) == gl.TRUE
}

func (s Shader) InfoLogLength() int32 {
	return s.IV(gl.INFO_LOG_LENGTH)
}

func (s Shader) InfoLog() (log string) {
	len := s.InfoLogLength()
	log = strings.Repeat("\x00", int(len+1))
	gl.GetShaderInfoLog(uint32(s), len, nil, gl.Str(log))
	return
}
