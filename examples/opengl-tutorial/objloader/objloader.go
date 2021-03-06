// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objloader

import (
	"bufio"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"os"
)

type MeshObject struct {
	Vertices []mgl32.Vec3
	UVs      []mgl32.Vec2
	Normals  []mgl32.Vec3
}

func LoadObject(fname string, invertV bool) *MeshObject {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	vertices, uvs, normals := make([]mgl32.Vec3, 0), make([]mgl32.Vec2, 0), make([]mgl32.Vec3, 0)
	vIndices, uvIndices, nIndices := make([]uint, 0), make([]uint, 0), make([]uint, 0)

	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		lineReader := bufio.NewReader(StringReader(line))
		header, _ := lineReader.ReadString(' ')
		restOfLine, _ := lineReader.ReadString('\n')

		switch header[:len(header)-1] {
		case "v":
			vert := mgl32.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f\n", &vert[0], &vert[1], &vert[2])
			if count != 3 {
				panic("Wrong vert count")
			}
			vertices = append(vertices, vert)

		case "vt":
			uv := mgl32.Vec2{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f\n", &uv[0], &uv[1])
			if count != 2 {
				panic("Wrong uv count")
			}
			if invertV {
				// For DDS textures
				uv = mgl32.Vec2{uv[0], 1 - uv[1]}
			}
			uvs = append(uvs, uv)
		case "vn":
			norm := mgl32.Vec3{}
			count, _ := fmt.Sscanf(restOfLine, "%f %f %f\n", &norm[0], &norm[1], &norm[2])
			if count != 3 {
				panic("Wrong norm count")
			}
			normals = append(normals, norm)
		case "f":
			//vert1, vert2, vert3 string
			vIndex, uvIndex, nIndex := [3]uint{}, [3]uint{}, [3]uint{}
			matches, _ := fmt.Sscanf(restOfLine, "%d/%d/%d %d/%d/%d %d/%d/%d\n", &vIndex[0], &uvIndex[0], &nIndex[0], &vIndex[1], &uvIndex[1], &nIndex[1], &vIndex[2], &uvIndex[2], &nIndex[2])
			if matches != 9 {
				panic("Can't read file")
			}
			vIndices = append(vIndices, vIndex[:]...)
			uvIndices = append(uvIndices, uvIndex[:]...)
			nIndices = append(nIndices, nIndex[:]...)
		default:
			// eat line
		}
	}

	//fmt.Println(vertices)

	obj := &MeshObject{make([]mgl32.Vec3, 0, len(vIndices)), make([]mgl32.Vec2, 0, len(uvIndices)), make([]mgl32.Vec3, 0, len(nIndices))}
	for i := range vIndices {
		vIndex, uvIndex, nIndex := vIndices[i], uvIndices[i], nIndices[i]

		vert, uv, norm := vertices[vIndex-1], uvs[uvIndex-1], normals[nIndex-1]

		obj.Vertices = append(obj.Vertices, vert)
		obj.UVs = append(obj.UVs, uv)
		obj.Normals = append(obj.Normals, norm)
	}

	return obj
}

type StringReader string

func (s StringReader) Read(byt []byte) (n int, err error) {
	copy(byt, string(s))
	return len(byt), nil
}
