package main

// Borrowed from:
// https://github.com/polygo/polygo/blob/master/mouse.go

import "github.com/go-gl/glfw/v3.1/glfw"

type MouseButton int
type Mouse [glfw.MouseButtonLast]bool

//Mouse structure manipulation
func (m *Mouse) glfwMouseButtonCallback(
	window *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	mod glfw.ModifierKey,
) {
	if action == glfw.Press {
		m[button] = true
	} else if action == glfw.Release {
		m[button] = false
	}
}

//Test if mouse button is down
func (m *Mouse) IsDown(button MouseButton) bool {
	return m[glfw.MouseButton(int(button))]
}

//Input constants
const (
	MouseButton1      MouseButton = MouseButton(int(glfw.MouseButton1))
	MouseButton2      MouseButton = MouseButton(int(glfw.MouseButton2))
	MouseButton3      MouseButton = MouseButton(int(glfw.MouseButton3))
	MouseButton4      MouseButton = MouseButton(int(glfw.MouseButton4))
	MouseButton5      MouseButton = MouseButton(int(glfw.MouseButton5))
	MouseButton6      MouseButton = MouseButton(int(glfw.MouseButton6))
	MouseButton7      MouseButton = MouseButton(int(glfw.MouseButton7))
	MouseButton8      MouseButton = MouseButton(int(glfw.MouseButton8))
	MouseButtonLast   MouseButton = MouseButton(int(glfw.MouseButtonLast))
	MouseButtonLeft   MouseButton = MouseButton(int(glfw.MouseButtonLeft))
	MouseButtonRight  MouseButton = MouseButton(int(glfw.MouseButtonRight))
	MouseButtonMiddle MouseButton = MouseButton(int(glfw.MouseButtonMiddle))
)
