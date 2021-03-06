// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Making vectorf.go")
	vecs := GenVec()
	//fmt.Println(vecs)
	vecf, err := os.Create("../mgl32/vector.go")
	if err != nil {
		panic(err)
	}
	defer vecf.Close()

	_, err = vecf.Write([]byte(vecs))
	if err != nil {
		panic(err)
	}

	//fmt.Println("Making matrixf.go")
	mats := GenMat()
	//fmt.Println(mats)
	matf, err := os.Create("../mgl32/matrix.go")
	if err != nil {
		panic(err)
	}
	defer matf.Close()

	_, err = matf.Write([]byte(mats))
	if err != nil {
		panic(err)
	}
	//fmt.Println(mats)
	//fmt.Println("Done")
}

func GenVec() (s string) {
	vecs := `// Copyright 2014 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import(
	"math"
)

`

	for m := 2; m <= 4; m++ {
		vecs += GenVecDef(m)
	}

	vecs += "\n"
	for m := 2; m <= 4; m++ {
		vecs += GenVecAdd(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecSub(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecMul(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecDot(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecLen(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecNormalize(m)
	}

	vecs += GenVecCross()

	for m := 2; m <= 4; m++ {
		vecs += GenVecEq(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecThresholdEq(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecFuncEq(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecElement(m)
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			vecs += GenVecOuterProd(m, n)
		}
	}

	/*for m := 2; m <= 4; m++ {
		for o := 2; o <= 4; o++ {
			vecs += GenVecMatMul(m,o)
		}
	}*/

	return vecs
}

func GenVecCross() string {
	header := `// The vector cross product is an operation only defined on 3D vectors. It is equivalent to
// Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}. 
// Another interpretation is that it's the vector whose magnitude is |v1||v2|sin(theta) 
// where theta is the angle between v1 and v2.
//
// The cross product is most often used for finding surface normals. The cross product of vectors
// will generate a vector that is perpendicular to the plane they form.
//
// Technically, a generalized cross product exists as an "(N-1)ary" operation 
// (that is, the 4D cross product requires 3 4D vectors). But the binary
// 3D (and 7D) cross product is the most important. It can be considered 
// the area of a parallelogram with sides v1 and v2.
//
// Like the dot product, the cross product is roughly a measure of directionality. 
// Two normalized perpendicular vectors will return a vector with a magnitude of
// 1.0 or -1.0 and two parallel vectors will return a vector with magnitude 0.0. 
// The cross product is "anticommutative" meaning v1.Cross(v2) = -v2.Cross(v1),
// this property can be useful to know when finding normals, 
// as taking the wrong cross product can lead to the opposite normal of the one you want.
`
	return header + "func (v1 Vec3) Cross(v2 Vec3) Vec3 {\n\treturn Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}\n}\n\n"
}

func GenVecDef(m int) (s string) {
	return fmt.Sprintf("type Vec%d [%d]float32\n", m, m)
}

/*func GenVecMatMul(m, n, o int) (s string) {
	s = fmt.Sprintf("func (v1 %s) Mul%s(v2 %s) %s {\n\treturn %s{", GenMatName(m,n), GenMatName(n,o), GenMatName(m,o), GenMatName(m,o))
	for j := 0; j < o; j++ {
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ {
				s += fmt.Sprintf("m1[%d] * m2[%d]", i+k*m, j*n+k)
				s += " + "
			}
			s = s[:len(s)-3]
			s += ", "
		}
	}
	s = s[:len(s)-2]
	s += "}\n}\n\n"
	return s
}*/

func GenVecAdd(m int) (s string) {
	s = `// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
`
	s += fmt.Sprintf("func (v1 %s) Add(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] + v2[%d]", i, i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecSub(m int) (s string) {
	s = `// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
`
	s += fmt.Sprintf("func (v1 %s) Sub(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] - v2[%d]", i, i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecMul(m int) (s string) {
	s = `// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
`
	s += fmt.Sprintf("func (v1 %s) Mul(c float32) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * c", i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"

	return s
}

func GenVecDot(m int) (s string) {
	s = `// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
`
	s += fmt.Sprintf("func (v1 %s) Dot(v2 %s) float32 {\n\treturn ", VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * v2[%d]", i, i)
		if i != m-1 {
			s += "+"
		}
	}
	s += "\n}\n\n"

	return s
}

func GenVecLen(m int) (s string) {
	s = `// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's 
// math.Hypot(v[0], v[1]).
`
	if m != 2 {
		s += fmt.Sprintf("func (v1 %s) Len() float32 {\n\treturn float32(math.Sqrt(float64(", VecName(m))
		for i := 0; i < m; i++ {
			s += fmt.Sprintf("v1[%d] * v1[%d]", i, i)
			if i != m-1 {
				s += "+"
			}
		}
		s += ")))\n}\n\n"
	} else {
		s += fmt.Sprintf("func (v1 %s) Len() float32 {\n\treturn float32(math.Hypot(float64(v1[0]), float64(v1[1])))\n}\n\n", VecName(m))
	}
	return s
}

func GenVecNormalize(m int) (s string) {
	s = `// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
`
	s += fmt.Sprintf("func (v1 %s) Normalize() %s {\n\tl := 1.0/v1.Len()\n\t return %s{", VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * l", i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecEq(m int) (s string) {
	s = `// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
`
	s += fmt.Sprintf("func (v1 %s) ApproxEqual(v2 %s) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqual(v1[i],v2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecThresholdEq(m int) (s string) {
	s = `// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
`
	s += fmt.Sprintf("func (v1 %s) ApproxEqualThreshold(v2 %s, threshold float32) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqualThreshold(v1[i],v2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecFuncEq(m int) (s string) {
	s = `// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
`
	s += fmt.Sprintf("func (v1 %s) ApproxFuncEqual(v2 %s, eq func(float32,float32) bool) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !eq(v1[i],v2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecElement(m int) (s string) {
	elName := func(el int) string {
		switch el {
		case 0:
			return "X"
		case 1:
			return "Y"
		case 2:
			return "Z"
		case 3:
			return "W"
		default:
			panic("Can't generate that element")
		}
	}

	for i := 0; i < m; i++ {
		s += `// This is an element access func, it is equivalent to v[n] where
// n is some valid index. The mappings are XYZW (X=0, Y=1 etc). Benchmarks
// show that this is more or less as fast as direct acces, probably due to
// inlining, so use v[0] or v.X() depending on personal preference.
`

		s += fmt.Sprintf("func (v %s) %s() float32 {\n\treturn v[%d]\n}\n\n", VecName(m), elName(i), i)
	}

	return
}

func GenVecOuterProd(m, n int) string {
	vecs := `// Does the vector outer product
// of two vectors. The outer product produces an
// NxM matrix. E.G. a Vec2 * Vec3 = Mat2x3.
//
// The outer product can be thought of as the "opposite"
// of the Dot product. The Dot product treats both vectors like matrices
// oriented such that the left one has N columns and the right has N rows.
// So Vec3.Vec3 = Mat1x3*Mat3x1 = Mat1 = Scalar.
//
// The outer product orients it so they're facing "outward": Vec2*Vec3
// = Mat2x1*Mat1x3 = Mat2x3.
`

	vecs += fmt.Sprintf("func (v1 %s) OuterProd%d(v2 %s) %s {\n\treturn %s{", VecName(m), n, VecName(n), GenMatName(m, n), GenMatName(m, n))

	for c := 0; c < n; c++ {
		for r := 0; r < m; r++ {
			if r != 0 || c != 0 {
				vecs += ","
			}

			vecs += fmt.Sprintf("v1[%d]*v2[%d]", r, c)
		}
	}

	vecs += "}\n}\n\n"

	return vecs
}

func VecName(m int) (s string) {
	return fmt.Sprintf("Vec%d", m)
}

func GenMat() string {
	mats := `// Copyright 2014 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import(
	// "math"
)

`

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatDef(m, n)
		}
	}
	mats += "\n"

	for m := 2; m <= 4; m++ {
		mats += GenMatIden(m)
	}

	for m := 2; m <= 4; m++ {
		mats += GenMatDiag(m)
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatFromRows(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatFromCols(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatAdd(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatSub(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenScalarMul(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			for o := 1; o <= 4; o++ {
				mats += GenMatMul(m, n, o)
			}
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenTranspose(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		mats += GenDet(m)
	}

	for m := 2; m <= 4; m++ {
		mats += GenInv(m)
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatEq(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatThresholdEq(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatFuncEq(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatAt(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatSet(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatIndex(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatRow(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatRows(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatCol(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatCols(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		mats += GenMatTrace(m)
	}

	//fmt.Println(mats)
	return mats
}

func GenMatIden(m int) (s string) {
	s = `// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
`
	s += fmt.Sprintf("func Ident%d() %s {\n\treturn %s{", m, GenMatName(m, m), GenMatName(m, m))

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				s += "0"
			} else {
				s += "1"
			}

			if i != m-1 || j != m-1 {
				s += ","
			}
		}
	}

	s += "}\n}\n\n"

	return s
}

func GenMatDef(m, n int) (s string) {
	return fmt.Sprintf("type %s [%d]float32\n", GenMatName(m, n), m*n)
}

func GenMatAdd(m, n int) (s string) {
	s = `// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
`
	s += fmt.Sprintf("func (m1 %s) Add(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] + m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenMatSub(m, n int) (s string) {
	s = `// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
`
	s += fmt.Sprintf("func (m1 %s) Sub(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] - m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenScalarMul(m, n int) (s string) {
	s = `// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
`
	s += fmt.Sprintf("func (m1 %s) Mul(c float32) %s {\n\treturn %s{", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] *c", i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenMatMul(m, n, o int) (s string) {
	s = `// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
`
	s += "func " + "(m1 " + GenMatName(m, n) + ") Mul" + fmt.Sprintf("%d", n)

	if n != o {
		s += fmt.Sprintf("x%d", o)
	}

	s += "(m2 " + GenMatName(n, o) + ") " + GenMatName(m, o) + " {\n\treturn " + GenMatName(m, o) + "{"
	for j := 0; j < o; j++ { // For each element of the output array
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ { // For each element of the vector we're multiplying
				s += fmt.Sprintf("m1[%d] * m2[%d]", i+k*m, j*n+k)
				s += " + "
			}
			s = s[:len(s)-3]
			s += ", "
		}
	}
	s = s[:len(s)-2]
	s += "}\n}\n\n"
	return s
}

func GenTranspose(m, n int) (s string) {
	s = `// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]    
`
	s += fmt.Sprintf("func (m1 %s) Transpose() %s {\n\treturn %s{", GenMatName(m, n), GenMatName(n, m), GenMatName(n, m))

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if r != 0 || c != 0 {
				s += ","
			}
			s += fmt.Sprintf("m1[%d]", c*m+r)
		}
	}
	s += "}\n}\n\n"

	return s
}

/*func GenDet(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %s) Mul() float32 {\n\treturn %s{", GenMatName(m,n), GenMatName(n,o), GenMatName(m,o), GenMatName(m,o))
}*/

func GenMatName(m, n int) string {
	if m == 1 {
		return fmt.Sprintf("Vec%d", n)
	}

	if n == 1 {
		return fmt.Sprintf("Vec%d", m)
	}

	s := fmt.Sprintf("Mat%d", m)

	if m != n {
		s += fmt.Sprintf("x%d", n)
	}

	return s
}

func GenDet(m int) string {
	s := `// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
`
	s += fmt.Sprintf("func (m %s) Det() float32 {\n\treturn ", GenMatName(m, m))

	switch m {
	case 2:
		s += "m[0] * m[2] - m[1] * m[3]"
	case 3:
		s += "m[0]*m[4]*m[8] + m[3] * m[7] * m[2] + m[6] * m[1] * m[5] - m[6] * m[4] * m[2] - m[3] * m[1] * m[8] - m[0] * m[7] * m[5]"
	case 4:
		s += "m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13]" +
			" - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12]" +
			" + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12]" +
			" - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]"
	}
	s += "\n}\n\n"

	return s
}

func GenInv(m int) string {
	s := `// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function does not check the determinant to see if the matrix is invertible. 
// If the determinant is 0.0, the value of all elements will be
// infinite. (See here for why: http://play.golang.org/p/Aaj7SnbqIp ) 
// Therefore, if the program really cares, it should check the determinant first.
// In the future, an alternate function may be written which takes in a pre-computed determinant. 
`
	s += fmt.Sprintf("func (m %s) Inv() %s {\n\t", GenMatName(m, m), GenMatName(m, m))
	s += "det := m.Det()\n\t if FloatEqual(det,float32(0.0)) { \n\t\t return " + GenMatName(m, m) + "{}\n\t}\n\t"
	s += "retMat := " + GenMatName(m, m) + "{"

	switch m {
	case 2:
		s += "m[3], -m[1], -m[2], m[0]"
	case 3:
		s += "m[4] * m[8] -m[5] * m[7] , m[2] * m[7] -m[1] * m[8] ,m[1] * m[5] -m[2] * m[4] ,m[5] * m[6] -m[3] * m[8] ,m[0] * m[8] -m[2] * m[6] ,m[2] * m[3] -m[0] * m[5] ,m[3] * m[7] -m[4] * m[6] ,m[1] * m[6] -m[0] * m[7] ,m[0] * m[4] -m[1] * m[3]"
	case 4:
		s += "-m[7] * m[10] * m[13] +m[6] * m[11] * m[13] +m[7] * m[9] * m[14] -m[5] * m[11] * m[14] -m[6] * m[9] * m[15] +m[5] * m[10] * m[15] ," +
			"m[3] * m[10] * m[13] -m[2] * m[11] * m[13] -m[3] * m[9] * m[14] +m[1] * m[11] * m[14] +m[2] * m[9] * m[15] -m[1] * m[10] * m[15] ," +
			"-m[3] * m[6] * m[13] +m[2] * m[7] * m[13] +m[3] * m[5] * m[14] -m[1] * m[7] * m[14] -m[2] * m[5] * m[15] +m[1] * m[6] * m[15] ," +
			"m[3] * m[6] * m[9] -m[2] * m[7] * m[9] -m[3] * m[5] * m[10] +m[1] * m[7] * m[10] +m[2] * m[5] * m[11] -m[1] * m[6] * m[11] ," +
			"m[7] * m[10] * m[12] -m[6] * m[11] * m[12] -m[7] * m[8] * m[14] +m[4] * m[11] * m[14] +m[6] * m[8] * m[15] -m[4] * m[10] * m[15] ," +
			"-m[3] * m[10] * m[12] +m[2] * m[11] * m[12] +m[3] * m[8] * m[14] -m[0] * m[11] * m[14] -m[2] * m[8] * m[15] +m[0] * m[10] * m[15] ," +
			" m[3] * m[6] * m[12] -m[2] * m[7] * m[12] -m[3] * m[4] * m[14] +m[0] * m[7] * m[14] +m[2] * m[4] * m[15] -m[0] * m[6] * m[15] ," +
			"-m[3] * m[6] * m[8] +m[2] * m[7] * m[8] +m[3] * m[4] * m[10] -m[0] * m[7] * m[10] -m[2] * m[4] * m[11] +m[0] * m[6] * m[11] ," +
			"-m[7] * m[9] * m[12] +m[5] * m[11] * m[12] +m[7] * m[8] * m[13] -m[4] * m[11] * m[13] -m[5] * m[8] * m[15] +m[4] * m[9] * m[15] ," +
			"m[3] * m[9] * m[12] -m[1] * m[11] * m[12] -m[3] * m[8] * m[13] +m[0] * m[11] * m[13] +m[1] * m[8] * m[15] -m[0] * m[9] * m[15] ," +
			"-m[3] * m[5] * m[12] +m[1] * m[7] * m[12] +m[3] * m[4] * m[13] -m[0] * m[7] * m[13] -m[1] * m[4] * m[15] +m[0] * m[5] * m[15] ," +
			"m[3] * m[5] * m[8] -m[1] * m[7] * m[8] -m[3] * m[4] * m[9] +m[0] * m[7] * m[9] +m[1] * m[4] * m[11] -m[0] * m[5] * m[11] ," +
			"m[6] * m[9] * m[12] -m[5] * m[10] * m[12] -m[6] * m[8] * m[13] +m[4] * m[10] * m[13] +m[5] * m[8] * m[14] -m[4] * m[9] * m[14] ," +
			"-m[2] * m[9] * m[12] +m[1] * m[10] * m[12] +m[2] * m[8] * m[13] -m[0] * m[10] * m[13] -m[1] * m[8] * m[14] +m[0] * m[9] * m[14] ," +
			"m[2] * m[5] * m[12] -m[1] * m[6] * m[12] -m[2] * m[4] * m[13] +m[0] * m[6] * m[13] +m[1] * m[4] * m[14] -m[0] * m[5] * m[14] ," +
			"-m[2] * m[5] * m[8] +m[1] * m[6] * m[8] +m[2] * m[4] * m[9] -m[0] * m[6] * m[9] -m[1] * m[4] * m[10] +m[0] * m[5] * m[10]"
	}

	s += "}\n\t return retMat.Mul(1/det)\n}\n\n"

	return s
}

func GenMatEq(m, n int) (s string) {
	s = `// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
`
	s += fmt.Sprintf("func (m1 %s) ApproxEqual(m2 %s) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqual(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatThresholdEq(m, n int) (s string) {
	s = `// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
`
	s += fmt.Sprintf("func (m1 %s) ApproxEqualThreshold(m2 %s, threshold float32) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqualThreshold(m1[i],m2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatFuncEq(m, n int) (s string) {
	s = `// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
`
	s += fmt.Sprintf("func (m1 %s) ApproxFuncEqual(m2 %s, eq func(float32,float32) bool) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !eq(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatAt(m, n int) (s string) {
	s = `// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
`

	s += fmt.Sprintf("func (m %s) At(row,col int) float32 {\n\treturn m[col * %d + row]\n}\n\n", GenMatName(m, n), m)

	return
}

func GenMatSet(m, n int) (s string) {
	s = `// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
`

	s += fmt.Sprintf("func (m *%s) Set(row,col int, value float32) {\n\tm[col * %d + row] = value\n}\n\n", GenMatName(m, n), m)

	return
}

func GenMatIndex(m, n int) (s string) {
	s = `// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
`
	s += fmt.Sprintf("func (m %s) Index(row,col int) int {\n\treturn col * %d + row\n}\n\n", GenMatName(m, n), m)

	return
}

func GenMatRow(m, n int) (s string) {
	s = `// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
`

	s += fmt.Sprintf("func (m %s) Row(row int) %s {\n\treturn %s{", GenMatName(m, n), VecName(n), VecName(n))

	for i := 0; i < n; i++ {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("m[row + %d]", m*i)
	}

	s += "}\n}\n\n"

	return
}

var rowNames = [...]string{"row0", "row1", "row2", "row3"}
var colNames = [...]string{"col0", "col1", "col2", "col3"}

func GenMatRows(m, n int) (s string) {
	s = `// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
`
	args := strings.Join(rowNames[:m], ", ")
	s += fmt.Sprintf("func (m %s) Rows() (%s %s) {\n\treturn ", GenMatName(m, n), args, VecName(n))

	for i := 0; i < m; i++ {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("m.Row(%d)", i)
	}

	s += "\n}\n\n"

	return
}

func GenMatFromRows(m, n int) (s string) {
	s = ` // Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.


`

	args := strings.Join(rowNames[:m], ", ")
	s += fmt.Sprintf("func %sFromRows(%s %s) %s {\n\treturn %s{", GenMatName(m, n), args, VecName(n), GenMatName(m, n), GenMatName(m, n))

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j+i != 0 {
				s += ","
			}

			s += fmt.Sprintf("%s[%d]", rowNames[j], i)
		}
	}

	s += "}\n}\n\n"

	return
}

func GenMatCol(m, n int) (s string) {
	s = `// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
`

	s += fmt.Sprintf("func (m %s) Col(col int) %s {\n\treturn %s{", GenMatName(m, n), VecName(m), VecName(m))

	for i := 0; i < m; i++ {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("m[col * %d + %d]", m, i)
	}

	s += "}\n}\n\n"

	return
}

func GenMatCols(m, n int) (s string) {
	s = `// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
`
	args := strings.Join(colNames[:n], ", ")
	s += fmt.Sprintf("func (m %s) Cols() (%s %s) {\n\treturn ", GenMatName(m, n), args, VecName(m))

	for i := 0; i < n; i++ {
		if i != 0 {
			s += ","
		}
		s += fmt.Sprintf("m.Col(%d)", i)
	}

	s += "\n}\n\n"

	return
}

func GenMatFromCols(m, n int) (s string) {
	s = ` // Mat<Size>FromCols builds a new matrix from column vectors.
`
	args := strings.Join(colNames[:n], ", ")
	s += fmt.Sprintf("func %sFromCols(%s %s) %s {\n\treturn %s{", GenMatName(m, n), args, VecName(m), GenMatName(m, n), GenMatName(m, n))

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j+i != 0 {
				s += ","
			}

			s += fmt.Sprintf("%s[%d]", colNames[i], j)
		}
	}

	s += "}\n}\n\n"
	return
}

func GenMatTrace(m int) (s string) {
	s = `// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
`

	s += fmt.Sprintf("func (m %s) Trace() float32{\n\treturn ", GenMatName(m, m))

	for i := 0; i < m; i++ {
		if i != 0 {
			s += "+"
		}

		s += fmt.Sprintf("m[%d]", i*m+i)
	}

	s += "\n}\n\n"

	return s
}

func GenMatDiag(m int) (s string) {
	s = `// Diag creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
`

	s += fmt.Sprintf("func Diag%d(v %s) %s {\n\tm := %s{}\n\t", m, VecName(m), GenMatName(m, m), GenMatName(m, m))

	for i := 0; i < m; i++ {
		if i != 0 {
			s += ","
		}

		s += fmt.Sprintf("m[%d]", i*m+i)
	}

	s += "="

	for i := 0; i < m; i++ {
		if i != 0 {
			s += ","
		}

		s += fmt.Sprintf("v[%d]", i)
	}

	s += "\n\treturn m\n}\n\n"

	return s
}
