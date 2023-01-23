package grid

import (
	"fmt"
	"math"
)

type Vec struct {
	X int
	Y int
}

func Abs(v1, v2 int) int {
	return int(math.Abs(float64(v1 - v2)))
}

func Distance(v1, v2 Vec) int {
	return Abs(v1.X, v2.X) + Abs(v1.Y, v2.Y)
}

func (v Vec) Add(v2 Vec) Vec {
	return Vec{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec) Sub(v2 Vec) Vec {
	return Vec{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec) Print() {
	println(v.X, v.Y)
}

func (v Vec) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}
