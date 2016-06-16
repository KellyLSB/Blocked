package main

import "github.com/go-gl/gl/v4.5-core/gl"

type Location uint32

func (l Location) UniformMatrix4fv(
	count int32, transpose bool, value *float32,
) {
	gl.UniformMatrix4fv(int32(l), count, transpose, value)
}

func (l Location) Uniform1I(v0 int32) {
	gl.Uniform1i(int32(l), v0)
}

func (l Location) EnableVertexAttribArray() {
	gl.EnableVertexAttribArray(uint32(l))
}

// VertexAttribPointer
//
// Arguments:
//   numDimensions - X,Y,Z = 3; X,Y,Z,U,V = 4; etc...
//   xtype         - gl library C type (i.e. `gl.FLOAT`)
//   normalized    - specifies if the normals are already calculated
//                   (i.e. If you are providing a `gl.NORMAL_ARRAY`;
//                    if I remember)
//   strideBytes   - Total bytes for all included dimensions.
//                   (i.e. X,Y,Z @ uint32 = 3*4 = 12)
//   pointer       - The `layout(location = n)` pointer qualifier
//                   https://www.opengl.org/wiki/Layout_Qualifier_(GLSL)
func (l Location) VertexAttribPointer(
	numDimensions int32,
	xtype uint32, normalized bool,
	strideBytes int32, pointer int,
) {
	gl.VertexAttribPointer(
		uint32(l), numDimensions,
		xtype, normalized, strideBytes,
		gl.PtrOffset(pointer),
	)
}
