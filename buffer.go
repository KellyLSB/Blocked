package main

import "github.com/go-gl/gl/v4.5-core/gl"

type Buffer [2]uint32

func GenBuffer(bufType uint32) (b Buffer) {
	b = Buffer{bufType, 0}
	gl.GenBuffers(1, &b[1])
	return
}

func (b Buffer) BindBuffer() {
	gl.BindBuffer(b[0], b[1])
}

func (b Buffer) BufferData(
	size int, data interface{}, usage uint32,
) {
	b.BindBuffer()
	gl.BufferData(b[0], size, gl.Ptr(data), usage)
}
