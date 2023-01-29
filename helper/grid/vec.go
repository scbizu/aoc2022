package grid

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
)

type VecMatrix[T any] map[Vec]T

func NewVecMatrix[T any]() VecMatrix[T] {
	return make(VecMatrix[T])
}

func (m VecMatrix[T]) Add(v Vec, data T) {
	m[v] = data
}

func (m VecMatrix[T]) Get(v Vec) (T, bool) {
	data, ok := m[v]
	return data, ok
}

func (m VecMatrix[T]) Rows() int {
	var maxY, minY int
	for v := range m {
		maxY = v.Y
		minY = v.Y
		break
	}
	for v := range m {
		if v.Y > maxY {
			maxY = v.Y
		}
		if v.Y < minY {
			minY = v.Y
		}
	}
	return maxY - minY + 1
}

func (m VecMatrix[T]) Cols() int {
	var maxX, minX int
	for v := range m {
		maxX = v.X
		minX = v.X
		break
	}
	for v := range m {
		if v.X > maxX {
			maxX = v.X
		}
		if v.X < minX {
			minX = v.X
		}
	}
	return maxX - minX + 1
}

func (m VecMatrix[T]) GetNeighbors8(v Vec) []Vec {
	return []Vec{
		v.Add(Vec{X: -1, Y: -1}),
		v.Add(Vec{X: -1, Y: 0}),
		v.Add(Vec{X: -1, Y: 1}),
		v.Add(Vec{X: 0, Y: -1}),
		v.Add(Vec{X: 0, Y: 1}),
		v.Add(Vec{X: 1, Y: -1}),
		v.Add(Vec{X: 1, Y: 0}),
		v.Add(Vec{X: 1, Y: 1}),
	}
}

func (m VecMatrix[T]) GetNeighbor(v Vec) []Vec {
	return []Vec{
		v.Add(Vec{X: -1, Y: 0}),
		v.Add(Vec{X: 0, Y: -1}),
		v.Add(Vec{X: 0, Y: 1}),
		v.Add(Vec{X: 1, Y: 0}),
	}
}

func (m VecMatrix[T]) ForEach(f func(Vec, T)) {
	for v, data := range m {
		f(v, data)
	}
}

func (m VecMatrix[T]) String(format string) string {
	buffer := bytes.NewBuffer(nil)
	m.Print(buffer, format)
	return buffer.String()
}

func (m VecMatrix[T]) Print(w io.Writer, format string) {
	byline := make(map[int][]Vec)
	for v := range m {
		byline[v.Y] = append(byline[v.Y], v)
	}
	for _, v := range byline {
		sort.Slice(v, func(i, j int) bool {
			return v[i].X < v[j].X
		})
	}
	var lines [][]Vec
	for line := range byline {
		lines = append(lines, byline[line])
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0].Y < lines[j][0].Y
	})

	for _, line := range lines {
		for _, v := range line {
			fmt.Fprintf(w, format, m[v])
		}
		fmt.Fprintln(w)
	}
}

func (m VecMatrix[T]) Reset(c T) {
	for v := range m {
		m[v] = c
	}
}

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
