package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

func main() {
	f := input.NewTXTFile("input.txt")
	if err := f.ReadByBlock(context.Background(), "\n\n", countTop3Handler); err != nil {
		log.Fatalln(err)
	}
}

func countTop3Handler(block []string) error {
	// create top 3 carry
	top3Carry := [3]int64{0, 0, 0}
	for _, lines := range block {
		if lines == "" {
			continue
		}
		var carry int64
		for _, line := range strings.Split(lines, "\n") {
			if line == "" {
				continue
			}
			i, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				return err
			}
			carry += i
		}
		if carry > top3Carry[0] {
			top3Carry[2] = top3Carry[1]
			top3Carry[1] = top3Carry[0]
			top3Carry[0] = carry
			continue
		}
		if carry > top3Carry[1] {
			top3Carry[2] = top3Carry[1]
			top3Carry[1] = carry
			continue
		}
		if carry > top3Carry[2] {
			top3Carry[2] = carry
			continue
		}
	}
	fmt.Fprintf(os.Stdout, "carry: %d\n", top3Carry[0]+top3Carry[1]+top3Carry[2])
	return nil
}
