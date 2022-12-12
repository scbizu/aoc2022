package main

import (
	"context"
	"fmt"
	"math"

	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/matrix"
)

var m = matrix.NewMatrix[rune](41, 70)

var (
	start = grid.Vec{X: 0, Y: 0}
	end   = grid.Vec{X: 0, Y: 0}
)

var dist = make(map[grid.Vec]int)

func filterNeighbor(f grid.Vec, vs []grid.Vec) []grid.Vec {
	var ns []grid.Vec
	for _, v := range vs {
		vv := m.Get(v.X, v.Y)
		if vv < m.Get(f.X, f.Y)-1 {
			continue
		}
		ns = append(ns, v)
	}
	return ns
}

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	m.Print()
	dist[end] = 0
	q := []grid.Vec{end}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		ns := m.GetNeighbor(v.X, v.Y)
		ns = filterNeighbor(v, ns)
		for _, n := range ns {
			vec := grid.Vec{
				X: n.X,
				Y: n.Y,
			}
			if _, ok := dist[vec]; !ok {
				dist[vec] = dist[v] + 1
				q = append(q, vec)
			}
		}
	}

	fmt.Printf("p1: %d\n", dist[start])

	min := math.MaxInt64

	m.ForEach(func(x, y int, v rune) {
		if v != 'a' {
			return
		}
		if d, ok := dist[grid.Vec{X: x, Y: y}]; ok && d < min {
			min = d
		}
	})
	fmt.Printf("p2: %d\n", min)
}

var rowCount int

func handler(line string) error {
	if len(line) == 0 {
		return nil
	}
	for i, c := range line {
		if c == 'S' {
			start.X = rowCount
			start.Y = i
			c = 'a'
		}
		if c == 'E' {
			end.X = rowCount
			end.Y = i
			c = 'z'
		}
		m.Add(rowCount, i, c)
	}
	rowCount++
	return nil
}
