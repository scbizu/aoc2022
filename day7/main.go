package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/scbizu/aoc2022/helper/input"
)

type nodeKind int

const (
	kindDir nodeKind = iota
	kindFile
)

type fsTree struct {
	root *inode
}

type inode struct {
	parent   *inode
	children []*inode
	kind     nodeKind
	name     string
	size     int64
}

var tree = fsTree{
	root: &inode{
		parent: nil,
		kind:   kindDir,
		name:   "/",
	},
}

var (
	currentNode     *inode = tree.root
	currentNodeName string = "/"
)

func (in *inode) addFileNode(name string, size int64) *inode {
	inode := &inode{
		parent: in,
		kind:   kindFile,
		name:   name,
		size:   size,
	}
	in.children = append(in.children, inode)
	return inode
}

func (in *inode) addDirNode(name string) *inode {
	inode := &inode{
		parent: in,
		kind:   kindDir,
		name:   name,
	}
	in.children = append(in.children, inode)
	return inode
}

func (in *inode) findNamedNode(name string) *inode {
	if name == "/" {
		return tree.root
	}
	for _, child := range in.children {
		if child.name == name {
			return child
		}
		child.findNamedNode(name)
	}
	return nil
}

func (in *inode) traverse(fn func(i *inode)) {
	fn(in)
	for _, child := range in.children {
		child.traverse(fn)
	}
}

func (in *inode) cal() {
	if in.size == 0 && in.kind == kindDir {
		for _, child := range in.children {
			child.cal()
			in.size += child.size
		}
	}
}

func main() {
	var size int64
	input.NewTXTFile("input.txt").ReadByBlockEx(context.Background(), sep, wHandler())
	tree.root.cal()
	tree.root.traverse(func(i *inode) {
		if i.kind == kindDir && i.size <= 100000 {
			size += i.size
		}
	})
	fmt.Fprintf(os.Stdout, "p1: size: %d\n", size)
	var min int64 = math.MaxInt64
	tree.root.traverse(func(i *inode) {
		if i.kind == kindDir && i.size > 30000000-(70000000-tree.root.size) {
			if i.size < min {
				min = i.size
			}
		}
	})
	fmt.Fprintf(os.Stdout, "p2: min: %d\n", min)
}

func wHandler() func(lines []string) error {
	return func(lines []string) error {
		for _, line := range lines {
			switch {
			case strings.HasPrefix(line, "$ cd"):
				currentNodeName = strings.TrimPrefix(line, "$ cd ")
				if currentNodeName != ".." {
					currentNode = currentNode.findNamedNode(currentNodeName)
				} else {
					currentNode = currentNode.parent
					currentNodeName = currentNode.name
				}
			case line == "$ ls":
			case strings.HasPrefix(line, "dir"):
				currentNode.addDirNode(strings.TrimPrefix(line, "dir "))
			// file node
			default:
				parts := strings.Split(line, " ")
				size, err := strconv.Atoi(parts[0])
				if err != nil {
					return fmt.Errorf("failed to parse size: %w", err)
				}
				currentNode.addFileNode(parts[1], int64(size))
			}
		}
		return nil
	}
}

func sep(i int, line string) bool {
	return strings.HasPrefix(line, "$ cd")
}
