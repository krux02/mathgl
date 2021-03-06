// Copyright 2014 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
// "math"
)

type Mat2 [4]float32
type Mat2x3 [6]float32
type Mat2x4 [8]float32
type Mat3x2 [6]float32
type Mat3 [9]float32
type Mat3x4 [12]float32
type Mat4x2 [8]float32
type Mat4x3 [12]float32
type Mat4 [16]float32

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident2() Mat2 {
	return Mat2{1, 0, 0, 1}
}

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident3() Mat3 {
	return Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident4() Mat4 {
	return Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// Diag creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag2(v Vec2) Mat2 {
	m := Mat2{}
	m[0], m[3] = v[0], v[1]
	return m
}

// Diag creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag3(v Vec3) Mat3 {
	m := Mat3{}
	m[0], m[4], m[8] = v[0], v[1], v[2]
	return m
}

// Diag creates a diagonal matrix from the entries of the input vector.
// That is, for each pointer for row==col, vector[row] is the entry. Otherwise it's 0.
//
// Another way to think about it is that the identity is this function where the every vector element is 1.
func Diag4(v Vec4) Mat4 {
	m := Mat4{}
	m[0], m[5], m[10], m[15] = v[0], v[1], v[2], v[3]
	return m
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat2FromRows(row0, row1 Vec2) Mat2 {
	return Mat2{row0[0], row1[0], row0[1], row1[1]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat2x3FromRows(row0, row1 Vec3) Mat2x3 {
	return Mat2x3{row0[0], row1[0], row0[1], row1[1], row0[2], row1[2]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat2x4FromRows(row0, row1 Vec4) Mat2x4 {
	return Mat2x4{row0[0], row1[0], row0[1], row1[1], row0[2], row1[2], row0[3], row1[3]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat3x2FromRows(row0, row1, row2 Vec2) Mat3x2 {
	return Mat3x2{row0[0], row1[0], row2[0], row0[1], row1[1], row2[1]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat3FromRows(row0, row1, row2 Vec3) Mat3 {
	return Mat3{row0[0], row1[0], row2[0], row0[1], row1[1], row2[1], row0[2], row1[2], row2[2]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat3x4FromRows(row0, row1, row2 Vec4) Mat3x4 {
	return Mat3x4{row0[0], row1[0], row2[0], row0[1], row1[1], row2[1], row0[2], row1[2], row2[2], row0[3], row1[3], row2[3]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat4x2FromRows(row0, row1, row2, row3 Vec2) Mat4x2 {
	return Mat4x2{row0[0], row1[0], row2[0], row3[0], row0[1], row1[1], row2[1], row3[1]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat4x3FromRows(row0, row1, row2, row3 Vec3) Mat4x3 {
	return Mat4x3{row0[0], row1[0], row2[0], row3[0], row0[1], row1[1], row2[1], row3[1], row0[2], row1[2], row2[2], row3[2]}
}

// Mat<Size>FromRows builds a new matrix from row vectors.
// The resulting matrix will still be in column major order, but this can be
// good for hand-building matrices.

func Mat4FromRows(row0, row1, row2, row3 Vec4) Mat4 {
	return Mat4{row0[0], row1[0], row2[0], row3[0], row0[1], row1[1], row2[1], row3[1], row0[2], row1[2], row2[2], row3[2], row0[3], row1[3], row2[3], row3[3]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat2FromCols(col0, col1 Vec2) Mat2 {
	return Mat2{col0[0], col0[1], col1[0], col1[1]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat2x3FromCols(col0, col1, col2 Vec2) Mat2x3 {
	return Mat2x3{col0[0], col0[1], col1[0], col1[1], col2[0], col2[1]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat2x4FromCols(col0, col1, col2, col3 Vec2) Mat2x4 {
	return Mat2x4{col0[0], col0[1], col1[0], col1[1], col2[0], col2[1], col3[0], col3[1]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat3x2FromCols(col0, col1 Vec3) Mat3x2 {
	return Mat3x2{col0[0], col0[1], col0[2], col1[0], col1[1], col1[2]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat3FromCols(col0, col1, col2 Vec3) Mat3 {
	return Mat3{col0[0], col0[1], col0[2], col1[0], col1[1], col1[2], col2[0], col2[1], col2[2]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat3x4FromCols(col0, col1, col2, col3 Vec3) Mat3x4 {
	return Mat3x4{col0[0], col0[1], col0[2], col1[0], col1[1], col1[2], col2[0], col2[1], col2[2], col3[0], col3[1], col3[2]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat4x2FromCols(col0, col1 Vec4) Mat4x2 {
	return Mat4x2{col0[0], col0[1], col0[2], col0[3], col1[0], col1[1], col1[2], col1[3]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat4x3FromCols(col0, col1, col2 Vec4) Mat4x3 {
	return Mat4x3{col0[0], col0[1], col0[2], col0[3], col1[0], col1[1], col1[2], col1[3], col2[0], col2[1], col2[2], col2[3]}
}

// Mat<Size>FromCols builds a new matrix from column vectors.
func Mat4FromCols(col0, col1, col2, col3 Vec4) Mat4 {
	return Mat4{col0[0], col0[1], col0[2], col0[3], col1[0], col1[1], col1[2], col1[3], col2[0], col2[1], col2[2], col2[3], col3[0], col3[1], col3[2], col3[3]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2) Add(m2 Mat2) Mat2 {
	return Mat2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2x3) Add(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2x4) Add(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3x2) Add(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3) Add(m2 Mat3) Mat3 {
	return Mat3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3x4) Add(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4x2) Add(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4x3) Add(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4) Add(m2 Mat4) Mat4 {
	return Mat4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2) Sub(m2 Mat2) Mat2 {
	return Mat2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2x3) Sub(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2x4) Sub(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3x2) Sub(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3) Sub(m2 Mat3) Mat3 {
	return Mat3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3x4) Sub(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4x2) Sub(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4x3) Sub(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4) Sub(m2 Mat4) Mat4 {
	return Mat4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2) Mul(c float32) Mat2 {
	return Mat2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2x3) Mul(c float32) Mat2x3 {
	return Mat2x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2x4) Mul(c float32) Mat2x4 {
	return Mat2x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3x2) Mul(c float32) Mat3x2 {
	return Mat3x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3) Mul(c float32) Mat3 {
	return Mat3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3x4) Mul(c float32) Mat3x4 {
	return Mat3x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4x2) Mul(c float32) Mat4x2 {
	return Mat4x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4x3) Mul(c float32) Mat4x3 {
	return Mat4x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4) Mul(c float32) Mat4 {
	return Mat4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x1(m2 Vec2) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2(m2 Mat2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x3(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x4(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x1(m2 Vec3) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x2(m2 Mat3x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3(m2 Mat3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x4(m2 Mat3x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x1(m2 Vec4) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x2(m2 Mat4x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x3(m2 Mat4x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4(m2 Mat4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x1(m2 Vec2) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2(m2 Mat2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x3(m2 Mat2x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x4(m2 Mat2x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x1(m2 Vec3) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x2(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3(m2 Mat3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x4(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x1(m2 Vec4) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x2(m2 Mat4x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x3(m2 Mat4x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4(m2 Mat4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x1(m2 Vec2) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2(m2 Mat2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x3(m2 Mat2x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x4(m2 Mat2x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x1(m2 Vec3) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x2(m2 Mat3x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3(m2 Mat3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x4(m2 Mat3x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x1(m2 Vec4) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x2(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x3(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4(m2 Mat4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2) Transpose() Mat2 {
	return Mat2{m1[0], m1[2], m1[1], m1[3]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2x3) Transpose() Mat3x2 {
	return Mat3x2{m1[0], m1[2], m1[4], m1[1], m1[3], m1[5]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2x4) Transpose() Mat4x2 {
	return Mat4x2{m1[0], m1[2], m1[4], m1[6], m1[1], m1[3], m1[5], m1[7]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3x2) Transpose() Mat2x3 {
	return Mat2x3{m1[0], m1[3], m1[1], m1[4], m1[2], m1[5]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3) Transpose() Mat3 {
	return Mat3{m1[0], m1[3], m1[6], m1[1], m1[4], m1[7], m1[2], m1[5], m1[8]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3x4) Transpose() Mat4x3 {
	return Mat4x3{m1[0], m1[3], m1[6], m1[9], m1[1], m1[4], m1[7], m1[10], m1[2], m1[5], m1[8], m1[11]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4x2) Transpose() Mat2x4 {
	return Mat2x4{m1[0], m1[4], m1[1], m1[5], m1[2], m1[6], m1[3], m1[7]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4x3) Transpose() Mat3x4 {
	return Mat3x4{m1[0], m1[4], m1[8], m1[1], m1[5], m1[9], m1[2], m1[6], m1[10], m1[3], m1[7], m1[11]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4) Transpose() Mat4 {
	return Mat4{m1[0], m1[4], m1[8], m1[12], m1[1], m1[5], m1[9], m1[13], m1[2], m1[6], m1[10], m1[14], m1[3], m1[7], m1[11], m1[15]}
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat2) Det() float32 {
	return m[0]*m[2] - m[1]*m[3]
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat3) Det() float32 {
	return m[0]*m[4]*m[8] + m[3]*m[7]*m[2] + m[6]*m[1]*m[5] - m[6]*m[4]*m[2] - m[3]*m[1]*m[8] - m[0]*m[7]*m[5]
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat4) Det() float32 {
	return m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13] - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12] + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12] - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
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
func (m Mat2) Inv() Mat2 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat2{}
	}
	retMat := Mat2{m[3], -m[1], -m[2], m[0]}
	return retMat.Mul(1 / det)
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
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
func (m Mat3) Inv() Mat3 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat3{}
	}
	retMat := Mat3{m[4]*m[8] - m[5]*m[7], m[2]*m[7] - m[1]*m[8], m[1]*m[5] - m[2]*m[4], m[5]*m[6] - m[3]*m[8], m[0]*m[8] - m[2]*m[6], m[2]*m[3] - m[0]*m[5], m[3]*m[7] - m[4]*m[6], m[1]*m[6] - m[0]*m[7], m[0]*m[4] - m[1]*m[3]}
	return retMat.Mul(1 / det)
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
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
func (m Mat4) Inv() Mat4 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat4{}
	}
	retMat := Mat4{-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15], m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15], -m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15], m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11], m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15], -m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15], m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15], -m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11], -m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15], m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15], -m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15], m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11], m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14], -m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14], m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14], -m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10]}
	return retMat.Mul(1 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2) ApproxEqual(m2 Mat2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2x3) ApproxEqual(m2 Mat2x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2x4) ApproxEqual(m2 Mat2x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3x2) ApproxEqual(m2 Mat3x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3) ApproxEqual(m2 Mat3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3x4) ApproxEqual(m2 Mat3x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4x2) ApproxEqual(m2 Mat4x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4x3) ApproxEqual(m2 Mat4x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4) ApproxEqual(m2 Mat4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2) ApproxEqualThreshold(m2 Mat2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2x3) ApproxEqualThreshold(m2 Mat2x3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2x4) ApproxEqualThreshold(m2 Mat2x4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3x2) ApproxEqualThreshold(m2 Mat3x2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3) ApproxEqualThreshold(m2 Mat3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3x4) ApproxEqualThreshold(m2 Mat3x4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4x2) ApproxEqualThreshold(m2 Mat4x2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4x3) ApproxEqualThreshold(m2 Mat4x3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4) ApproxEqualThreshold(m2 Mat4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2) ApproxFuncEqual(m2 Mat2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2x3) ApproxFuncEqual(m2 Mat2x3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2x4) ApproxFuncEqual(m2 Mat2x4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3x2) ApproxFuncEqual(m2 Mat3x2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3) ApproxFuncEqual(m2 Mat3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3x4) ApproxFuncEqual(m2 Mat3x4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4x2) ApproxFuncEqual(m2 Mat4x2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4x3) ApproxFuncEqual(m2 Mat4x3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4) ApproxFuncEqual(m2 Mat4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat2) At(row, col int) float32 {
	return m[col*2+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat2x3) At(row, col int) float32 {
	return m[col*2+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat2x4) At(row, col int) float32 {
	return m[col*2+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat3x2) At(row, col int) float32 {
	return m[col*3+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat3) At(row, col int) float32 {
	return m[col*3+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat3x4) At(row, col int) float32 {
	return m[col*3+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat4x2) At(row, col int) float32 {
	return m[col*4+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat4x3) At(row, col int) float32 {
	return m[col*4+row]
}

// At returns the matrix element at the given row and column.
// This is equivalent to mat[col * numRow + row] where numRow is constant
// (E.G. for a Mat3x2 it's equal to 3)
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// At(5,0) will work just like At(1,1). Or it may panic if it's out of bounds.
func (m Mat4) At(row, col int) float32 {
	return m[col*4+row]
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat2) Set(row, col int, value float32) {
	m[col*2+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat2x3) Set(row, col int, value float32) {
	m[col*2+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat2x4) Set(row, col int, value float32) {
	m[col*2+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat3x2) Set(row, col int, value float32) {
	m[col*3+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat3) Set(row, col int, value float32) {
	m[col*3+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat3x4) Set(row, col int, value float32) {
	m[col*3+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat4x2) Set(row, col int, value float32) {
	m[col*4+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat4x3) Set(row, col int, value float32) {
	m[col*4+row] = value
}

// Set sets the corresponding matrix element at the given row and column.
// This has a pointer receiver because it mutates the matrix.
//
// This method is garbage-in garbage-out. For instance, on a Mat4 asking for
// Set(5,0,val) will work just like Set(1,1,val). Or it may panic if it's out of bounds.
func (m *Mat4) Set(row, col int, value float32) {
	m[col*4+row] = value
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat2) Index(row, col int) int {
	return col*2 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat2x3) Index(row, col int) int {
	return col*2 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat2x4) Index(row, col int) int {
	return col*2 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat3x2) Index(row, col int) int {
	return col*3 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat3) Index(row, col int) int {
	return col*3 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat3x4) Index(row, col int) int {
	return col*3 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat4x2) Index(row, col int) int {
	return col*4 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat4x3) Index(row, col int) int {
	return col*4 + row
}

// Index returns the index of the given row and column, to be used with direct
// access. E.G. Index(0,0) = 0.
//
// This is a garbage-in garbage-out method. For instance, on a Mat4 asking for the index of
// (5,0) will work the same as asking for (1,1). Or it may give you a value that will cause
// a panic if you try to access the array with it if it's truly out of bounds.
func (m Mat4) Index(row, col int) int {
	return col*4 + row
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat2) Row(row int) Vec2 {
	return Vec2{m[row+0], m[row+2]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat2x3) Row(row int) Vec3 {
	return Vec3{m[row+0], m[row+2], m[row+4]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat2x4) Row(row int) Vec4 {
	return Vec4{m[row+0], m[row+2], m[row+4], m[row+6]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat3x2) Row(row int) Vec2 {
	return Vec2{m[row+0], m[row+3]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat3) Row(row int) Vec3 {
	return Vec3{m[row+0], m[row+3], m[row+6]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat3x4) Row(row int) Vec4 {
	return Vec4{m[row+0], m[row+3], m[row+6], m[row+9]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat4x2) Row(row int) Vec2 {
	return Vec2{m[row+0], m[row+4]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat4x3) Row(row int) Vec3 {
	return Vec3{m[row+0], m[row+4], m[row+8]}
}

// Row returns a vector representing the corresponding row (starting at row 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecM for a MxN matrix.
func (m Mat4) Row(row int) Vec4 {
	return Vec4{m[row+0], m[row+4], m[row+8], m[row+12]}
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat2) Rows() (row0, row1 Vec2) {
	return m.Row(0), m.Row(1)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat2x3) Rows() (row0, row1 Vec3) {
	return m.Row(0), m.Row(1)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat2x4) Rows() (row0, row1 Vec4) {
	return m.Row(0), m.Row(1)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat3x2) Rows() (row0, row1, row2 Vec2) {
	return m.Row(0), m.Row(1), m.Row(2)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat3) Rows() (row0, row1, row2 Vec3) {
	return m.Row(0), m.Row(1), m.Row(2)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat3x4) Rows() (row0, row1, row2 Vec4) {
	return m.Row(0), m.Row(1), m.Row(2)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat4x2) Rows() (row0, row1, row2, row3 Vec2) {
	return m.Row(0), m.Row(1), m.Row(2), m.Row(3)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat4x3) Rows() (row0, row1, row2, row3 Vec3) {
	return m.Row(0), m.Row(1), m.Row(2), m.Row(3)
}

// Rows decomposes a matrix into its corresponding row vectors.
// This is equivalent to calling mat.Row for each row.
func (m Mat4) Rows() (row0, row1, row2, row3 Vec4) {
	return m.Row(0), m.Row(1), m.Row(2), m.Row(3)
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat2) Col(col int) Vec2 {
	return Vec2{m[col*2+0], m[col*2+1]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat2x3) Col(col int) Vec2 {
	return Vec2{m[col*2+0], m[col*2+1]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat2x4) Col(col int) Vec2 {
	return Vec2{m[col*2+0], m[col*2+1]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat3x2) Col(col int) Vec3 {
	return Vec3{m[col*3+0], m[col*3+1], m[col*3+2]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat3) Col(col int) Vec3 {
	return Vec3{m[col*3+0], m[col*3+1], m[col*3+2]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat3x4) Col(col int) Vec3 {
	return Vec3{m[col*3+0], m[col*3+1], m[col*3+2]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat4x2) Col(col int) Vec4 {
	return Vec4{m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat4x3) Col(col int) Vec4 {
	return Vec4{m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3]}
}

// Col returns a vector representing the corresponding column (starting at col 0).
// This package makes no distinction between row and column vectors, so it
// will be a normal VecN for a MxN matrix.
func (m Mat4) Col(col int) Vec4 {
	return Vec4{m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3]}
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat2) Cols() (col0, col1 Vec2) {
	return m.Col(0), m.Col(1)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat2x3) Cols() (col0, col1, col2 Vec2) {
	return m.Col(0), m.Col(1), m.Col(2)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat2x4) Cols() (col0, col1, col2, col3 Vec2) {
	return m.Col(0), m.Col(1), m.Col(2), m.Col(3)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat3x2) Cols() (col0, col1 Vec3) {
	return m.Col(0), m.Col(1)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat3) Cols() (col0, col1, col2 Vec3) {
	return m.Col(0), m.Col(1), m.Col(2)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat3x4) Cols() (col0, col1, col2, col3 Vec3) {
	return m.Col(0), m.Col(1), m.Col(2), m.Col(3)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat4x2) Cols() (col0, col1 Vec4) {
	return m.Col(0), m.Col(1)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat4x3) Cols() (col0, col1, col2 Vec4) {
	return m.Col(0), m.Col(1), m.Col(2)
}

// Cols decomposes a matrix into its corresponding column vectors.
// This is equivalent to calling mat.Col for each column.
func (m Mat4) Cols() (col0, col1, col2, col3 Vec4) {
	return m.Col(0), m.Col(1), m.Col(2), m.Col(3)
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat2) Trace() float32 {
	return m[0] + m[3]
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat3) Trace() float32 {
	return m[0] + m[4] + m[8]
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat4) Trace() float32 {
	return m[0] + m[5] + m[10] + m[15]
}
