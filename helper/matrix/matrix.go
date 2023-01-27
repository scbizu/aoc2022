package matrix

import (
	"fmt"

	"github.com/scbizu/aoc2022/helper/grid"

	"golang.org/x/exp/constraints"
)

type Matrix[T constraints.Ordered] [][]T

type Point[T constraints.Ordered] struct {
	X     int
	Y     int
	Value T
}

type Points[T constraints.Ordered] []Point[T]

func (p Points[T]) Len() int {
	return len(p)
}

func (p Points[T]) Less(i, j int) bool {
	if p[i].Value < p[j].Value {
		return true
	}
	if p[i].Value == p[j].Value {
		if p[i].X < p[j].X {
			return true
		}
		if p[i].X == p[j].X {
			return p[i].Y < p[j].Y
		}
	}
	return false
}

func (p Points[T]) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Points[T]) Push(x any) {
	*p = append(*p, x.(Point[T]))
	var s string
	for _, v := range *p {
		s += fmt.Sprintf("%v ", v)
	}
}

func (p *Points[T]) Pop() any {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p *Points[T]) String() string {
	var s string
	for _, v := range *p {
		s += fmt.Sprintf("%v ", v)
	}
	return s
}

func NewMatrix[T constraints.Ordered](rows, cols int) Matrix[T] {
	m := make(Matrix[T], rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]T, cols)
	}
	return m
}

func (m Matrix[T]) Add(x, y int, v T) {
	m[x][y] = v
}

func (m Matrix[T]) Get(x, y int) T {
	return m[x][y]
}

func (m Matrix[T]) Rows() int {
	return len(m)
}

func (m Matrix[T]) Cols() int {
	return len(m[0])
}

func (m Matrix[T]) ForEach(f func(x, y int, v T)) {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			f(i, j, m.Get(i, j))
		}
	}
}

func (m Matrix[T]) GetNeighbor(x, y int) []grid.Vec {
	var ns []grid.Vec
	if x > 0 {
		ns = append(ns, grid.Vec{X: x - 1, Y: y})
	}
	if x < m.Rows()-1 {
		ns = append(ns, grid.Vec{X: x + 1, Y: y})
	}
	if y > 0 {
		ns = append(ns, grid.Vec{X: x, Y: y - 1})
	}
	if y < m.Cols()-1 {
		ns = append(ns, grid.Vec{X: x, Y: y + 1})
	}
	return ns
}

func (m Matrix[T]) Reset(c T) {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			// set the default value of matrix
			m[i][j] = c
		}
	}
}

func (m Matrix[T]) GetNeighbors8(x, y int) []grid.Vec {
	var ns []grid.Vec
	if x > 0 {
		ns = append(ns, grid.Vec{X: x - 1, Y: y})
	}
	if x < m.Rows()-1 {
		ns = append(ns, grid.Vec{X: x + 1, Y: y})
	}
	if y > 0 {
		ns = append(ns, grid.Vec{X: x, Y: y - 1})
	}
	if y < m.Cols()-1 {
		ns = append(ns, grid.Vec{X: x, Y: y + 1})
	}
	if x > 0 && y > 0 {
		ns = append(ns, grid.Vec{X: x - 1, Y: y - 1})
	}
	if x < m.Rows()-1 && y < m.Cols()-1 {
		ns = append(ns, grid.Vec{X: x + 1, Y: y + 1})
	}
	if x > 0 && y < m.Cols()-1 {
		ns = append(ns, grid.Vec{X: x - 1, Y: y + 1})
	}
	if x < m.Rows()-1 && y > 0 {
		ns = append(ns, grid.Vec{X: x + 1, Y: y - 1})
	}
	return ns
}

func (m Matrix[T]) Print() {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			fmt.Printf("%v ", m.Get(i, j))
		}
		fmt.Println()
	}
}
