package main

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type Program uint32

func (p Program) AttachShader(shader Shader) {
	gl.AttachShader(uint32(p), uint32(shader))
}

func (p Program) Link() {
	gl.LinkProgram(uint32(p))

	if !p.LinkStatus() {
		panic(errors.New(p.InfoLog()))
	}
}

func (p Program) IV(pname uint32) (value int32) {
	gl.GetProgramiv(uint32(p), pname, &value)
	return
}

func (p Program) LinkStatus() bool {
	return p.IV(gl.LINK_STATUS) == gl.TRUE
}

func (p Program) InfoLogLength() int32 {
	return p.IV(gl.INFO_LOG_LENGTH)
}

func (p Program) InfoLog() (log string) {
	len := p.InfoLogLength()
	log = strings.Repeat("\x00", int(len+1))
	gl.GetProgramInfoLog(uint32(p), len, nil, gl.Str(log))
	return
}

func (p Program) Use() {
	gl.UseProgram(uint32(p))
}

func (p Program) GetUniformLocation(attr string) Location {
	return Location(gl.GetUniformLocation(uint32(p), gl.Str(attr+"\x00")))
}

func (p Program) GetAttribLocation(attr string) Location {
	return Location(gl.GetAttribLocation(uint32(p), gl.Str(attr+"\x00")))
}

func (p Program) BindFragDataLocation(color uint32, attr string) {
	gl.BindFragDataLocation(uint32(p), color, gl.Str(attr+"\x00"))
}

func NewProgram(shaders ...Shader) (p Program) {
	p = Program(gl.CreateProgram())

	for _, shader := range shaders {
		p.AttachShader(shader)
	}

	p.Link()
	return
}
