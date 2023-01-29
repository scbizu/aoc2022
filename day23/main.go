package main

import (
	"context"
	"fmt"
	"math"
	"os"

	"github.com/magejiCoder/set"
	"github.com/scbizu/aoc2022/helper/grid"
	"github.com/scbizu/aoc2022/helper/input"
	"github.com/scbizu/aoc2022/helper/seq"
)

var (
	board  map[grid.Vec]rune          = make(map[grid.Vec]rune)
	elfLoc *set.Set[grid.Vec]         = set.New[grid.Vec]()
	dirs   *seq.Circular[dirCategory] = seq.NewCircular([]dirCategory{dirNorth, dirSouth, dirWest, dirEast})
)

const (
	tryTimes = 10
)

type dirCategory uint8

type elf struct {
	grid.Vec
	cIndex int
}

const (
	// N,NW,NE
	dirNorth dirCategory = iota
	// S,SW,SE
	dirSouth
	// W,NW,SW
	dirWest
	// E,NE,SE
	dirEast
)

func main() {
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	m := grid.NewVecMatrix[rune]()
	for g, e := range board {
		m.Add(g, e)
	}
	m.Print(os.Stdout, "%c")
	elfs := make(map[grid.Vec]*elf)
	elfLoc.Each(func(item grid.Vec) bool {
		elfs[item] = &elf{
			Vec:    item,
			cIndex: int(dirNorth),
		}
		// check if there is any elf in the 8 neighbors
		s8 := m.GetNeighbors8(item)
		if elfLoc.HasAny(s8...) {
			return true
		}
		// fmt.Printf("no elf in 8 neighbors of %v\n", item)
		elfLoc.Remove(item)
		return true
	})
	bs := boardState{
		board:          m,
		dir:            dirNorth,
		elfLocs:        elfLoc,
		elfs:           elfs,
		proposedLocs:   set.New[grid.Vec](),
		proposeMapping: map[grid.Vec]grid.Vec{},
		emptySlots:     len(board) - len(elfs),
	}
	bs = moveUntilTimes(bs, tryTimes)

	fmt.Fprintf(os.Stdout, "p1: empty slots: %d\n", bs.emptySlots)

	elfLoc = set.New[grid.Vec]()
	input.NewTXTFile("input.txt").ReadByLine(context.Background(), handler)
	m = grid.NewVecMatrix[rune]()
	for g, e := range board {
		m.Add(g, e)
	}
	m.Print(os.Stdout, "%c")
	elfs = make(map[grid.Vec]*elf)

	elfLoc.Each(func(item grid.Vec) bool {
		elfs[item] = &elf{
			Vec:    item,
			cIndex: int(dirNorth),
		}
		// check if there is any elf in the 8 neighbors
		s8 := m.GetNeighbors8(item)
		if elfLoc.HasAny(s8...) {
			return true
		}
		// fmt.Printf("no elf in 8 neighbors of %v\n", item)
		elfLoc.Remove(item)
		return true
	})
	bs = boardState{
		board:          m,
		dir:            dirNorth,
		elfLocs:        elfLoc,
		elfs:           elfs,
		proposedLocs:   set.New[grid.Vec](),
		proposeMapping: map[grid.Vec]grid.Vec{},
		emptySlots:     len(board) - len(elfs),
		round:          0,
	}
	bs = moveUntilSilent(bs)

	fmt.Fprintf(os.Stdout, "p2: round: %d\n", bs.round+1)
}

type boardState struct {
	board          grid.VecMatrix[rune]
	dir            dirCategory
	elfs           map[grid.Vec]*elf
	elfLocs        *set.Set[grid.Vec]
	proposedLocs   *set.Set[grid.Vec]
	proposeMapping map[grid.Vec]grid.Vec
	emptySlots     int
	round          int
}

func (bs boardState) reset() boardState {
	bs.elfLocs = set.New[grid.Vec]()
	bs.dir += 1
	for vec := range bs.elfs {
		if vec2, ok := bs.proposeMapping[vec]; ok {
			bs.elfs[vec2] = &elf{
				Vec:    vec2,
				cIndex: int(bs.dir),
			}
			delete(bs.elfs, vec)
			bs.elfLocs.Add(vec2)
		} else {
			bs.elfs[vec] = &elf{
				Vec:    vec,
				cIndex: int(bs.dir),
			}
			bs.elfLocs.Add(vec)
		}
	}
	bs.elfLocs.Each(func(item grid.Vec) bool {
		// check if there is any elf in the 8 neighbors
		s8 := bs.board.GetNeighbors8(item)
		if bs.elfLocs.HasAny(s8...) {
			return true
		}
		// fmt.Printf("no elf in 8 neighbors of %s\n", item)
		bs.elfLocs.Remove(item)
		return true
	})
	bs.proposedLocs = set.New[grid.Vec]()
	bs.proposeMapping = map[grid.Vec]grid.Vec{}

	east, south := math.MaxInt, math.MaxInt
	west, north := math.MinInt, math.MinInt

	for vec := range bs.elfs {
		if vec.X < east {
			east = vec.X
		}
		if vec.X > west {
			west = vec.X
		}
		if vec.Y < south {
			south = vec.Y
		}
		if vec.Y > north {
			north = vec.Y
		}
	}

	bs.board = grid.NewVecMatrix[rune]()

	for y := south; y <= north; y++ {
		for x := east; x <= west; x++ {
			if _, ok := bs.elfs[grid.Vec{X: x, Y: y}]; ok {
				bs.board.Add(grid.Vec{X: x, Y: y}, '#')
			} else {
				bs.board.Add(grid.Vec{X: x, Y: y}, '.')
			}
		}
	}

	// bs.board.Print("%c")

	bs.emptySlots = (north-south+1)*(west-east+1) - len(bs.elfs)

	return bs
}

func getLocs(raw grid.Vec, d dirCategory) []grid.Vec {
	switch d {
	case dirNorth:
		return []grid.Vec{raw.Add(grid.Vec{X: 0, Y: -1}), raw.Add(grid.Vec{X: -1, Y: -1}), raw.Add(grid.Vec{X: 1, Y: -1})}
	case dirSouth:
		return []grid.Vec{raw.Add(grid.Vec{X: 0, Y: 1}), raw.Add(grid.Vec{X: -1, Y: 1}), raw.Add(grid.Vec{X: 1, Y: 1})}
	case dirWest:
		return []grid.Vec{raw.Add(grid.Vec{X: -1, Y: 0}), raw.Add(grid.Vec{X: -1, Y: -1}), raw.Add(grid.Vec{X: -1, Y: 1})}
	case dirEast:
		return []grid.Vec{raw.Add(grid.Vec{X: 1, Y: 0}), raw.Add(grid.Vec{X: 1, Y: -1}), raw.Add(grid.Vec{X: 1, Y: 1})}
	}
	panic("invalid dirCategory")
}

func moveDir(raw grid.Vec, d dirCategory) grid.Vec {
	switch d {
	case dirNorth:
		return raw.Add(grid.Vec{X: 0, Y: -1})
	case dirSouth:
		return raw.Add(grid.Vec{X: 0, Y: 1})
	case dirWest:
		return raw.Add(grid.Vec{X: -1, Y: 0})
	case dirEast:
		return raw.Add(grid.Vec{X: 1, Y: 0})
	}
	panic("invalid dirCategory")
}

func moveUntilSilent(bs boardState) boardState {
	for {
		bs.elfLocs.Each(func(item grid.Vec) bool {
			if _, ok := bs.elfs[item]; !ok {
				panic("invalid elf loc")
			}
			// first half
			bs = move(bs, bs.elfs[item], 0)
			return true
		})
		if len(bs.proposeMapping) == 0 {
			return bs
		}
		// second half
		// process the proposed locs
		// reset board
		bs = bs.reset()
		bs.round += 1
	}
}

func moveUntilTimes(bs boardState, times int) boardState {
	for i := 0; i < times; i++ {
		bs.elfLocs.Each(func(item grid.Vec) bool {
			if _, ok := bs.elfs[item]; !ok {
				panic("invalid elf loc")
			}
			// first half
			bs = move(bs, bs.elfs[item], 0)
			return true
		})
		// second half
		// process the proposed locs
		// reset board
		bs = bs.reset()
	}

	return bs
}

func move(bs boardState, e *elf, index int) boardState {
	if index >= dirs.BaseLen() {
		// fmt.Println("stay")
		return bs
	}
	// fmt.Printf("elf at location %s\n", e.Vec)
	locs := getLocs(e.Vec, dirs.AtIndex(e.cIndex))
	if bs.elfLocs.HasAny(locs...) {
		e.cIndex += 1
		bs = move(bs, e, index+1)
	} else {
		newVec := moveDir(e.Vec, dirs.AtIndex(e.cIndex))
		if bs.proposedLocs.HasAny(newVec) {
			// fmt.Printf("proposing failed: %s\n", newVec)
			bs.proposedLocs.Remove(newVec)
			bs.proposeMapping = withdrawProposal(bs.proposeMapping, newVec)
		} else {
			// fmt.Printf("proposing to move to %s\n", newVec)
			bs.proposedLocs.Add(newVec)
			bs.proposeMapping[e.Vec] = newVec
		}
		e.cIndex += 1
	}
	return bs
}

func withdrawProposal(m map[grid.Vec]grid.Vec, v grid.Vec) map[grid.Vec]grid.Vec {
	for k, v2 := range m {
		if v2 == v {
			delete(m, k)
		}
	}
	return m
}

var x, y int

func handler(line string) error {
	x = len(line)
	for col, c := range line {
		xy := grid.Vec{
			X: col,
			Y: y,
		}
		if c == '#' {
			elfLoc.Add(xy)
		}
		board[xy] = c
	}
	y++
	return nil
}
