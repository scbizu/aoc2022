package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/input"
)

var (
	interCount   int
	interCountP2 int
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	fmt.Fprintln(os.Stdout, "P1:", interCount)
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handlerP2)
	fmt.Fprintln(os.Stdout, "P2", interCountP2)
}

func handler(line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	sections := strings.Split(parts[0], "-")
	if len(sections) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	from, _ := strconv.Atoi(sections[0])
	to, _ := strconv.Atoi(sections[1])
	p1set := set.New[int]()
	for i := from; i <= to; i++ {
		p1set.Add(i)
	}
	sections = strings.Split(parts[1], "-")
	if len(sections) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	from, _ = strconv.Atoi(sections[0])
	to, _ = strconv.Atoi(sections[1])
	p2set := set.New[int]()
	for i := from; i <= to; i++ {
		p2set.Add(i)
	}
	if set.Intersection(p1set, p2set).Size() == p1set.Size() || set.Intersection(p1set, p2set).Size() == p2set.Size() {
		interCount++
	}
	return nil
}

func handlerP2(line string) error {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	sections := strings.Split(parts[0], "-")
	if len(sections) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	from, _ := strconv.Atoi(sections[0])
	to, _ := strconv.Atoi(sections[1])
	p1set := set.New[int]()
	for i := from; i <= to; i++ {
		p1set.Add(i)
	}
	sections = strings.Split(parts[1], "-")
	if len(sections) != 2 {
		return fmt.Errorf("invalid line: %s", line)
	}
	from, _ = strconv.Atoi(sections[0])
	to, _ = strconv.Atoi(sections[1])
	p2set := set.New[int]()
	for i := from; i <= to; i++ {
		p2set.Add(i)
	}
	if set.Intersection(p1set, p2set).Size() > 0 {
		interCountP2++
	}
	return nil
}
