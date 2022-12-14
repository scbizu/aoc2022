package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	v := grid.Vec{}
	for v.Y < maxY {
		v = dropFrom(grid.Vec{X: 500, Y: 0}, maxY)
	}
	var count int
	for _, v := range path {
		if v == kindSand {
			count++
		}
	}
	fmt.Fprintf(os.Stdout, "p1: count: %d\n", count)

	path = make(map[grid.Vec]kind)
	maxY = 0
	v = grid.Vec{}
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	for {
		if v.X == 500 && v.Y == 0 {
			break
		}
		v = dropFrom(grid.Vec{X: 500, Y: 0}, maxY+2)
	}
	count = 0
	for _, v := range path {
		if v == kindSand {
			count++
		}
	}

	fmt.Fprintf(os.Stdout, "p2: count: %d\n", count)
}

type kind int

const (
	kindAir kind = iota
	kindRock
	kindSand
)

var (
	path = make(map[grid.Vec]kind)
	maxY int
)

func handler(line string) error {
	if len(line) == 0 {
		return nil
	}
	parts := strings.Split(line, "->")
	from := strings.Split(parts[0], ",")
	fromX, _ := strconv.Atoi(strings.TrimSpace(from[0]))
	fromY, _ := strconv.Atoi(strings.TrimSpace(from[1]))
	if fromY > maxY {
		maxY = fromY
	}
	fromVec := grid.Vec{X: fromX, Y: fromY}
	for _, v := range parts[1:] {
		to := strings.Split(v, ",")
		toX, _ := strconv.Atoi(strings.TrimSpace(to[0]))
		toY, _ := strconv.Atoi(strings.TrimSpace(to[1]))
		if toY > maxY {
			maxY = toY
		}
		toVec := grid.Vec{X: toX, Y: toY}
		l := grid.NewLine(fromVec, toVec)
		l.OnDraw(func(v grid.Vec) error {
			path[grid.Vec{X: v.X, Y: v.Y}] = kindRock
			return nil
		})
		fromVec = toVec
	}
	return nil
}

func dropFrom(v grid.Vec, breakAt int) grid.Vec {
	if v.Y == breakAt {
		path[grid.Vec{X: v.X, Y: v.Y}] = kindRock
		return v
	}
	_, ok1 := path[grid.Vec{X: v.X, Y: v.Y + 1}]
	if !ok1 {
		return dropFrom(grid.Vec{X: v.X, Y: v.Y + 1}, breakAt)
	}
	_, ok2 := path[grid.Vec{X: v.X - 1, Y: v.Y + 1}]
	if !ok2 {
		return dropFrom(grid.Vec{X: v.X - 1, Y: v.Y + 1}, breakAt)
	}
	_, ok3 := path[grid.Vec{X: v.X + 1, Y: v.Y + 1}]
	if !ok3 {
		return dropFrom(grid.Vec{X: v.X + 1, Y: v.Y + 1}, breakAt)
	}
	if ok1 && ok2 && ok3 {
		path[grid.Vec{X: v.X, Y: v.Y}] = kindSand
		return v
	}
	panic("unreachable")
}
