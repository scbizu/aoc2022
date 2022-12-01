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
	if err := f.ReadByBlock(context.Background(), "\n\n", countHandler); err != nil {
		log.Fatalln(err)
	}
}

func countHandler(block []string) error {
	var maxCarry int64
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
		if carry > maxCarry {
			maxCarry = carry
		}
	}
	fmt.Fprintf(os.Stdout, "max carry: %d\n", maxCarry)
	return nil
}
