// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Renders a textured spinning cube using GLFW 3.1 and OpenGL 4.1 core forward-compatible profile.
package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const windowWidth = 800
const windowHeight = 600

func main() {
	window := NewWindow("Block", windowWidth, windowHeight, false, true, true)

	window.OnRun(func(_ *glfw.Window) {
		// Configure the vertex and fragment shaders
		vShader := CompileShader(gl.VERTEX_SHADER, vertexShader)
		fShader := CompileShader(gl.FRAGMENT_SHADER, fragmentShader)
		program := NewProgram(vShader, fShader)
		program.Use()

		// Projection load uniform
		projection := CProjection(
			program.GetUniformLocation("projection"),
		)

		// Projection set perspective
		projection.Perspective(
			45.0, float32(
				window.Width/window.Height,
			), 0.1, 100.0,
		)

		// Projection zoom on scroll
		window.OnScroll(func(
			_ *glfw.Window,
			xoff, yoff float64,
		) {
			var scale float32

			if yoff > 0 {
				scale = 0.95
			} else {
				scale = 1.05
			}

			projection.Zoom(scale)
		})

		camera := CCamera(
			program.GetUniformLocation("camera"),
		)

		camera.LookAtV(
			mgl32.Vec3{3, 3, 3},
			mgl32.Vec3{0, 0, 0},
			mgl32.Vec3{0, 1, 0},
		)

		model := mgl32.Ident4()
		var modelYaw, modelPitch float32
		modelUniform := program.GetUniformLocation("model")
		modelUniform.UniformMatrix4fv(1, false, &model[0])

		window.OnCursorPos(func(_ *glfw.Window, xpos, ypos float64) {
			if window.Mouse.IsDown(MouseButtonRight) {
				midx := float64(window.Width >> 1)
				midy := float64(window.Height >> 1)

				if midx == xpos && midy == ypos {
					return
				}

				window.SetCursorPos(midx, midy)
				modelYaw += float32((midx-xpos)/1000) * 2
				modelPitch += float32((midy-ypos)/1000) * 2

				model = mgl32.HomogRotate3D(modelYaw, mgl32.Vec3{0, 1, 0}).
					Mul4(mgl32.HomogRotate3D(modelPitch, mgl32.Vec3{0, 0, 1}))
				modelUniform.UniformMatrix4fv(1, false, &model[0])
			}
		})

		textureUniform := program.GetUniformLocation("tex")
		textureUniform.Uniform1I(0)

		program.BindFragDataLocation(0, "outputColor")

		// Load the texture
		// texture, err := newTexture("square.png")
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// Configure the vertex data
		vao := GenVertexArray()
		vao.BindVertexArray()

		stl := OpenSTL("model.stl")
		stl.Scale(0.25)
		vertices := stl.Vertices()
		fmt.Println(vertices[0:20])

		vbo := GenBuffer(gl.ARRAY_BUFFER)
		vbo.BufferData(len(vertices)*4, vertices, gl.STATIC_DRAW)

		vertAttrib := program.GetAttribLocation("vert")
		vertAttrib.EnableVertexAttribArray()
		vertAttrib.VertexAttribPointer(3, gl.FLOAT, false, 5*4, 0)

		texCoordAttrib := program.GetAttribLocation("vertTexCoord")
		texCoordAttrib.EnableVertexAttribArray()
		texCoordAttrib.VertexAttribPointer(2, gl.FLOAT, false, 5*4, 3*4)

		// Configure global settings
		gl.Enable(gl.DEPTH_TEST)
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(gl.BACK)
		gl.DepthFunc(gl.LESS)
		gl.ClearColor(1.0, 1.0, 1.0, 1.0)

		angle := 0.0
		previousTime := glfw.GetTime()

		window.OnDraw(func(_ *glfw.Window) {

			// Update
			time := glfw.GetTime()
			elapsed := time - previousTime
			previousTime = time

			angle += elapsed
			//model = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

			// Render
			program.Use()
			//modelUniform.UniformMatrix4fv(1, false, &model[0])

			vao.BindVertexArray()

			//gl.ActiveTexture(gl.TEXTURE0)
			//gl.BindTexture(gl.TEXTURE_2D, texture)

			gl.DrawArrays(gl.TRIANGLES, 0, 6*2*3)
		})
	})

	window.Run()
}

func newTexture(file string) (uint32, error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return 0, fmt.Errorf("texture %q not found on disk: %v", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return 0, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return 0, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.GenTextures(1, &texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return texture, nil
}

var vertexShader = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var fragmentShader = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,

	// Top
	-1.0, 1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 1.0,

	// Left
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 1.0, 0.0,

	// Right
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, -1.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0,
}

// Set the working directory to the root of Go package, so that its assets can be accessed.
// func init() {
// 	dir, err := importPathToDir("github.com/go-gl/examples/glfw31-gl41core-cube")
// 	if err != nil {
// 		log.Fatalln("Unable to find Go package in your GOPATH, it's needed to load assets:", err)
// 	}
// 	err = os.Chdir(dir)
// 	if err != nil {
// 		log.Panicln("os.Chdir:", err)
// 	}
// }
//
// // importPathToDir resolves the absolute path from importPath.
// // There doesn't need to be a valid Go package inside that import path,
// // but the directory must exist.
// func importPathToDir(importPath string) (string, error) {
// 	p, err := build.Import(importPath, "", build.FindOnly)
// 	if err != nil {
// 		return "", err
// 	}
// 	return p.Dir, nil
// }
