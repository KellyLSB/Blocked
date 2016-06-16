package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

// System Update() calls per second = fixedDeltaTime / 1.0.
const fixedDeltaTime float64 = 1.0 / 60.0

type Window struct {
	Title                        string
	Width, Height, FrameRate     int
	Fullscreen, Resizable, VSync bool

	*glfw.Window
	Keyboard
	Mouse

	callbacks struct {
		key             []glfw.KeyCallback
		framebufferSize []glfw.FramebufferSizeCallback
		cursorPos       []glfw.CursorPosCallback
		mouseButton     []glfw.MouseButtonCallback
		scroll          []glfw.ScrollCallback
		run             []func(*glfw.Window)
		draw            []func(*glfw.Window)
	}
}

func NewWindow(
	title string,
	width, height int,
	fullscreen, resizable, vsync bool,
) *Window {
	return &Window{
		Title:      title,
		Width:      width,
		Height:     height,
		Fullscreen: fullscreen,
		Resizable:  resizable,
		VSync:      vsync,
	}
}

func (w *Window) Run() {
	runtime.LockOSThread()

	// Init GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// GLFW WindowHints
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Start with window hidden
	glfw.WindowHint(glfw.Visible, glfw.False)

	// Get primary monitor video mode
	monitor := glfw.GetPrimaryMonitor()
	videoMode := monitor.GetVideoMode()

	// Sync window and monitor color bits
	glfw.WindowHint(glfw.RedBits, videoMode.RedBits)
	glfw.WindowHint(glfw.GreenBits, videoMode.GreenBits)
	glfw.WindowHint(glfw.BlueBits, videoMode.BlueBits)

	// Sync window and monitor refresh rate
	if w.VSync {
		glfw.WindowHint(glfw.RefreshRate, videoMode.RefreshRate)
	}

	// Is window resizable
	glfw.WindowHint(glfw.Resizable, func() int {
		if w.Resizable {
			return glfw.True
		}

		return glfw.False
	}())

	var err error

	// Create Window
	if w.Fullscreen {
		w.Window, err = glfw.CreateWindow(
			videoMode.Width, videoMode.Height,
			w.Title, monitor, nil,
		)
	} else {
		w.Window, err = glfw.CreateWindow(
			w.Width, w.Height,
			w.Title, nil, nil,
		)
	}

	// Panic on Error
	if err != nil {
		panic(err)
	}

	// Update Width and Height and create a callback
	w.Width, w.Height = w.Window.GetFramebufferSize()
	w.OnFramebufferSize(func(_ *glfw.Window, width, height int) {
		w.Width, w.Height = width, height
	})

	// Centre the window on the screen
	if !w.Fullscreen {
		w.Window.SetPos(
			videoMode.Width/2-w.Width/2,
			videoMode.Height/2-w.Height/2,
		)
	}

	// Set keyboard/mouse state callbacks
	w.OnKey(w.Keyboard.glfwKeyCallback)
	w.OnMouseButton(w.Mouse.glfwMouseButtonCallback)

	// Set Callbacks
	w.Window.SetScrollCallback(
		w.callScroll)
	w.Window.SetKeyCallback(
		w.callKey)
	w.Window.SetCursorPosCallback(
		w.callCursorPos)
	w.Window.SetMouseButtonCallback(
		w.callMouseButton)
	w.Window.SetFramebufferSizeCallback(
		w.callFramebufferSize)

	// Show the window and contextualize
	w.Window.Show()
	w.Window.MakeContextCurrent()

	// Initialize OpenGL.
	if err = gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	// If we wanted VSync, now's
	// the time to tell OpenGL about it.
	// @TODO: Replace comment
	if w.VSync {
		glfw.SwapInterval(1)
	}

	// Print some OpenGL information to stdout.
	fmt.Println("=== OpenGl Information ===")
	fmt.Println("Version: " + gl.GoStr(gl.GetString(gl.VERSION)))
	fmt.Println("Vendor: " + gl.GoStr(gl.GetString(gl.VENDOR)))
	fmt.Println("Renderer: " + gl.GoStr(gl.GetString(gl.RENDERER)))
	fmt.Println("GLSL Version: " + gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION)))
	fmt.Println("==========================")

	// Initialize any application specific stuff
	w.callRun(w.Window)

	// Preparing some variables to assist in keeping track of time, independent
	// of the system we're running on. By declaring these now we can prevent
	// creating potentially millions of objects for the garbage collector to
	// indulge upon.
	var newTime float64         // Subtracted from last currentTime = frameTime
	var frameTime float64       // Time elapsed since last frame
	var accumulator float64     // Time accumulated by calling Render()
	referenceTime := time.Now() // Reference time
	var frameCount int          // Number of frames, for determining framerate
	var frameCountReset float64 // Amount of time, for determining framerate

	// Begin the game loop
	currentTime := time.Since(referenceTime).Seconds()

	for !w.Window.ShouldClose() {
		// Get this frame's frameTime (the time elapsed since the last frame)
		newTime = time.Since(referenceTime).Seconds()
		frameTime = newTime - currentTime

		// Prevent the spiral-of-death-desynchronization that occurs between
		// updates and renders when our Render() calls take too long to
		// complete.
		if frameTime > 0.25 {
			frameTime = 0.25
		}

		// Update currentTime actually be the current time, so that we can
		// effectively measure the next frame's frameTime.
		currentTime = newTime

		// Add this frame's time to the accumulator. Once we've accumulated
		// enough time for an update (i.e. the amount of time specified by
		// fixedDeltaTime, probably 1.0/60.0s), we can call the update callback
		// on the game developer's systems.
		accumulator += frameTime

		// If we've accumulated enough time, i.e. the amount specified by
		// fixedDeltaTime, call Polygo's update callback to pass out a game
		// update to the game deveoper's systems.
		for accumulator >= fixedDeltaTime {
			//polygo.callUpdate(fixedDeltaTime)
			// @TODO: Determine what this is for
			accumulator -= fixedDeltaTime
		}

		// Increment frame count (number of frames which have passed since the
		// last frame count reset, for counting frames / second)
		frameCount++

		// Add this frame's frameTime to the frameCountReset (amount of time
		// which has passed since the last frame count reset, for counting
		// frames / second)
		frameCountReset += frameTime

		// Check when a full second has passed. When it has, we can look at
		// the number of frames which passed during that second to determine
		// a framerate (measured in frames per second).
		if frameCountReset >= 1 {
			w.FrameRate = frameCount
			frameCount = 0
			frameCountReset -= 1
		}

		// Clear the color & depth buffers so we don't see previously rendered
		// frames behind this frame.
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Trigger the Polygo callbacks system to call the Draw callbacks for
		// the game developer's own systems.
		w.callDraw(w.Window)

		// Switch the buffer we just rendered the game state to with the buffer
		// currently displayed on the screen. The currently displayed buffer
		// will then become the buffer we render to next time.
		w.Window.SwapBuffers()

		// Poll the GLFW for keyboard/mouse/etc events, which will then be
		// passed to the Polygo's keyboard + mouse objects, as well as the
		// callbacks system for triggering the game developer's systems.
		//
		// Currently this is coupled to the render loop. Investigate either
		// moving it to the update loop (call it BEFORE calling Update()s) or
		// completely decoupling it from both the render and update loops if
		// bad input performance is experienced on systems with a low framerate.
		glfw.PollEvents()
	}
}

func (w *Window) OnKey(
	cbs ...glfw.KeyCallback,
) {
	w.callbacks.key = append(
		w.callbacks.key, cbs...,
	)
}

func (w *Window) OnScroll(
	cbs ...glfw.ScrollCallback,
) {
	w.callbacks.scroll = append(
		w.callbacks.scroll, cbs...,
	)
}

func (w *Window) OnCursorPos(
	cbs ...glfw.CursorPosCallback,
) {
	w.callbacks.cursorPos = append(
		w.callbacks.cursorPos, cbs...,
	)
}

func (w *Window) OnMouseButton(
	cbs ...glfw.MouseButtonCallback,
) {
	w.callbacks.mouseButton = append(
		w.callbacks.mouseButton, cbs...,
	)
}

func (w *Window) OnFramebufferSize(
	cbs ...glfw.FramebufferSizeCallback,
) {
	w.callbacks.framebufferSize = append(
		w.callbacks.framebufferSize, cbs...,
	)
}

func (w *Window) OnRun(
	cbs ...func(window *glfw.Window),
) {
	w.callbacks.run = append(
		w.callbacks.run, cbs...,
	)
}

func (w *Window) OnDraw(
	cbs ...func(window *glfw.Window),
) {
	w.callbacks.draw = append(
		w.callbacks.draw, cbs...,
	)
}

func (w *Window) callKey(
	window *glfw.Window,
	key glfw.Key,
	scancode int,
	action glfw.Action,
	mods glfw.ModifierKey,
) {
	for _, cb := range w.callbacks.key {
		cb(window, key, scancode, action, mods)
	}
}

func (w *Window) callScroll(
	window *glfw.Window,
	xoff, yoff float64,
) {
	for _, cb := range w.callbacks.scroll {
		cb(window, xoff, yoff)
	}
}

func (w *Window) callCursorPos(
	window *glfw.Window,
	xpos float64,
	ypos float64,
) {
	for _, cb := range w.callbacks.cursorPos {
		cb(window, xpos, ypos)
	}
}

func (w *Window) callMouseButton(
	window *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mod glfw.ModifierKey,
) {
	for _, cb := range w.callbacks.mouseButton {
		cb(window, button, action, mod)
	}
}

func (w *Window) callFramebufferSize(
	window *glfw.Window,
	width, height int,
) {
	for _, cb := range w.callbacks.framebufferSize {
		cb(window, width, height)
	}
}

func (w *Window) callRun(window *glfw.Window) {
	for _, cb := range w.callbacks.run {
		cb(window)
	}
}

func (w *Window) callDraw(window *glfw.Window) {
	for _, cb := range w.callbacks.draw {
		cb(window)
	}
}
