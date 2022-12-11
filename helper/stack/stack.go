package stack

import (
	"container/list"
	"fmt"
	"strings"
)

type Stack[T comparable] struct {
	*list.List
}

func (s *Stack[T]) Push(v T) {
	s.PushBack(v)
}

func (s *Stack[T]) PushN(vs ...T) {
	for _, v := range vs {
		s.Push(v)
	}
}

func (s *Stack[T]) PushReverseN(vs ...T) {
	for i := len(vs) - 1; i >= 0; i-- {
		s.Push(vs[i])
	}
}

func (s *Stack[T]) Pop() T {
	var res T
	e := s.Back()
	if e != nil {
		s.Remove(e)
		return e.Value.(T)
	}
	return res
}

func (s *Stack[T]) PopN(n int) []interface{} {
	var res []interface{}
	for i := 0; i < n; i++ {
		res = append(res, s.Pop())
	}
	return res
}

func (s *Stack[T]) Peek() interface{} {
	e := s.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *Stack[T]) String() string {
	var res []string
	for e := s.Front(); e != nil; e = e.Next() {
		res = append(res, fmt.Sprintf("%v", e.Value))
	}
	return strings.Join(res, ",")
}
