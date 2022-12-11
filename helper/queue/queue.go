package queue

import (
	"container/list"
	"fmt"
	"strings"
)

type Queue[T comparable] struct {
	*list.List
}

func (q *Queue[T]) Push(v T) {
	q.PushBack(v)
}

func (q *Queue[T]) PushN(vs ...T) {
	for _, v := range vs {
		q.Push(v)
	}
}

func (q *Queue[T]) Peek() T {
	var res T
	e := q.Front()
	if e != nil {
		return e.Value.(T)
	}
	return res
}

func (q *Queue[T]) Pop() T {
	var res T
	e := q.Front()
	if e != nil {
		q.Remove(e)
		return e.Value.(T)
	}
	return res
}

func (q *Queue[T]) String() string {
	var res []string
	for e := q.Front(); e != nil; e = e.Next() {
		res = append(res, fmt.Sprintf("%v", e.Value))
	}
	return strings.Join(res, ",")
}
