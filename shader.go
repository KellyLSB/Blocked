package main

import (
	"errors"
	"io/ioutil"
	"strings"
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
)

func glStr(str string, fn func(str **uint8)) {
	cstr, free := gl.Strs(str)
	defer free()
	fn(cstr)
}

type ShaderType uint32

const (
	VertexShader                ShaderType = gl.VERTEX_SHADER
	TesselationControlShader    ShaderType = gl.TESS_CONTROL_SHADER
	TesselationEvaluationShader ShaderType = gl.TESS_EVALUATION_SHADER
	GeometryShader              ShaderType = gl.GEOMETRY_SHADER
	FragmentShader              ShaderType = gl.FRAGMENT_SHADER
	ComputeShader               ShaderType = gl.COMPUTE_SHADER
)

type Shader uint32

func NewShader(typ ShaderType) (s Shader) {
	return Shader(gl.CreateShader(uint32(typ)))
}

func (s Shader) Source(source string) Shader {
	glStr(source, func(source **uint8) {
		gl.ShaderSource(uint32(s), 1, source, nil)
	})

	return s
}

func (s Shader) SourceFile(file string) Shader {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return s.Source(*(*string)(unsafe.Pointer(&byt)))
}

func (s Shader) Compile() Shader {
	gl.CompileShader(uint32(s))

	if !s.CompileStatus() {
		panic(errors.New(s.InfoLog()))
	}

	return s
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
