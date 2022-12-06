package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
)

var lastIndex int

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	fmt.Fprintf(os.Stdout, "last index: %d\n", lastIndex)

	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handlerP2)
	fmt.Fprintf(os.Stdout, "P2 last index: %d\n", lastIndex)
}

func handler(line string) error {
	for i := 0; i < len(line)-3; i++ {
		window := make(map[rune]struct{})
		for _, c := range line[i : i+4] {
			window[c] = struct{}{}
		}
		if len(window) == 4 {
			lastIndex = i + 4
			break
		}
	}
	return nil
}

func handlerP2(line string) error {
	for i := 0; i < len(line)-13; i++ {
		window := make(map[rune]struct{})
		for _, c := range line[i : i+14] {
			window[c] = struct{}{}
		}
		if len(window) == 14 {
			lastIndex = i + 14
			break
		}
	}
	return nil
}
