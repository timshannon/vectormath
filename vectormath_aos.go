// Copyright (c) 2006, 2007 Sony Computer Entertainment Inc.
// Copyright (c) 2012 James Helferty
// All rights reserved.

package vectormath

type Vector3 struct {
	X, Y, Z, _ float32
}

type Vector4 struct {
	X, Y, Z, W float32
}

type Point3 struct {
	X, Y, Z, _ float32
}

type Quat struct {
	X, Y, Z, W float32
}

type Matrix3 struct {
	Col0, Col1, Col2 Vector3
}

type Matrix4 struct {
	Col0, Col1, Col2, Col3 Vector4
}

type Transform3 struct {
	Col0, Col1, Col2, Col3 Vector3
}

func SliceToM4(m *Matrix4, slice []float32) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			//column major
			m.SetElem(c, r, slice[(c*4)+r])
		}
	}
}

func SliceToV3(v3 *Vector3, slice []float32) {
	v3.SetElem(0, slice[0])
	v3.SetElem(1, slice[1])
	v3.SetElem(2, slice[2])
}

func M4ToSlice(slice []float32, m *Matrix4) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			//column major
			slice[(c*4)+r] = m.GetElem(c, r)
		}
	}
}

func V3ToSlice(slice []float32, v3 *Vector3) {
	slice[0] = v3.GetElem(0)
	slice[1] = v3.GetElem(1)
	slice[2] = v3.GetElem(2)
}
