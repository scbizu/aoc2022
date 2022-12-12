package priority

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type Ele[T constraints.Ordered] struct {
	data []T
}

type PriorityQueue struct {
	heap.Interface
}

func (e *Ele[T]) Push(x any) {
	e.data = append(e.data, x.(T))
}

func (e *Ele[T]) Pop() any {
	n := len(e.data)
	x := e.data[n-1]
	e.data = e.data[:n-1]
	return x
}

// PriorityQueue[T] implements sort.Interface for generic usage
func (e Ele[T]) Len() int { return len(e.data) }

func (e Ele[T]) Less(i, j int) bool {
	return e.data[i] < e.data[j]
}

func (e Ele[T]) Swap(i, j int) {
	e.data[i], e.data[j] = e.data[j], e.data[i]
}
