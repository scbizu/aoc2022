package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

var (
	register []int = []int{1}
	sum      int
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	for i := 1; 20*(2*i-1) < len(register); i++ {
		index := 20 * (i*2 - 1)
		sum += index * register[index-1]
	}
	fmt.Fprintf(os.Stdout, "p1: sum: %d\n", sum)

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			v := i*40 + j + 1
			if register[v-1]-1 <= j && register[v-1]+1 >= j {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		println()
	}
}

var next int

func handler(line string) error {
	if line == "" {
		return nil
	}
	parts := strings.Split(line, " ")
	if next != 0 {
		last := register[len(register)-1]
		register = append(register, last+next, last+next)
	} else {
		register = append(register, register[len(register)-1])
	}
	switch parts[0] {
	case "noop":
		next = 0
	case "addx":
		value, _ := strconv.Atoi(parts[1])
		next = value
	}
	return nil
}
