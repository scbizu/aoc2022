package main

import (
	"container/list"
	"context"
	"fmt"
	"os"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/stack"
)

var (
	cubes     map[grid.XYZVec]struct{} = make(map[grid.XYZVec]struct{})
	space     map[grid.XYZVec]struct{} = make(map[grid.XYZVec]struct{})
	min, max                           = -1, 50
	cubeStack                          = &stack.Stack[grid.XYZVec]{List: list.New()}
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	// to determine the air space, perform BFS(stack) to iterate all the bounding cubes
	c := grid.XYZVec{X: min, Y: min, Z: min}
	cubeStack.Push(c)
	for cubeStack.Len() > 0 {
		cube := cubeStack.Pop()
		if cube.X < min || cube.X > max || cube.Y < min || cube.Y > max || cube.Z < min || cube.Z > max {
			continue
		}
		if _, ok := cubes[cube]; ok {
			continue
		}
		space[cube] = struct{}{}
		for _, n := range cube.Neighbors6() {
			if _, ok := space[n]; ok {
				continue
			}
			cubeStack.Push(n)
		}
	}

	lavaCubes := set.New[grid.XYZVec]()
	// and the remaining is the lava space , the internal lava space cant be reached by the outside air
	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			for k := min; k <= max; k++ {
				cube := grid.XYZVec{X: i, Y: j, Z: k}
				if _, ok := space[cube]; !ok && !lavaCubes.Has(cube) {
					lavaCubes.Add(cube)
				}
			}
		}
	}
	// the p1 solution
	var surface int
	lavaCubes.Each(func(cube grid.XYZVec) bool {
		surface += exposeSurface(cube, lavaCubes)
		return true
	})
	fmt.Fprintf(os.Stdout, "p2: res: %d\n", surface)
}

func handler(line string) error {
	if line == "" {
		return nil
	}
	var cube grid.XYZVec
	if _, err := fmt.Sscanf(line, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z); err != nil {
		return err
	}
	cubes[cube] = struct{}{}
	return nil
}

func exposeSurface(cube grid.XYZVec, cubeSet *set.Set[grid.XYZVec]) int {
	var surface int

	for _, c := range cube.Neighbors6() {
		if !cubeSet.Has(c) {
			surface++
		}
	}

	return surface
}
