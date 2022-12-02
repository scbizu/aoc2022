package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
)

var lineMap = map[string]int{
	"A X": 1 + 3,
	"B X": 1 + 0,
	"C X": 1 + 6,

	"A Y": 2 + 6,
	"B Y": 2 + 3,
	"C Y": 2 + 0,

	"A Z": 3 + 0,
	"B Z": 3 + 6,
	"C Z": 3 + 3,
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
