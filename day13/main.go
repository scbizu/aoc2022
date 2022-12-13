package main

import (
	"container/list"
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/stack"
)

func main() {
	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", handler)
	input.NewTXTFile("input.txt").ReadByBlock(context.Background(), "\n\n", handler2)
}

type kind int

const (
	kindInt kind = iota
	kindList
)

type item struct {
	parent   *item
	children []*item
	// if kind is kindList, data is []int
	// if kind is kindInt, data is int
	t    kind
	data any
}

func (i *item) print() {
	if i.t == kindInt {
		fmt.Printf("%d", i.data.(int))
		return
	}
	fmt.Printf("[")
	for _, c := range i.children {
		c.print()
	}
	fmt.Printf("]")
}

func (i *item) String() string {
	if i.t == kindInt {
		return fmt.Sprintf("%d", i.data.(int))
	}
	var s strings.Builder
	s.WriteString("[")
	for _, c := range i.children {
		s.WriteString(c.String())
	}
	s.WriteString("]")
	return s.String()
}

type cmp int

const (
	cmpTrue cmp = iota
	cmpFalse
	cmpEqual
)

func (c cmp) String() string {
	switch c {
	case cmpTrue:
		return "true"
	case cmpFalse:
		return "false"
	case cmpEqual:
		return "equal"
	}
	return "unknown"
}

func (i *item) InRightOrder(other *item) cmp {
	if i.t == kindInt && other.t == kindInt {
		switch {
		case i.data.(int) == other.data.(int):
			return cmpEqual
		case i.data.(int) < other.data.(int):
			return cmpTrue
		case i.data.(int) > other.data.(int):
			return cmpFalse
		}
	}
	if i.t == kindList && other.t == kindList {
		for i, c := range i.children {
			if i >= len(other.children) {
				return cmpFalse
			}
			order := c.InRightOrder(other.children[i])
			if order != cmpEqual {
				return order
			}
		}
		if len(i.children) < len(other.children) {
			return cmpTrue
		}
		i.parent.children = i.parent.children[1:]
		other.parent.children = other.parent.children[1:]
		return i.parent.InRightOrder(other.parent)
	}
	if i.t == kindInt && other.t == kindList {
		i.t = kindList
		i.children = append(i.children, &item{
			t:      kindInt,
			data:   i.data,
			parent: i,
		})
		return i.InRightOrder(other)
	}
	if i.t == kindList && other.t == kindInt {
		other.t = kindList
		other.children = append(other.children, &item{
			t:      kindInt,
			data:   other.data,
			parent: other,
		})
		return i.InRightOrder(other)
	}

	panic("unreachable")
}

func indexBracketMap(line string) map[int]int {
	index := make(map[int]int)
	s := stack.Stack[int]{
		List: list.New(),
	}
	for i, c := range line {
		switch c {
		case '[':
			s.Push(i)
		case ']':
			index[i] = s.Pop()
		}
	}
	indexRev := make(map[int]int)
	for k, v := range index {
		indexRev[v] = k
	}
	return indexRev
}

func newItemFromLine(node *item, line string) *item {
	indexMap := indexBracketMap(line)
	if len(line) == 0 {
		return node
	}
	newNode := &item{}
	switch {
	case strings.HasPrefix(line, "["):
		index := strings.Index(line, "[")
		newNode.t = kindList
		newNode.data = []int{}
		// trim [ ]
		newItemFromLine(newNode, line[index+1:indexMap[index]])
		for _, c := range newNode.children {
			if c.t == kindInt {
				newNode.data = append(newNode.data.([]int), c.data.(int))
			}
		}
		newNode.parent = node
		node.children = append(node.children, newNode)
		if indexMap[index] == len(line)-1 {
			return newNode
		}
		return newItemFromLine(node, line[indexMap[index]+1:])
	case strings.HasPrefix(line, ","):
		return newItemFromLine(node, line[1:])
	case strings.HasPrefix(line, "]"):
		return newItemFromLine(node, line[1:])
	default:
		newNode.t = kindInt
		var readOffset int
		var reads string
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == ',' || c == ']' {
				readOffset = i
				break
			}
			reads += string(c)
			readOffset++
		}
		readin, _ := strconv.Atoi(reads)
		newNode.data = readin
		newNode.parent = node
		node.children = append(node.children, newNode)
		return newItemFromLine(node, line[readOffset:])
	}
}

func handler(line []string) error {
	if len(line) == 0 {
		return nil
	}
	var sum int
	for index, part := range line {
		l := strings.Split(part, "\n")
		im := newItemFromLine(&item{}, l[0])
		im.print()
		println()
		im2 := newItemFromLine(&item{}, l[1])
		im2.print()
		println()
		if im.InRightOrder(im2) == cmpTrue {
			sum += index + 1
		}
	}
	fmt.Fprintf(os.Stdout, "p1:res: %d\n", sum)
	return nil
}

func handler2(line []string) error {
	if len(line) == 0 {
		return nil
	}
	var its []string
	for _, part := range line {
		l := strings.Split(part, "\n")
		its = append(its, l[0], l[1])
	}
	its = append(its, "[[2]]", "[[6]]")

	sort.SliceStable(its, func(i, j int) bool {
		return newItemFromLine(&item{}, its[i]).InRightOrder(newItemFromLine(&item{}, its[j])) == cmpTrue
	})

	var i1, i2 int
	for i, it := range its {
		if it == "[[2]]" {
			i1 = i + 1
		}
		if it == "[[6]]" {
			i2 = i + 1
		}
	}
	fmt.Fprintf(os.Stdout, "p2:res:  %d\n", i1*i2)
	return nil
}
