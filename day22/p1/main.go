package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

type loc struct {
	grid.Vec
	face byte
}

// loc implements fmt.Stringer
func (l loc) String() string {
	return fmt.Sprintf("(%d,%d,%c)", l.X, l.Y, l.face)
}

func main() {
	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", handler)
	height := len(board)
	current := loc{}
	t, s := parsePath(path)
	for i := 0; i < len(t); i++ {
		current.face = turn(current.face, t[i])
		remStep := s[i]
		vec := current.Add(direction(current.face, 1))
		for remStep > 0 {
			fmt.Printf("vec: %v,rem: %d\n", vec, remStep)
			if vec.Y > height-1 {
				vec.Y = 0
			}
			if vec.Y < 0 {
				vec.Y = height - 1
			}
			if vec.X > width-1 {
				vec.X = 0
			}
			if vec.X < 0 {
				vec.X = width - 1
			}
			if v, ok := board[grid.Vec{X: vec.X, Y: vec.Y}]; ok && v == '#' {
				break
			}
			if v, ok := board[grid.Vec{X: vec.X, Y: vec.Y}]; ok && v == '.' {
				current.Vec = vec
				remStep--
				fmt.Printf("current: %s,rem: %d\n", current, remStep)
			}
			vec = vec.Add(direction(current.face, 1))
		}
	}
	fmt.Fprintf(os.Stdout, "current: %s\n", current)
	fmt.Fprintf(os.Stdout, "password: %d\n", 1000*(current.Y+1)+4*(current.X+1)+faceNum(current.face))
}

func faceNum(face byte) int {
	switch face {
	case 'R':
		return 0
	case 'L':
		return 2
	case 'U':
		return 3
	case 'D':
		return 1
	}
	panic("invalid face")
}

var (
	block int
	board map[grid.Vec]byte
	path  string
)

func parsePath(path string) ([]byte, []int) {
	di := []byte{'R'}
	var steps []int

	for len(path) > 0 {
		var step int
		for i, c := range path {
			if c >= '0' && c <= '9' {
				step = step*10 + int(c-'0')
				if i == len(path)-1 {
					path = path[i+1:]
					break
				}
			} else {
				path = path[i:]
				break
			}
		}
		steps = append(steps, step)
		if len(path) > 0 {
			di = append(di, path[0])
			path = path[1:]
		}
	}

	if len(di) != len(steps) {
		fmt.Printf("di: %v, steps: %v\n", di, steps)
		panic("invalid path")
	}

	return di, steps
}

func direction(face byte, step int) grid.Vec {
	switch face {
	case 'R':
		return grid.Vec{X: step}
	case 'L':
		return grid.Vec{X: -step}
	case 'U':
		return grid.Vec{Y: -step}
	case 'D':
		return grid.Vec{Y: step}
	}
	panic("invalid face")
}

func turn(raw byte, to byte) byte {
	switch raw {
	case 'R':
		if to == 'L' {
			return 'U'
		}
		return 'D'
	case 'L':
		if to == 'L' {
			return 'D'
		}
		return 'U'
	case 'U':
		if to == 'L' {
			return 'L'
		}
		return 'R'
	case 'D':
		if to == 'L' {
			return 'R'
		}
		return 'L'
	}
	return to
}

var width int

func handler(lines []string) error {
	if len(lines) == 0 {
		return nil
	}
	if len(lines) != 2 {
		panic("invalid input")
	}
	bl := strings.Split(lines[0], "\n")
	board = make(map[grid.Vec]byte)
	for i, line := range bl {
		if len(line) > width {
			width = len(line)
		}
		for j, c := range line {
			board[grid.Vec{X: j, Y: i}] = byte(c)
		}
	}
	path = strings.TrimSuffix(lines[1], "\n")
	return nil
}
