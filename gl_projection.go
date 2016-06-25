package main

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

// Projection is a helper object
// to interact with a projection uniform.
type Projection struct {
	Object
}

// CProjection
// Cast a Location into a Projection utility object.
func CProjection(location interface{}) *Projection {
	switch location := location.(type) {
	case *Object:
		return &Projection{*location}
	case Object:
		return &Projection{location}
	case Location:
		return &Projection{*CObject(location)}
	default:
		panic(fmt.Errorf(
			"Bad type passed %# v; expected Location or Object", location,
		))
	}
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
	p.Set4(mgl32.Perspective(
		mgl32.DegToRad(degree),
		aspect, near, far,
	))

	p.Render(true)
}

// Zoom
// Performs a simple Affinine Scale
// on the projection; this provides
// an efficiency boost then moving the camera.
func (p *Projection) Zoom(scale float32) {
	p.M4.Set(0, 0, p.M4.At(0, 0)*scale)
	p.M4.Set(1, 1, p.M4.At(1, 1)*scale)
	p.M4.Set(2, 2, p.M4.At(2, 2)*scale)
	p.M4.Set(3, 3, p.M4.At(3, 3)*scale)
	p.Render(true)
}
