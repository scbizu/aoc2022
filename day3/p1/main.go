package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
)

var sum int

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	fmt.Fprintln(os.Stdout, sum)
}

func handler(line string) error {
	lineLen := len(line)
	parts := make(map[int]struct{})
	for _, c := range line[:lineLen/2] {
		parts[int(c)] = struct{}{}
	}
	for _, c := range line[lineLen/2:] {
		if _, ok := parts[int(c)]; ok {
			if rune('a') <= c && c <= rune('z') {
				sum += int(c - 'a' + 1)
				break
			}
			if rune('A') <= c && c <= rune('Z') {
				sum += int(c - 'A' + 27)
				break
			}
		}
	}
	return nil
}
