// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"math"
)

// A rotation order is the order in which
// rotations will be transformed for the purposes of AnglesToQuat
type RotationOrder int

const (
	XYX RotationOrder = iota
	XYZ
	XZX
	XZY
	YXY
	YXZ
	YZY
	YZX
	ZYZ
	ZYX
	ZXZ
	ZXY
)

// A Quaternion is an extension of the imaginary numbers; there's all sorts of
// interesting theory behind it. In 3D graphics we mostly use it as a cheap way of
// representing rotation since quaternions are cheaper to multiply by, and easier to
// interpolate than matrices.
//
// A Quaternion has two parts: W, the so-called scalar component,
// and "V", the vector component. The vector component is considered to
// be the part in 3D space, while W (loosely interpreted) is its 4D coordinate.
type Quat struct {
	W float32
	V Vec3
}

// The quaternion identity: W=1; V=(0,0,0).
//
// As with all identities, multiplying any quaternion by this will yield the same
// quaternion you started with.
func QuatIdent() Quat {
	return Quat{1., Vec3{0, 0, 0}}
}

// Creates an angle from an axis and an angle relative to that axis.
//
// This is cheaper than HomogRotate3D.
func QuatRotate(angle float32, axis Vec3) Quat {
	// angle = (float32(math.Pi) * angle) / 180.0

	c, s := float32(math.Cos(float64(angle/2))), float32(math.Sin(float64(angle/2)))

	return Quat{c, axis.Mul(s)}
}

// This function is deprecated. Instead, use AnglesToQuat
//
// The behavior of this function should be equivalent to AnglesToQuat(zAngle, yAngle, xAngle, ZYX)
func EulerToQuat(xAngle, yAngle, zAngle float32) Quat {
	sinz, cosz := math.Sincos(float64(zAngle))
	sz, cz := float32(sinz), float32(cosz)

	siny, cosy := math.Sincos(float64(yAngle))
	sy, cy := float32(siny), float32(cosy)

	sinx, cosx := math.Sincos(float64(xAngle))
	sx, cx := float32(sinx), float32(cosx)

	return Quat{
		W: cx*cy*cz + sx*sy*sz,
		V: Vec3{
			sx*cy*cz - cx*sy*sz,
			cx*sy*cz + sx*cy*sz,
			cx*cy*sz - sx*sy*cz,
		},
	}
}

// A convenient alias for q.V[0]
func (q Quat) X() float32 {
	return q.V[0]
}

// A convenient alias for q.V[1]
func (q Quat) Y() float32 {
	return q.V[1]
}

// A convenient alias for q.V[2]
func (q Quat) Z() float32 {
	return q.V[2]
}

// Adds two quaternions. It's no more complicated than
// adding their W and V components.
func (q1 Quat) Add(q2 Quat) Quat {
	return Quat{q1.W + q2.W, q1.V.Add(q2.V)}
}

// Subtracts two quaternions. It's no more complicated than
// subtracting their W and V components.
func (q1 Quat) Sub(q2 Quat) Quat {
	return Quat{q1.W - q2.W, q1.V.Sub(q2.V)}
}

// Multiplies two quaternions. This can be seen as a rotation. Note that
// Multiplication is NOT commutative, meaning q1.Mul(q2) does not necessarily
// equal q2.Mul(q1).
func (q1 Quat) Mul(q2 Quat) Quat {
	return Quat{q1.W*q2.W - q1.V.Dot(q2.V), q1.V.Cross(q2.V).Add(q2.V.Mul(q1.W)).Add(q1.V.Mul(q2.W))}
}

// Scales every element of the quaternion by some constant factor.
func (q1 Quat) Scale(c float32) Quat {
	return Quat{q1.W * c, Vec3{q1.V[0] * c, q1.V[1] * c, q1.V[2] * c}}
}

// Returns the conjugate of a quaternion. Equivalent to
// Quat{q1.W, q1.V.Mul(-1)}
func (q1 Quat) Conjugate() Quat {
	return Quat{q1.W, q1.V.Mul(-1)}
}

// Returns the Length of the quaternion, also known as its Norm. This is the same thing as
// the Len of a Vec4
func (q1 Quat) Len() float32 {
	return float32(math.Sqrt(float64(q1.W*q1.W + q1.V[0]*q1.V[0] + q1.V[1]*q1.V[1] + q1.V[2]*q1.V[2])))
}

// Norm() is an alias for Len() since both are very common terms.
func (q1 Quat) Norm() float32 {
	return q1.Len()
}

// Normalizes the quaternion, returning its versor (unit quaternion).
//
// This is the same as normalizing it as a Vec4.
func (q1 Quat) Normalize() Quat {
	length := q1.Len()

	if FloatEqual(1, length) {
		return q1
	}

	return Quat{q1.W * 1 / length, q1.V.Mul(1 / length)}
}

// The inverse of a quaternion. The inverse is equivalent
// to the conjugate divided by the square of the length.
//
// This method computes the square norm by directly adding the sum
// of the squares of all terms instead of actually squaring q1.Len(),
// both for performance and percision.
func (q1 Quat) Inverse() Quat {
	return q1.Conjugate().Scale(1 / q1.Dot(q1))
}

// Rotates a vector by the rotation this quaternion represents.
// This will result in a 3D vector. Strictly speaking, this is
// equivalent to q1.v.q* where the "."" is quaternion multiplication and v is interpreted
// as a quaternion with W 0 and V v. In code:
// q1.Mul(Quat{0,v}).Mul(q1.Conjugate()), and
// then retrieving the imaginary (vector) part.
//
// In practice, we hand-compute this in the general case and simplify
// to save a few operations.
func (q1 Quat) Rotate(v Vec3) Vec3 {
	cross := q1.V.Cross(v)
	// v + 2q_w * (q_v x v) + 2q_v x (q_v x v)
	return v.Add(cross.Mul(2 * q1.W)).Add(q1.V.Mul(2).Cross(cross))
}

// Returns the homogeneous 3D rotation matrix corresponding to the quaternion.
func (q1 Quat) Mat4() Mat4 {
	w, x, y, z := q1.W, q1.V[0], q1.V[1], q1.V[2]
	return Mat4{1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0, 2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*w*x, 0, 2*x*z + 2*w*y, 2*y*z - 2*w*x, 1 - 2*x*x - 2*y*y, 0, 0, 0, 0, 1}
}

// The dot product between two quaternions, equivalent to if this was a Vec4
func (q1 Quat) Dot(q2 Quat) float32 {
	return q1.W*q1.W + q1.V[0]*q1.V[0] + q1.V[1]*q1.V[1] + q1.V[2]*q1.V[2]
}

// Returns whether the quaternions are approximately equal, as if
// FloatEqual was called on each matching element
func (q1 Quat) ApproxEqual(q2 Quat) bool {
	return FloatEqual(q1.W, q2.W) && q1.V.ApproxEqual(q2.V)
}

// Returns whether the quaternions are approximately equal with a given tolerence, as if
// FloatEqualThreshold was called on each matching element with the given epsilon
func (q1 Quat) ApproxEqualThreshold(q2 Quat, epsilon float32) bool {
	return FloatEqualThreshold(q1.W, q2.W, epsilon) && q1.V.ApproxEqualThreshold(q2.V, epsilon)
}

// Returns whether the quaternions are approximately equal using the given comparison function, as if
// the function had been called on each individual element
func (q1 Quat) ApproxEqualFunc(q2 Quat, f func(float32, float32) bool) bool {
	return f(q1.W, q2.W) && q1.V.ApproxFuncEqual(q2.V, f)
}

// Slerp is *S*pherical *L*inear Int*erp*olation, a method of interpolating
// between two quaternions. This always takes the straightest path on the sphere between
// the two quaternions, and maintains constant velocity.
//
// However, it's expensive and QuatSlerp(q1,q2) is not the same as QuatSlerp(q2,q1)
func QuatSlerp(q1, q2 Quat, amount float32) Quat {
	q1, q2 = q1.Normalize(), q2.Normalize()
	dot := q1.Dot(q2)

	// This is here for precision errors, I'm perfectly aware the *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it
	dot = Clamp(dot, -1, 1)

	theta := float32(math.Acos(float64(dot))) * amount
	c, s := float32(math.Cos(float64(theta))), float32(math.Sin(float64(theta)))
	rel := q2.Sub(q1.Scale(dot)).Normalize()

	return q2.Sub(q1.Scale(c)).Add(rel.Scale(s))
}

// *L*inear Int*erp*olation between two Quaternions, cheap and simple.
//
// Not excessively useful, but uses can be found.
func QuatLerp(q1, q2 Quat, amount float32) Quat {
	return q1.Add(q2.Sub(q1).Scale(amount))
}

// *Normalized* *L*inear Int*erp*olation between two Quaternions. Cheaper than Slerp
// and usually just as good. This is literally Lerp with Normalize() called on it.
//
// Unlike Slerp, constant velocity isn't maintained, but it's much faster and
// Nlerp(q1,q2) and Nlerp(q2,q1) return the same path. You should probably
// use this more often unless you're suffering from choppiness due to the
// non-constant velocity problem.
func QuatNlerp(q1, q2 Quat, amount float32) Quat {
	return QuatLerp(q1, q2, amount).Normalize()
}

// Performs a rotation in the specified order. If the order is not
// a valid RotationOrder, this function will panic
//
// The rotation "order" is more of an axis descriptor. For instance XZX would
// tell the function to interpret angle1 as a rotation about the X axis, angle2 about
// the Z axis, and angle3 about the X axis again.
//
// Based off the code for the Matlab function "angle2quat", though this implementation
// only supports 3 single angles as opposed to multiple angles.
func AnglesToQuat(angle1, angle2, angle3 float32, order RotationOrder) Quat {
	s := [3]float64{}
	c := [3]float64{}

	s[0], c[0] = math.Sincos(float64(angle1 / 2))
	s[1], c[1] = math.Sincos(float64(angle2 / 2))
	s[2], c[2] = math.Sincos(float64(angle3 / 2))

	ret := Quat{}
	switch order {
	case ZYX:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
			float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
		}
	case ZYZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*c[1]*c[2] + c[0]*c[1]*s[2]),
		}
	case ZXY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
		}
	case ZXZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
		}
	case YXZ:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
			float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
		}
	case YXY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*c[1]*c[2] + c[0]*c[1]*s[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
		}
	case YZX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
		}
	case YZY:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		}
	case XYZ:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*s[1]*s[2])
		ret.V = Vec3{float32(c[0]*s[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] - s[0]*c[1]*s[2]),
			float32(c[0]*c[1]*s[2] + s[0]*s[1]*c[2]),
		}
	case XYX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
			float32(s[0]*s[1]*c[2] - c[0]*s[1]*s[2]),
		}
	case XZY:
		ret.W = float32(c[0]*c[1]*c[2] + s[0]*s[1]*s[2])
		ret.V = Vec3{float32(s[0]*c[1]*c[2] - c[0]*s[1]*s[2]),
			float32(c[0]*c[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*c[1]*s[2]),
		}
	case XZX:
		ret.W = float32(c[0]*c[1]*c[2] - s[0]*c[1]*s[2])
		ret.V = Vec3{float32(c[0]*c[1]*s[2] + s[0]*c[1]*c[2]),
			float32(c[0]*s[1]*s[2] - s[0]*s[1]*c[2]),
			float32(c[0]*s[1]*c[2] + s[0]*s[1]*s[2]),
		}
	default:
		panic("Unsupported rotation order")
	}
	return ret
}
