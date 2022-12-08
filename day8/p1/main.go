package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/matrix"
)

var count int

func main() {
	input.NewTXTFile("input.txt").ReadByLineEx(context.Background(), handler)
	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			if i == 0 || i == m.Rows()-1 {
				count++
				continue
			}
			if j == 0 || j == m.Cols()-1 {
				count++
				continue
			}
			if isTop(m, m.Get(i, j), i, j) {
				count++
			}
		}
	}
	fmt.Fprintf(os.Stdout, "p1: result: %d\n", count)
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

func isTop(m matrix.Matrix[int], v, row, col int) bool {
	for i := row + 1; i < m.Rows(); i++ {
		if m.Get(i, col) >= v {
			for i := row - 1; i >= 0; i-- {
				if m.Get(i, col) >= v {
					for i := col + 1; i < m.Cols(); i++ {
						if m.Get(row, i) >= v {
							for i := col - 1; i >= 0; i-- {
								if m.Get(row, i) >= v {
									return false
								}
							}
						}
					}
				}
			}
		}
	}
	return true
}
