package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

var (
	vecMap = map[string]grid.Vec{
		"R": {X: 1, Y: 0},
		"L": {X: -1, Y: 0},
		"U": {X: 0, Y: 1},
		"D": {X: 0, Y: -1},
	}
	s = set.New(grid.Vec{
		X: 0,
		Y: 0,
	})
	s2 = set.New(grid.Vec{
		X: 0,
		Y: 0,
	})
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	fmt.Fprintf(os.Stdout, "p1: result: %d\n", s.Size())
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler2)
	fmt.Fprintf(os.Stdout, "p2: result: %d\n", s2.Size())
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var ht = []grid.Vec{
	{X: 0, Y: 0},
	{X: 0, Y: 0},
}

func handler(line string) error {
	if line == "" {
		return nil
	}
	parts := strings.Split(line, " ")
	step, _ := strconv.Atoi(parts[1])
	for i := 0; i < step; i++ {
		ht[0] = ht[0].Add(vecMap[parts[0]])
		ht = compare(ht)
		if !s.Has(ht[1]) {
			s.Add(ht[1])
		}
	}
	return nil
}

var nodes = make([]grid.Vec, 10)

func handler2(line string) error {
	if line == "" {
		return nil
	}
	parts := strings.Split(line, " ")
	step, _ := strconv.Atoi(parts[1])
	for i := 0; i < step; i++ {
		nodes[0] = nodes[0].Add(vecMap[parts[0]])
		nodes = compare(nodes)
		if !s2.Has(nodes[9]) {
			s2.Add(nodes[9])
		}
	}
	return nil
}

func compare(nodes []grid.Vec) []grid.Vec {
	for i := 1; i < len(nodes); i++ {
		xStep := nodes[i-1].X - nodes[i].X
		yStep := nodes[i-1].Y - nodes[i].Y
		if xStep == 0 && abs(yStep) > 1 {
			// same row
			nodes[i] = nodes[i].Add(grid.Vec{X: 0, Y: yStep / abs(yStep)})
			continue
		}
		if yStep == 0 && abs(xStep) > 1 {
			// same col
			nodes[i] = nodes[i].Add(grid.Vec{X: xStep / abs(xStep), Y: 0})
			continue
		}
		if max(abs(xStep), abs(yStep)) > 1 {
			// diagonal
			nodes[i] = nodes[i].Add(grid.Vec{X: xStep / abs(xStep), Y: yStep / abs(yStep)})
			continue
		}
	}
	return nodes
}
