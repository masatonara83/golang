package main

import "fmt"



func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

// func Area(v Vertex) int {
// 	return v.X * v.Y
// }

type Vertex struct {
	x, y int
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4, 5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
