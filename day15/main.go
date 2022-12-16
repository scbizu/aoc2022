package main

import (
	"context"
	"fmt"
	"os"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
)

var lineExpr = "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"

var ls = []*grid.HorizontalLine{}

var (
	minX, maxX = 0, 4000000
	minY, maxY = 0, 4000000
)

var y = 2000000

var (
	count int
	b     = set.New[grid.Vec]()
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	ls = grid.Merge(ls)
	var count int
	for _, l := range ls {
		count += l.Len()
	}
	fmt.Fprintf(os.Stdout, "count: %d\n", count-b.Size())

	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler2)

	println("p2:")

	for i, l := range lines {
		l = grid.Merge(l)
		if len(l) > 1 {
			for _, l2 := range l {
				l2.Print()
			}
			fmt.Fprintf(os.Stdout, "x: %d\n", (l[0].ToX()+1)*4000000+i)
			break
		}
	}
}

func handler(line string) error {
	var sensor grid.Vec
	var beacon grid.Vec
	if _, err := fmt.Sscanf(line, lineExpr, &sensor.X, &sensor.Y, &beacon.X, &beacon.Y); err != nil {
		return err
	}
	distance := grid.Distance(sensor, beacon)

	if sensor.Y+distance < y || sensor.Y-distance > y {
		return nil
	}

	var x0, x1 int
	if sensor.Y <= y {
		x0 = sensor.X - (distance - y + sensor.Y)
		x1 = sensor.X + (distance - y + sensor.Y)
	} else {
		x0 = sensor.X - (distance + y - sensor.Y)
		x1 = sensor.X + (distance + y - sensor.Y)
	}
	var temp int
	if x0 < x1 {
		temp = x0
		x0 = x1
		x1 = temp
	}

	ls = append(ls, grid.NewHorizontalLine(x1, x0, y))

	if beacon.Y == y && !b.Has(beacon) {
		b.Add(beacon)
	}

	return nil
}

var lines = make(map[int][]*grid.HorizontalLine)

func handler2(line string) error {
	var sensor grid.Vec
	var beacon grid.Vec
	if _, err := fmt.Sscanf(line, lineExpr, &sensor.X, &sensor.Y, &beacon.X, &beacon.Y); err != nil {
		return err
	}

	distance := grid.Distance(sensor, beacon)
	y1, y2 := sensor.Y-distance, sensor.Y+distance

	if y1 < minY {
		y1 = minY
	}
	if y2 > maxY {
		y2 = maxY
	}

	for y := y1; y <= y2; y++ {
		var x0, x1 int
		if sensor.Y <= y {
			x0 = sensor.X - (distance - y + sensor.Y)
			x1 = sensor.X + (distance - y + sensor.Y)
		} else {
			x0 = sensor.X - (distance + y - sensor.Y)
			x1 = sensor.X + (distance + y - sensor.Y)
		}
		var temp int
		if x0 < x1 {
			temp = x0
			x0 = x1
			x1 = temp
		}

		if x0 > maxX {
			x0 = maxX
		}
		if x1 < minX {
			x1 = minX
		}
		lines[y] = append(lines[y], grid.NewHorizontalLine(x1, x0, y))
	}

	return nil
}
