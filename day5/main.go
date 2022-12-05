package main

import (
	"container/list"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

type stack struct {
	*list.List
}

func (s *stack) Push(v interface{}) {
	s.PushBack(v)
}

func (s *stack) PushN(vs ...interface{}) {
	for _, v := range vs {
		s.Push(v)
	}
}

func (s *stack) PushReverseN(vs ...interface{}) {
	for i := len(vs) - 1; i >= 0; i-- {
		s.Push(vs[i])
	}
}

func (s *stack) Pop() interface{} {
	e := s.Back()
	if e != nil {
		s.Remove(e)
		return e.Value
	}
	return nil
}

func (s *stack) PopN(n int) []interface{} {
	var res []interface{}
	for i := 0; i < n; i++ {
		res = append(res, s.Pop())
	}
	return res
}

func (s *stack) Peek() interface{} {
	e := s.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (s *stack) String() string {
	var res []string
	for e := s.Front(); e != nil; e = e.Next() {
		res = append(res, fmt.Sprintf("%v", e.Value))
	}
	return strings.Join(res, ",")
}

func mvN(from *stack, to *stack, n int) {
	for i := 0; i < n; i++ {
		to.Push(from.Pop())
	}
}

func mvNReverse(from *stack, to *stack, n int) {
	data := from.PopN(n)
	to.PushReverseN(data...)
}

type process struct {
	n              int
	fromStackIndex int
	toStackIndex   int
}

func main() {
	var stacks [9]*stack
	stacks[0] = &stack{list.New()}
	stacks[0].PushReverseN("V", "J", "B", "D")
	stacks[1] = &stack{list.New()}
	stacks[1].PushReverseN("F", "D", "R", "W", "B", "V", "P")
	stacks[2] = &stack{list.New()}
	stacks[2].PushReverseN("Q", "W", "C", "D", "L", "F", "G", "R")
	stacks[3] = &stack{list.New()}
	stacks[3].PushReverseN("B", "D", "N", "L", "M", "P", "J", "W")
	stacks[4] = &stack{list.New()}
	stacks[4].PushReverseN("Q", "S", "C", "P", "B", "N", "H")
	stacks[5] = &stack{list.New()}
	stacks[5].PushReverseN("G", "N", "S", "B", "D", "R")
	stacks[6] = &stack{list.New()}
	stacks[6].PushReverseN("H", "S", "F", "Q", "M", "P", "B", "Z")
	stacks[7] = &stack{list.New()}
	stacks[7].PushReverseN("F", "L", "W")
	stacks[8] = &stack{list.New()}
	stacks[8].PushReverseN("R", "M", "F", "V", "S")

	// var stacks [3]*stack
	// stacks[0] = &stack{list.New()}
	// stacks[0].PushReverseN("N", "Z")
	// stacks[1] = &stack{list.New()}
	// stacks[1].PushReverseN("D", "C", "M")
	// stacks[2] = &stack{list.New()}
	// stacks[2].PushReverseN("P")

	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", func(block []string) error {
		lines := block[1]
		for _, line := range strings.Split(lines, "\n") {
			p := process{}
			if n, err := fmt.Sscanf(line, "move %d from %d to %d", &p.n, &p.fromStackIndex, &p.toStackIndex); err == nil && n > 0 {
				mvN(stacks[p.fromStackIndex-1], stacks[p.toStackIndex-1], p.n)
			}
		}
		return nil
	})
	fmt.Fprint(os.Stdout, "P1:")
	for _, s := range stacks {
		fmt.Fprint(os.Stdout, s.Peek())
	}
	println()

	stacks[0] = &stack{list.New()}
	stacks[0].PushReverseN("V", "J", "B", "D")
	stacks[1] = &stack{list.New()}
	stacks[1].PushReverseN("F", "D", "R", "W", "B", "V", "P")
	stacks[2] = &stack{list.New()}
	stacks[2].PushReverseN("Q", "W", "C", "D", "L", "F", "G", "R")
	stacks[3] = &stack{list.New()}
	stacks[3].PushReverseN("B", "D", "N", "L", "M", "P", "J", "W")
	stacks[4] = &stack{list.New()}
	stacks[4].PushReverseN("Q", "S", "C", "P", "B", "N", "H")
	stacks[5] = &stack{list.New()}
	stacks[5].PushReverseN("G", "N", "S", "B", "D", "R")
	stacks[6] = &stack{list.New()}
	stacks[6].PushReverseN("H", "S", "F", "Q", "M", "P", "B", "Z")
	stacks[7] = &stack{list.New()}
	stacks[7].PushReverseN("F", "L", "W")
	stacks[8] = &stack{list.New()}
	stacks[8].PushReverseN("R", "M", "F", "V", "S")

	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", func(block []string) error {
		lines := block[1]
		for _, line := range strings.Split(lines, "\n") {
			p := process{}
			if n, err := fmt.Sscanf(line, "move %d from %d to %d", &p.n, &p.fromStackIndex, &p.toStackIndex); err == nil && n > 0 {
				mvNReverse(stacks[p.fromStackIndex-1], stacks[p.toStackIndex-1], p.n)
			}
		}
		return nil
	})
	fmt.Fprint(os.Stdout, "P2:")
	for _, s := range stacks {
		fmt.Fprint(os.Stdout, s.Peek())
	}
	println()
}
