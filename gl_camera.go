package main

import "github.com/go-gl/mathgl/mgl32"

// Camera is a helper object
// to interact with a camera uniform
type Camera struct {
	Location
	m4 mgl32.Mat4
}

// CCamera
// Cast a Location into a Camera utility object.
func CCamera(location Location) *Camera {
	return &Camera{Location: location}
}

// LookAtV
// Creates a 4 dimensional matrix
// in the affinine style to represent the
// camera location and focal length
// per frame of view
func (c *Camera) LookAtV(
	eye, center, up mgl32.Vec3,
) {
	c.m4 = mgl32.LookAtV(eye, center, up)
	c.UniformMatrix4fv(1, false)
}

// @TODO not sure if I want
// to include function to move
// the cameras... I'm unsure how
// that would play out sticking to
// a move the perspective angle...
//
// This is something worth figuring out.
// Especially if you consider how this
// applies to 3D printing.
//func (c *Camera) Pan(x, y float32)

// UniformMatrix4fv
// Performs the same as
// *Location.UniformMatrix4fv(); except
// this one already applies the
// camera model for you.
func (c *Camera) UniformMatrix4fv(
	count int32, transpose bool,
) {
	c.Location.UniformMatrix4fv(
		count, transpose, &c.m4[0],
	)
}
