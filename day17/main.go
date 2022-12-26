package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

const (
	maxDrop = 2022
	// magic number , to assert that last seen 128 lines are different enough
	detectLoopRows = 128
	maxDetectRows  = 100000
	target         = 1000000000000
)

func main() {
	tall := make(map[int]int)
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	for i := 0; i < maxDrop; i++ {
		drop(i)
		tall[i] = len(tetrisBuffer)
	}
	fmt.Fprintf(os.Stdout, "p1: res: %d\n", len(tetrisBuffer))
	tetrisBuffer = [][7]byte{}
	jetOffset = 0
	var i int
	loopState := make(map[string]int)
	var first, loop int
	var delta int
	for {
		if i > maxDetectRows {
			panic("try much larger detect number")
		}
		drop(i)
		if len(tetrisBuffer) > detectLoopRows {
			hash := hash(tetrisBuffer[:detectLoopRows])
			if last, ok := loopState[hash]; ok {
				first = last
				loop = i - last
				delta = len(tetrisBuffer) - tall[last]
				fmt.Fprintf(os.Stdout, "p2: found circle: %d,%d\n", i, last)
				fmt.Fprintf(os.Stdout, "p2: found circle: %d\n", loop)
				fmt.Fprintf(os.Stdout, "p2: found circle:tall %d\n", delta)
				break
			}
			loopState[hash] = i
		}
		i++
	}
	t := target - 1
	n := (t - first) / loop
	left := (t - first) % loop
	// -     | ....xx.|
	// left  | ....xx.| .  because the middle(delta) is collapsed , the left piece will be drop into `first`
	// -     | ..xxxx.|    so we should plus times (left + first) , say , we first drop 3 times , and left 3 times, and sub the first part's tall at last
	// -     | .......|
	// delta | .......|    can be collapsed , and total tall is n(loop times)*delta
	// -     | .......|
	// -     | ....xx.|
	// first | ....xx.|
	// -     | ..xxxx.|
	fmt.Fprintf(os.Stdout, "p2: res: %d\n", tall[first]+n*delta+tall[first+left]-tall[first])
}

// hash hashes the input with md5
func hash(s ShapePosition) string {
	// hash with md5
	h := md5.New()
	for _, row := range s {
		_, _ = h.Write(row[:])
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

type Shape int32

const (
	ShapeSub Shape = iota
	ShapePlus
	// the reverse L
	ShapeLR
	ShapeI
	ShapeSquare
)

func (s Shape) loc() ShapePosition {
	switch s {
	case ShapeSub:
		return [][7]byte{
			{0, 0, 1, 1, 1, 1, 0},
		}
	case ShapePlus:
		return [][7]byte{
			{0, 0, 0, 1, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 0, 0, 0},
		}
	case ShapeLR:
		return [][7]byte{
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 0, 0},
			{0, 0, 1, 1, 1, 0, 0},
		}
	case ShapeI:
		return [][7]byte{
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
		}
	case ShapeSquare:
		return [][7]byte{
			{0, 0, 1, 1, 0, 0, 0},
			{0, 0, 1, 1, 0, 0, 0},
		}
	}
	return nil
}

func left1(s ShapePosition) (ShapePosition, bool) {
	var new [][7]byte
	for _, row := range s {
		newLine := [7]byte{}
		for i := 0; i < 7; i++ {
			if row[i] == 1 {
				if i == 0 {
					return s, false
				}
				newLine[i-1] = 1
				newLine[i] = 0
			}
		}
		new = append(new, newLine)
	}
	return new, true
}

type ShapePosition [][7]byte

func (s ShapePosition) Print() {
	for _, row := range s {
		fmt.Print("|")
		for _, col := range row {
			if col == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
		fmt.Println()
	}
}

func right1(s ShapePosition) (ShapePosition, bool) {
	var new [][7]byte
	for _, row := range s {
		newLine := [7]byte{}
		for i := 6; i >= 0; i-- {
			if row[i] == 1 {
				if i == 6 {
					return s, false
				}
				newLine[i+1] = 1
				newLine[i] = 0
			}
		}
		new = append(new, newLine)
	}
	return new, true
}

func op(s ShapePosition, jp byte) (ShapePosition, bool) {
	switch jp {
	case '<': // left
		return left1(s)
	case '>': // right
		return right1(s)
	}
	panic("invalid op")
}

func newShape(times int) Shape {
	return Shape(times % 5)
}

var jetBuffer = bytes.NewBuffer(nil)

func handler(line string) error {
	line = strings.TrimSpace(line)
	jetBuffer.WriteString(line)
	return nil
}

var jetOffset int

func readJet() [1]byte {
	jet := jetBuffer.Bytes()[jetOffset%len(jetBuffer.Bytes())]
	jetOffset++
	return [1]byte{jet}
}

var tetrisBuffer = [][7]byte{}

func drop(times int) {
	s := newShape(times)
	sp := s.loc()
	for i := 0; i < 4; i++ {
		jet := readJet()
		sp, _ = op(sp, jet[0])
	}
	tetrisBuffer = mergeMatrix(tetrisBuffer, sp, s)
}

func merge(l1, l2 [7]byte) ([7]byte, bool) {
	var new [7]byte
	for i := 0; i < 7; i++ {
		if l1[i]&l2[i] == 1 {
			return new, false
		}
		new[i] = l1[i] | l2[i]
	}
	return new, true
}

func isOverlap(sp1, sp2 ShapePosition) bool {
	if len(sp1) == 0 {
		return false
	}
	if len(sp1) < len(sp2) {
		// reach the bottom
		return true
	}
	for i := 0; i < len(sp2); i++ {
		if _, ok := merge(sp1[i], sp2[i]); !ok {
			return true
		}
	}
	return false
}

func readLine(s ShapePosition) (ShapePosition, [7]byte) {
	if len(s) == 0 {
		return s, [7]byte{0, 0, 0, 0, 0, 0, 0}
	}
	sp := s[len(s)-1]
	s = s[:len(s)-1]
	return s, sp
}

func mergeAll(sp1, sp2 ShapePosition) ShapePosition {
	if len(sp1) == 0 {
		return sp2
	}
	if len(sp1) < len(sp2) {
		panic("base should be longer than new")
	}
	for i := 0; i < len(sp2); i++ {
		sp1[i], _ = merge(sp1[i], sp2[i])
	}
	return sp1
}

func mergeMatrix(sp1 ShapePosition, sp2 ShapePosition, _ Shape) ShapePosition {
	if len(sp1) == 0 {
		return sp2
	}
	var new ShapePosition
	sp2, item := readLine(sp2)
	newSp2 := ShapePosition([][7]byte{item})
	// move
	for {
		// sp1 conflict with sp2
		if isOverlap(sp1, newSp2) {
			if item != [7]byte{0, 0, 0, 0, 0, 0, 0} {
				sp2 = append(sp2, item)
			}
			new = mergeAll(sp1, newSp2[1:])
			break
		}
		jet := readJet()[0]
		tmp, ok := op(newSp2, jet)
		// if we can accept this move
		if !isOverlap(sp1, tmp) && ok {
			var ok bool
			sp2, ok = op(sp2, jet)
			if ok {
				newSp2 = tmp
			}
		}
		sp2, item = readLine(sp2)
		var sp ShapePosition
		sp = append(sp, item)
		sp = append(sp, newSp2...)
		newSp2 = sp
	}
	var res ShapePosition
	res = append(res, sp2...)
	res = append(res, new...)
	return res
}
