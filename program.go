package main

import (
	"errors"
	"strings"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type WindowProgram struct {
	window *Window
	Program
}

func NewWindowProgram(
	window *Window,
	shaders ...Shader,
) *WindowProgram {
	return &WindowProgram{
		window, NewProgram(shaders...),
	}
}

func (p *WindowProgram) GetUniformObject(attr string) (o *Object) {
	o = CObject(p.Program.GetUniformLocation(attr))
	o.EnableAutoRender()
	p.window.OnDraw(o.onDraw)
	return
}

func (p *WindowProgram) GetAttribObject(attr string) (o *Object) {
	o = CObject(p.Program.GetAttribLocation(attr))
	o.EnableAutoRender()
	p.window.OnDraw(o.onDraw)
	return
}

func (p *WindowProgram) Use() *WindowProgram {
	p.Program.Use()
	return p
}

type Program uint32

// NewProgram creates a new Program
func NewProgram(shaders ...Shader) (p Program) {
	p = Program(gl.CreateProgram())

	if len(shaders) > 0 {
		p.AttachShader(shaders...)
		p.Link()
		p.DetachShader(shaders...)
	}

	return
}

// AttachShader attaches a Shader to the Program
func (p Program) AttachShader(shaders ...Shader) {
	for _, shader := range shaders {
		gl.AttachShader(uint32(p), uint32(shader))
	}
}

// AttachShader Detaches a Shader from the Program
func (p Program) DetachShader(shaders ...Shader) {
	for _, shader := range shaders {
		gl.DetachShader(uint32(p), uint32(shader))
	}
}

func (p Program) Link() Program {
	gl.LinkProgram(uint32(p))

	if !p.LinkStatus() {
		panic(errors.New(p.InfoLog()))
	}

	return p
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

func (p Program) Use() Program {
	gl.UseProgram(uint32(p))
	return p
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

// func (p Program) GetProgramBinary() {
// 	gl.GetProgramBinary(program, bufSize, length, binaryFormat, binary)
// }
