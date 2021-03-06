// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"errors"
	"math"
)

func Ortho(left, right, bottom, top, near, far float64) Mat4 {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)

	return Mat4{float64(2. / rml), 0, 0, 0, 0, float64(2. / tmb), 0, 0, 0, 0, float64(-2. / fmn), 0, float64(-(right + left) / rml), float64(-(top + bottom) / tmb), float64(-(far + near) / fmn), 1}
}

// Equivalent to Ortho with the near and far planes being -1 and 1, respectively
func Ortho2D(left, right, top, bottom float64) Mat4 {
	return Ortho(left, right, top, bottom, -1, 1)
}

func Perspective(fovy, aspect, near, far float64) Mat4 {
	// fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	nmf, f := near-far, float64(1./math.Tan(float64(fovy)/2.0))

	return Mat4{float64(f / aspect), 0, 0, 0, 0, float64(f), 0, 0, 0, 0, float64((near + far) / nmf), -1, 0, 0, float64((2. * far * near) / nmf), 0}
}

func Frustum(left, right, bottom, top, near, far float64) Mat4 {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	A, B, C, D := (right+left)/rml, (top+bottom)/tmb, -(far+near)/fmn, (2*far*near)/fmn

	return Mat4{float64((2. * near) / rml), 0, 0, 0, 0, float64((2. * near) / tmb), 0, 0, float64(A), float64(B), float64(C), -1, 0, 0, float64(D), 0}
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) Mat4 {
	F := Vec3{
		float64(centerX - eyeX),
		float64(centerY - eyeY),
		float64(centerZ - eyeZ)}

	f := F.Normalize()

	Up := Vec3{
		float64(upX),
		float64(upY),
		float64(upZ)}

	Upp := Up.Normalize()

	s := f.Cross(Upp).Normalize()
	u := s.Cross(f)

	M := Mat4{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(-eyeX, -eyeY, -eyeZ))
}

func LookAtV(eye, center, up Vec3) Mat4 {
	F := center.Sub(eye)

	f := F.Normalize()

	Upp := up.Normalize()

	s := f.Cross(Upp).Normalize()
	u := s.Cross(f)

	M := Mat4{s[0], u[0], -f[0], 0, s[1], u[1], -f[1], 0, s[2], u[2], -f[2], 0, 0, 0, 0, 1}

	return M.Mul4(Translate3D(float64(-eye[0]), float64(-eye[1]), float64(-eye[2])))
}

// Transform a set of coordinates from object space (in obj) to window coordinates (with depth)
//
// Window coordinates are continuous, not discrete (well, as continuous as an IEEE Floating Point can be), so you won't get exact pixel locations
// without rounding or similar
func Project(obj Vec3, modelview, projection Mat4, initialX, initialY, width, height int) (win Vec3) {
	obj4 := Vec4{obj[0], obj[1], obj[2], 1.0}

	vpp := projection.Mul4(modelview).Mul4x1(obj4)
	win[0] = float64(initialX) + (float64(width)*(vpp[0]+1))/2
	win[1] = float64(initialY) + (float64(height)*(vpp[1]+1))/2
	win[2] = (vpp[2] + 1) / 2

	return win
}

// Transform a set of window coordinates to object space. If your MVP (projection.Mul(modelview) matrix is not invertible, this will return an error
//
// Note that the projection may not be perfect if you use strict pixel locations rather than the exact values given by Projectf.
// (It's still unlikely to be perfect due to precision errors, but it will be closer)
func UnProject(win Vec3, modelview, projection Mat4, initialX, initialY, width, height int) (obj Vec3, err error) {
	inv := projection.Mul4(modelview).Inv()
	blank := Mat4{}
	if inv == blank {
		return Vec3{}, errors.New("Could not find matrix inverse (projection times modelview is probably non-singular)")
	}

	obj[0] = (2 * (win[0] - float64(initialX)) / float64(width)) - 1
	obj[1] = (2 * (win[1] - float64(initialY)) / float64(height)) - 1
	obj[2] = 2*win[2] - 1

	return obj, nil
}
