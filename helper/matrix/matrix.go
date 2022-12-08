package matrix

import "fmt"

type Matrix[T comparable] [][]T

func NewMatrix[T comparable](rows, cols int) Matrix[T] {
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

func (m Matrix[T]) Print() {
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			fmt.Printf("%v ", m.Get(i, j))
		}
		fmt.Println()
	}
}
