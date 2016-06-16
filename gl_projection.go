package main

import "github.com/go-gl/mathgl/mgl32"

// Projection is a helper Object
// to interact with the projection uniform.
type Projection struct {
	Location
	p mgl32.Mat4
}

// CProjection
// Cast a location into a Projection Utility Object
func CProjection(location Location) *Projection {
	return &Projection{Location: location}
}

// Perspective
// Creates a 4 dimensional matrix
// in the affinine style to represent
// the distance between the observer
// and the object in view.
//
// Then applies with a UniformMatrix4fv() call.
func (p *Projection) Perspective(
	degree, aspect, near, far float32,
) {
	p.p = mgl32.Perspective(
		mgl32.DegToRad(degree),
		aspect, near, far,
	)

	p.UniformMatrix4fv(1, false)
}

// Zoom
// Performs a simple Affinine Scale
// on the projection; this provides
// an efficiency boost then moving the camera.
func (p *Projection) Zoom(scale float32) {
	// @TODO Will the third and fourth
	// following the decreasing linear
	// series of changes provide a better render?
	p.p.Set(0, 0, p.p.At(0, 0)*scale)
	p.p.Set(1, 1, p.p.At(1, 1)*scale)

	p.UniformMatrix4fv(1, false)
}

// UniformMatrix4fv
// Performs the same as
// *Location.UniformMatrix4fv(); except
// this one already applies the
// projection model for you.
func (p *Projection) UniformMatrix4fv(
	count int32, transpose bool,
) {
	p.Location.UniformMatrix4fv(
		count, transpose, &p.p[0],
	)
}
