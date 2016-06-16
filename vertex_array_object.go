package main

import (
	"unsafe"

	"github.com/go-gl/gl/v4.5-core/gl"
)

type VertexArrayObject uint32

func GenVertexArray() (vao VertexArrayObject) {
	gl.GenVertexArrays(1, (*uint32)(unsafe.Pointer(&vao)))
	return
}

func (vao VertexArrayObject) BindVertexArray() {
	gl.BindVertexArray(uint32(vao))
}
