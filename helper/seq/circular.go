package seq

import (
	"fmt"
)

type Circular[T comparable] struct {
	seq []T
}

func NewCircular[T comparable](seq []T) *Circular[T] {
	return &Circular[T]{seq: seq}
}

func (c *Circular[T]) Append(v T) {
	c.seq = append(c.seq, v)
}

func (c *Circular[T]) AtIndex(index int) T {
	return c.seq[c.Index(index)]
}

func (c *Circular[T]) Index(index int) int {
	if index < 0 {
		return len(c.seq) + index%len(c.seq)
	}
	return index % len(c.seq)
}

func (c *Circular[T]) List() []T {
	return c.seq
}

func (c *Circular[T]) BaseLen() int {
	return len(c.seq)
}

func (c *Circular[T]) Find(v T) int {
	for i, e := range c.seq {
		if e == v {
			return i
		}
	}
	return -1
}

func (c *Circular[T]) Print() {
	for _, e := range c.seq {
		fmt.Printf("%v ", e)
	}
	fmt.Println()
}

func (c *Circular[T]) Move(index int, step int) *Circular[T] {
	if index < 0 || index > len(c.seq) {
		panic("index out of range")
	}
	if step == 0 {
		return c
	}
	ic := c.seq[index]
	c.seq = append(c.seq[:index], c.seq[index+1:]...)
	pos := c.Index(index + step)

	var newSeq []T
	for i := 0; i < len(c.seq); i++ {
		if i == pos {
			newSeq = append(newSeq, ic)
		}
		newSeq = append(newSeq, c.seq[i])
	}
	c.seq = newSeq
	return c
}
