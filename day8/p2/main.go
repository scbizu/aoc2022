package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/matrix"
)

var max int

func main() {
	input.NewTXTFile("input.txt").ReadByLineEx(context.Background(), handler)
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			t := trees(m, m.Get(i, j), i, j)
			if t > max {
				max = t
			}
		}
	}
	fmt.Fprintf(os.Stdout, "p2: result: %d\n", max)
}

var row, col = 99, 99

var m = matrix.NewMatrix[int](row, col)

func handler(row int, line string) error {
	for col, c := range line {
		v, _ := strconv.Atoi(string(c))
		m.Add(row, col, v)
	}
	return nil
}

func trees(m matrix.Matrix[int], v, row, col int) int {
	var l, r, t, b int
	if row > 0 {
		for i := row - 1; i >= 0; i-- {
			if m.Get(i, col) >= v {
				t++
				break
			}
			if m.Get(i, col) <= v {
				t++
			}
		}
	}
	if row < m.Rows()-1 {
		for i := row + 1; i < m.Rows(); i++ {
			if m.Get(i, col) >= v {
				b++
				break
			}
			if m.Get(i, col) <= v {
				b++
			}
		}
	}
	if col > 0 {
		for i := col - 1; i >= 0; i-- {
			if m.Get(row, i) >= v {
				l++
				break
			}
			if m.Get(row, i) <= v {
				l++
			}
		}
	}
	if col < m.Cols()-1 {
		for i := col + 1; i < m.Cols(); i++ {
			if m.Get(row, i) >= v {
				r++
				break
			}
			if m.Get(row, i) <= v {
				r++
			}
		}
	}
	return l * r * t * b
}
