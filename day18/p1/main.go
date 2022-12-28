package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

var cubes = set.New[grid.XYZVec]()

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	var totalSurface int
	cubes.Each(func(cube grid.XYZVec) bool {
		totalSurface += exposeSurface(cube)
		return true
	})
	fmt.Fprintf(os.Stdout, "p1: res: %d\n", totalSurface)
}

func exposeSurface(cube grid.XYZVec) int {
	var surface int

	for _, c := range cube.Neighbors6() {
		if !cubes.Has(c) {
			surface++
		}
	}

	return surface
}

func handler(line string) error {
	if line == "" {
		return nil
	}
	var cube grid.XYZVec
	if _, err := fmt.Sscanf(line, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z); err != nil {
		return err
	}
	cubes.Add(cube)
	return nil
}
