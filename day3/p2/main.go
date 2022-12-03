package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
)

var sum int

func main() {
	input.NewTXTFile("input.txt").ReadByBlockEx(context.Background(), separator, handler)
	fmt.Fprintln(os.Stdout, sum)
}

func handler(lines []string) error {
	if len(lines) != 3 {
		return fmt.Errorf("invalid block size: %d", len(lines))
	}
	ascMap := make(map[uint8]struct{})
	for _, asc := range lines[0] {
		ascMap[uint8(rune(asc))] = struct{}{}
	}
	for _, asc := range lines[1] {
		if _, ok := ascMap[uint8(rune(asc))]; ok {
			for _, asc2 := range lines[2] {
				if uint8(rune(asc)) == uint8(rune(asc2)) {
					c := rune(asc)
					if rune('a') <= c && c <= rune('z') {
						sum += int(c - 'a' + 1)
					}
					if rune('A') <= c && c <= rune('Z') {
						sum += int(c - 'A' + 27)
					}
					goto next
				}
			}
		}
	}
next:
	return nil
}

func separator(i int, _ string) bool {
	if i == 0 {
		return false
	}
	return i%3 == 0
}
