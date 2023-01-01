package main

import (
	"context"
	"fmt"
	"os"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/seq"
)

const (
	baseNumber  = 811589153
	refillTimes = 10
)

type number struct {
	index int
	value int
}

func main() {
	ctx := context.Background()
	input.NewTXTFile("input.txt").ReadByLine(ctx, handler)
	for _, num := range q {
		seqIndex := sequence.Find(num)
		sequence = sequence.Move(seqIndex, num.value)
	}
	zeroIndex := sequence.Find(number{index: originZeroIndex, value: 0})
	// fmt.Printf("%d,%d,%d\n", sequence.AtIndex(1000+zeroIndex), sequence.AtIndex(2000+zeroIndex), sequence.AtIndex(3000+zeroIndex))
	fmt.Fprintf(os.Stdout, "p1: res: %d\n", sequence.AtIndex(1000+zeroIndex).value+sequence.AtIndex(2000+zeroIndex).value+sequence.AtIndex(3000+zeroIndex).value)

	// reset
	sequence = seq.NewCircular([]number{})
	index, originZeroIndex = 0, 0
	q = []number{}

	input.NewTXTFile("input.txt").ReadByLine(ctx, handler)
	circleLen := sequence.BaseLen() - 1

	for i := 0; i < refillTimes; i++ {
		for _, num := range q {
			seqIndex := sequence.Find(num)
			sequence = sequence.Move(seqIndex, (num.value*baseNumber)%circleLen)
		}
		// sequence.Print()
	}

	zeroIndex = sequence.Find(number{index: originZeroIndex, value: 0})
	fmt.Fprintf(os.Stdout, "p2: res: %d\n", (sequence.AtIndex(1000+zeroIndex).value+sequence.AtIndex(2000+zeroIndex).value+sequence.AtIndex(3000+zeroIndex).value)*baseNumber)
}

var sequence = seq.NewCircular([]number{})

var q = []number{}

var (
	index           int
	originZeroIndex int
)

func handler(line string) error {
	if input.Atoi(line) == 0 {
		originZeroIndex = index
	}
	sequence.Append(number{
		index: index,
		value: input.Atoi(line),
	})
	q = append(q, number{
		index: index,
		value: input.Atoi(line),
	})
	index++
	return nil
}
