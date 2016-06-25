package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/go-gl/glfw/v3.1/glfw"

	"github.com/go-gl/mathgl/mgl32"
)

// Object is a helper structure
// for interacting with shader pointers
type Object struct {
	Location
	M4      mgl32.Mat4
	m4Mutex sync.RWMutex

	AutoRender    bool
	renderReady   bool
	RenderTimeout time.Duration
	renderTimeout *time.Timer
}

// CObject - Cast a Location into
// an Object utility structure.
func CObject(location Location) (o *Object) {
	return &Object{
		Location:      location,
		RenderTimeout: time.Millisecond * 500,
	}
}

// Set4 - Sets the Object.M4 matrix.
func (o *Object) Set4(m4 mgl32.Mat4) {
	o.m4Mutex.Lock()
	defer o.m4Mutex.Unlock()
	o.M4 = m4
}

func (o *Object) EnableAutoRender() {
	o.AutoRender = true
	o.renderTimeout = time.AfterFunc(
		o.RenderTimeout,
		o.markRenderReady,
	)

	o.renderTimeout.Stop()
}

// SetRenderTimeout - Adds a timeout on the Object.Render function.
// the idea is to interupt renders so multiple translations may be applied
// before asking OpenGL to draw a frame.
func (o *Object) SetRenderTimeout(timeout time.Duration) {
	o.RenderTimeout = timeout
}

// Ident4 - Initalize Object with a 4x4 Identity Matrice.
func (o *Object) Ident4() {
	o.Set4(mgl32.Ident4())
}

// Mul4 - Perform a 4 dimensional multiplication.
// Useful for affinine translations.
func (o *Object) Mul4(m4 mgl32.Mat4) {
	o.Set4(o.M4.Mul4(m4))
}

// HomogRotate3D - Performs a homogonous 3D
// rotation using a 4D affinine matrix.
func (o *Object) HomogRotate3D(angle float32, vector mgl32.Vec3) {
	o.Mul4(mgl32.HomogRotate3D(angle, vector))
}

// UniformMatrix4fv - Performs the same as
// *Location.UniformMatrix4fv(); except
// this one already applies the matrix for you.
func (o *Object) UniformMatrix4fv(
	count int32, transpose bool,
) {
	o.m4Mutex.RLock()
	defer o.m4Mutex.RUnlock()
	o.Location.UniformMatrix4fv(
		count, transpose, &o.M4[0],
	)
}

// Render - Renders a 4 dimensional matrix (fv = foreground? vector?)
func (o *Object) Render(force ...bool) {
	fmt.Println("o.Render()")
	if o.renderTimeout == nil || !o.AutoRender || len(force) > 0 && force[0] {
		o.render()
		return
	}

	o.renderTimeout.Reset(o.RenderTimeout)
}

func (o *Object) render() {
	runtime.LockOSThread()
	fmt.Println("o.render()")
	o.UniformMatrix4fv(1, false)
}

func (o *Object) onDraw(_ *glfw.Window) {
	if o.renderReady {
		o.render()
		o.renderReady = false
	}
}

func (o *Object) markRenderReady() {
	o.renderReady = true
}
