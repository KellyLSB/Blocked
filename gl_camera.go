package main

import "github.com/go-gl/mathgl/mgl32"

// Camera is a helper object
// to interact with a camera uniform
type Camera struct {
	Object
}

// CCamera
// Cast a Location into a Camera utility object.
func CCamera(location Location) *Camera {
	return &Camera{*CObject(location)}
}

// LookAtV
// Creates a 4 dimensional matrix
// in the affinine style to represent the
// camera location and focal length
// per frame of view
func (c *Camera) LookAtV(
	eye, center, up mgl32.Vec3,
) {
	c.Set4(mgl32.LookAtV(eye, center, up))
	c.Render(true)
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
