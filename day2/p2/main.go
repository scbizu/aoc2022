package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
)

var lineMap = map[string]int{
	"A X": 0 + 3,
	"B X": 0 + 1,
	"C X": 0 + 2,

	"A Y": 3 + 1,
	"B Y": 3 + 2,
	"C Y": 3 + 3,

	"A Z": 6 + 2,
	"B Z": 6 + 3,
	"C Z": 6 + 1,
}

var total int

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), countScoreHandler)
	fmt.Fprintln(os.Stdout, total)
}

func countScoreHandler(line string) error {
	if _, ok := lineMap[line]; ok {
		total += lineMap[line]
	}
	return nil
}
