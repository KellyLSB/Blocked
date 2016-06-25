package main

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// Model is a helper object
// to interact with a model uniform
type Model struct {
	Object
	Pitch, Roll, Yaw float32
}

// CModel - Cast a Location into a Model utility object.
func CModel(location Location) *Model {
	return &Model{Object: *CObject(location)}
}

// IncPitch - Increments the Model.Pitch angle.
func (m *Model) IncPitch(angle float32) {
	m.SetPitch(incrementAngle(m.Pitch, angle))
	fmt.Printf("Pitch: %f\n", m.Pitch)
}

// IncRoll - Increments the Model.Roll angle.
func (m *Model) IncRoll(angle float32) {
	m.SetRoll(incrementAngle(m.Roll, angle))
	fmt.Printf("Roll: %f\n", m.Roll)
}

// IncYaw - Increments the Model.Yaw angle.
func (m *Model) IncYaw(angle float32) {
	m.SetYaw(incrementAngle(m.Yaw, angle))
	fmt.Printf("Yaw: %f\n", m.Yaw)
}

func incrementAngle(angle, n float32) float32 {
	angle = angle + n

	if n = angle / 360.0; n > 1.0 {
		angle = angle - float32(
			math.Floor(float64(n)),
		)*360.0
	}

	return angle
}

// SetPitch - Set and Perform the Model.Pitch rotation.
//
// @NOTE: Setting the Model.Pitch accessor
// does not perform the rotation.
// The accessor is to used for transitions.
func (m *Model) SetPitch(angle float32) {
	m.Pitch = angle
	m.HomogRotate3D(m.Pitch, mgl32.Vec3{1, 0, 0})
}

// SetRoll - Set and Perform the Model.Roll rotation.
//
// @NOTE: Setting the Model.Roll accessor
// does not perform the rotation.
// The accessor is to used for transitions.
func (m *Model) SetRoll(angle float32) {
	m.Roll = angle
	m.HomogRotate3D(m.Roll, mgl32.Vec3{0, 1, 0})
}

// SetYaw - Set and Perform the Model.Yaw rotation.
//
// @NOTE: Setting the Model.Yaw accessor
// does not perform the rotation.
// The accessor is to used for transitions.
func (m *Model) SetYaw(angle float32) {
	m.Yaw = angle
	m.HomogRotate3D(m.Yaw, mgl32.Vec3{0, 0, 1})
}
