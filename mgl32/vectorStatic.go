package mgl32

func (v Vec2) Vec3(z float32) Vec3 {
	return Vec3{v[0], v[1], z}
}

func (v Vec2) Vec4(z, w float32) Vec4 {
	return Vec4{v[0], v[1], z, w}
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) Vec4(w float32) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
}

func (v Vec4) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) Vec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// extracts the elements of the vector for direct value assignment
func (v Vec2) Elem() (x, y float32) {
	return v[0], v[1]
}

// extracts the elements of the vector for direct value assignment
func (v Vec3) Elem() (x, y, z float32) {
	return v[0], v[1], v[2]
}

// extracts the elements of the vector for direct value assignment
func (v Vec4) Elem() (x, y, z, w float32) {
	return v[0], v[1], v[2], v[3]
}
